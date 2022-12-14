package main

import (
	"embed"
	"flag"
	"fmt"
	"go/token"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongmattree/go/controllers"
	"github.com/fullstack-lang/gongmattree/go/models"
	"github.com/fullstack-lang/gongmattree/go/orm"

	// gong stack for model analysis
	gong_controllers "github.com/fullstack-lang/gong/go/controllers"
	gong_models "github.com/fullstack-lang/gong/go/models"
	gong_orm "github.com/fullstack-lang/gong/go/orm"

	// insertion point for gong front end import
	_ "github.com/fullstack-lang/gong/ng"

	// for diagrams
	gongdoc_controllers "github.com/fullstack-lang/gongdoc/go/controllers"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_orm "github.com/fullstack-lang/gongdoc/go/orm"

	// insertion point for gong front end import
	_ "github.com/fullstack-lang/gongdoc/ng"

	gongmattree "github.com/fullstack-lang/gongmattree"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	marshallOnStartup = flag.String("marshallOnStartup", "", "at startup, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	unmarshall        = flag.String("unmarshall", "", "unmarshall data from marshall name and '.go' (must be lowercased without spaces), If unmarshall arg is '', no unmarshalling")
	marshallOnCommit  = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams (takes a few seconds)")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	models.Stage.Checkout()
	models.Stage.Marshall(file, "github.com/fullstack-lang/gongmattree/go/models", "main")
}

func main() {

	log.SetPrefix("gongmattree: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM
	db := orm.SetupModels(*logDBFlag, "./test.db")
	dbDB, err := db.DB()

	//
	// gong and gongdoc databases do not need to be persisted.
	// therefore, they are in memory
	//
	db_inMemory := gong_orm.SetupModels(*logDBFlag, ":memory:")

	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db_inMemory.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	// add gongdocatabase
	gongdoc_orm.AutoMigrate(db_inMemory)

	// generate injection code from the stage
	if *marshallOnStartup != "" {

		if strings.Contains(*marshallOnStartup, " ") {
			log.Fatalln(*marshallOnStartup + " must not contains blank spaces")
		}
		if strings.ToLower(*marshallOnStartup) != *marshallOnStartup {
			log.Fatalln(*marshallOnStartup + " must be lowercases")
		}

		file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnStartup))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		models.Stage.Checkout()
		models.Stage.Marshall(file, "github.com/fullstack-lang/gongmattree/go/models", "main")
		os.Exit(0)
	}

	// setup the stage by injecting the code from code database
	if *unmarshall != "" {
		models.Stage.Checkout()
		models.Stage.Reset()
		models.Stage.Commit()
		if InjectionGateway[*unmarshall] != nil {
			InjectionGateway[*unmarshall]()
		}
		models.Stage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		models.Stage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		models.Stage.OnInitCommitFromFrontCallback = hook
	}

	// since the stack can be a multi threaded application. It is important to set up
	// only one open connexion at a time
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)

	if *diagrams {

		// Analyse package
		modelPkg := &gong_models.ModelPkg{}

		// since the source is embedded, one needs to
		// compute the Abstract syntax tree in a special manner
		pkgs := gong_models.ParseEmbedModel(gongmattree.GoDir, "go/models")

		gong_models.WalkParser(pkgs, modelPkg)
		modelPkg.SerializeToStage()
		gong_models.Stage.Commit()

		// create the diagrams
		// prepare the model views
		pkgelt := new(gongdoc_models.Pkgelt)

		// first, get all gong struct in the model
		for gongStruct := range gong_models.Stage.GongStructs {

			// let create the gong struct in the gongdoc models
			// and put the numbre of instances
			gongStruct_ := (&gongdoc_models.GongStruct{Name: gongStruct.Name}).Stage()
			nbInstances, ok := models.Stage.Map_GongStructName_InstancesNb[gongStruct.Name]
			if ok {
				gongStruct_.NbInstances = nbInstances
			}
		}

		// classdiagram can only be fully in memory when they are Unmarshalled
		// for instance, the Name of diagrams or the Name of the Link
		fset := new(token.FileSet)
		pkgsParser := gong_models.ParseEmbedModel(gongmattree.GoDir, "go/diagrams")
		if len(pkgsParser) != 1 {
			log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
		}
		if pkgParser, ok := pkgsParser["diagrams"]; ok {
			pkgelt.Unmarshall(modelPkg, pkgParser, fset, "go/diagrams")
		}
		pkgelt.SerializeToStage()
	}

	controllers.RegisterControllers(r)
	gongdoc_controllers.RegisterControllers(r)
	gong_controllers.RegisterControllers(r)
	gongdoc_models.Stage.Commit()
	gong_models.Stage.Commit()

	// intercept updates to the nodes
	// hook something on models.Stage.OnInitCommitFromFrontCallback

	// insertion point for serving the static file
	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(gongmattree.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	// get callbacks on node updates
	var onNodeCallbackStruct NodeCallbackStruct
	models.Stage.OnAfterNodeUpdateCallback = onNodeCallbackStruct
	models.Stage.OnAfterNodeCreateCallback = onNodeCallbackStruct
	models.Stage.OnAfterNodeDeleteCallback = onNodeCallbackStruct
	models.Stage.OnAfterNodeReadCallback = onNodeCallbackStruct

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

type NodeCallbackStruct struct {
}

func (onNodeUpdateCallbackStruct NodeCallbackStruct) OnAfterUpdate(
	stage *models.StageStruct,
	old, new *models.Node) {

	if old.IsChecked != new.IsChecked {
		log.Println("Node " + old.Name + " is updated with value IsChecked to " + strconv.FormatBool(new.IsChecked))
	}
	if old.IsExpanded != new.IsExpanded {
		log.Println("Node " + old.Name + " is updated with value IsExpanded to " + strconv.FormatBool(new.IsExpanded))
	}
}

func (onNodeUpdateCallbackStruct NodeCallbackStruct) OnAfterCreate(
	stage *models.StageStruct,
	node *models.Node) {
	log.Println("Node created " + node.Name)
}

func (onNodeUpdateCallbackStruct NodeCallbackStruct) OnAfterDelete(
	stage *models.StageStruct,
	old, new *models.Node) {
	log.Println("Node delete " + new.Name)
}

func (onNodeUpdateCallbackStruct NodeCallbackStruct) OnAfterRead(
	stage *models.StageStruct,
	node *models.Node) {
	log.Println("Node read " + node.Name)
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
