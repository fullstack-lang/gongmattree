// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongdoc/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_GongdocStatus_sql sql.NullBool
var dummy_GongdocStatus_time time.Duration
var dummy_GongdocStatus_sort sort.Float64Slice

// GongdocStatusAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongdocstatusAPI
type GongdocStatusAPI struct {
	gorm.Model

	models.GongdocStatus

	// encoding of pointers
	GongdocStatusPointersEnconding
}

// GongdocStatusPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongdocStatusPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// GongdocStatusDB describes a gongdocstatus in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongdocstatusDB
type GongdocStatusDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongdocstatusDB.Name
	Name_Data sql.NullString

	// Declation for basic field gongdocstatusDB.Status
	Status_Data sql.NullString

	// Declation for basic field gongdocstatusDB.CommandCompletionDate
	CommandCompletionDate_Data sql.NullString
	// encoding of pointers
	GongdocStatusPointersEnconding
}

// GongdocStatusDBs arrays gongdocstatusDBs
// swagger:response gongdocstatusDBsResponse
type GongdocStatusDBs []GongdocStatusDB

// GongdocStatusDBResponse provides response
// swagger:response gongdocstatusDBResponse
type GongdocStatusDBResponse struct {
	GongdocStatusDB
}

// GongdocStatusWOP is a GongdocStatus without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongdocStatusWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Status models.GongdocCommandType `xlsx:"2"`

	CommandCompletionDate string `xlsx:"3"`
	// insertion for WOP pointer fields
}

var GongdocStatus_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Status",
	"CommandCompletionDate",
}

type BackRepoGongdocStatusStruct struct {
	// stores GongdocStatusDB according to their gorm ID
	Map_GongdocStatusDBID_GongdocStatusDB *map[uint]*GongdocStatusDB

	// stores GongdocStatusDB ID according to GongdocStatus address
	Map_GongdocStatusPtr_GongdocStatusDBID *map[*models.GongdocStatus]uint

	// stores GongdocStatus according to their gorm ID
	Map_GongdocStatusDBID_GongdocStatusPtr *map[uint]*models.GongdocStatus

	db *gorm.DB
}

func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) GetDB() *gorm.DB {
	return backRepoGongdocStatus.db
}

// GetGongdocStatusDBFromGongdocStatusPtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) GetGongdocStatusDBFromGongdocStatusPtr(gongdocstatus *models.GongdocStatus) (gongdocstatusDB *GongdocStatusDB) {
	id := (*backRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID)[gongdocstatus]
	gongdocstatusDB = (*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB)[id]
	return
}

// BackRepoGongdocStatus.Init set up the BackRepo of the GongdocStatus
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) Init(db *gorm.DB) (Error error) {

	if backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr != nil {
		err := errors.New("In Init, backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr should be nil")
		return err
	}

	if backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB != nil {
		err := errors.New("In Init, backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB should be nil")
		return err
	}

	if backRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID != nil {
		err := errors.New("In Init, backRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.GongdocStatus, 0)
	backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr = &tmp

	tmpDB := make(map[uint]*GongdocStatusDB, 0)
	backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB = &tmpDB

	tmpID := make(map[*models.GongdocStatus]uint, 0)
	backRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID = &tmpID

	backRepoGongdocStatus.db = db
	return
}

// BackRepoGongdocStatus.CommitPhaseOne commits all staged instances of GongdocStatus to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for gongdocstatus := range stage.GongdocStatuss {
		backRepoGongdocStatus.CommitPhaseOneInstance(gongdocstatus)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongdocstatus := range *backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr {
		if _, ok := stage.GongdocStatuss[gongdocstatus]; !ok {
			backRepoGongdocStatus.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongdocStatus.CommitDeleteInstance commits deletion of GongdocStatus to the BackRepo
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) CommitDeleteInstance(id uint) (Error error) {

	gongdocstatus := (*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr)[id]

	// gongdocstatus is not staged anymore, remove gongdocstatusDB
	gongdocstatusDB := (*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB)[id]
	query := backRepoGongdocStatus.db.Unscoped().Delete(&gongdocstatusDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID), gongdocstatus)
	delete((*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr), id)
	delete((*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB), id)

	return
}

// BackRepoGongdocStatus.CommitPhaseOneInstance commits gongdocstatus staged instances of GongdocStatus to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) CommitPhaseOneInstance(gongdocstatus *models.GongdocStatus) (Error error) {

	// check if the gongdocstatus is not commited yet
	if _, ok := (*backRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID)[gongdocstatus]; ok {
		return
	}

	// initiate gongdocstatus
	var gongdocstatusDB GongdocStatusDB
	gongdocstatusDB.CopyBasicFieldsFromGongdocStatus(gongdocstatus)

	query := backRepoGongdocStatus.db.Create(&gongdocstatusDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID)[gongdocstatus] = gongdocstatusDB.ID
	(*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr)[gongdocstatusDB.ID] = gongdocstatus
	(*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB)[gongdocstatusDB.ID] = &gongdocstatusDB

	return
}

// BackRepoGongdocStatus.CommitPhaseTwo commits all staged instances of GongdocStatus to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongdocstatus := range *backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr {
		backRepoGongdocStatus.CommitPhaseTwoInstance(backRepo, idx, gongdocstatus)
	}

	return
}

// BackRepoGongdocStatus.CommitPhaseTwoInstance commits {{structname }} of models.GongdocStatus to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongdocstatus *models.GongdocStatus) (Error error) {

	// fetch matching gongdocstatusDB
	if gongdocstatusDB, ok := (*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB)[idx]; ok {

		gongdocstatusDB.CopyBasicFieldsFromGongdocStatus(gongdocstatus)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoGongdocStatus.db.Save(&gongdocstatusDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongdocStatus intance %s", gongdocstatus.Name))
		return err
	}

	return
}

// BackRepoGongdocStatus.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) CheckoutPhaseOne() (Error error) {

	gongdocstatusDBArray := make([]GongdocStatusDB, 0)
	query := backRepoGongdocStatus.db.Find(&gongdocstatusDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongdocstatusInstancesToBeRemovedFromTheStage := make(map[*models.GongdocStatus]any)
	for key, value := range models.Stage.GongdocStatuss {
		gongdocstatusInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongdocstatusDB := range gongdocstatusDBArray {
		backRepoGongdocStatus.CheckoutPhaseOneInstance(&gongdocstatusDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongdocstatus, ok := (*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr)[gongdocstatusDB.ID]
		if ok {
			delete(gongdocstatusInstancesToBeRemovedFromTheStage, gongdocstatus)
		}
	}

	// remove from stage and back repo's 3 maps all gongdocstatuss that are not in the checkout
	for gongdocstatus := range gongdocstatusInstancesToBeRemovedFromTheStage {
		gongdocstatus.Unstage()

		// remove instance from the back repo 3 maps
		gongdocstatusID := (*backRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID)[gongdocstatus]
		delete((*backRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID), gongdocstatus)
		delete((*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB), gongdocstatusID)
		delete((*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr), gongdocstatusID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongdocstatusDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongdocstatusDB
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) CheckoutPhaseOneInstance(gongdocstatusDB *GongdocStatusDB) (Error error) {

	gongdocstatus, ok := (*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr)[gongdocstatusDB.ID]
	if !ok {
		gongdocstatus = new(models.GongdocStatus)

		(*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr)[gongdocstatusDB.ID] = gongdocstatus
		(*backRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID)[gongdocstatus] = gongdocstatusDB.ID

		// append model store with the new element
		gongdocstatus.Name = gongdocstatusDB.Name_Data.String
		gongdocstatus.Stage()
	}
	gongdocstatusDB.CopyBasicFieldsToGongdocStatus(gongdocstatus)

	// preserve pointer to gongdocstatusDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongdocStatusDBID_GongdocStatusDB)[gongdocstatusDB hold variable pointers
	gongdocstatusDB_Data := *gongdocstatusDB
	preservedPtrToGongdocStatus := &gongdocstatusDB_Data
	(*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB)[gongdocstatusDB.ID] = preservedPtrToGongdocStatus

	return
}

// BackRepoGongdocStatus.CheckoutPhaseTwo Checkouts all staged instances of GongdocStatus to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongdocstatusDB := range *backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB {
		backRepoGongdocStatus.CheckoutPhaseTwoInstance(backRepo, gongdocstatusDB)
	}
	return
}

// BackRepoGongdocStatus.CheckoutPhaseTwoInstance Checkouts staged instances of GongdocStatus to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongdocstatusDB *GongdocStatusDB) (Error error) {

	gongdocstatus := (*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusPtr)[gongdocstatusDB.ID]
	_ = gongdocstatus // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitGongdocStatus allows commit of a single gongdocstatus (if already staged)
func (backRepo *BackRepoStruct) CommitGongdocStatus(gongdocstatus *models.GongdocStatus) {
	backRepo.BackRepoGongdocStatus.CommitPhaseOneInstance(gongdocstatus)
	if id, ok := (*backRepo.BackRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID)[gongdocstatus]; ok {
		backRepo.BackRepoGongdocStatus.CommitPhaseTwoInstance(backRepo, id, gongdocstatus)
	}
}

// CommitGongdocStatus allows checkout of a single gongdocstatus (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongdocStatus(gongdocstatus *models.GongdocStatus) {
	// check if the gongdocstatus is staged
	if _, ok := (*backRepo.BackRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID)[gongdocstatus]; ok {

		if id, ok := (*backRepo.BackRepoGongdocStatus.Map_GongdocStatusPtr_GongdocStatusDBID)[gongdocstatus]; ok {
			var gongdocstatusDB GongdocStatusDB
			gongdocstatusDB.ID = id

			if err := backRepo.BackRepoGongdocStatus.db.First(&gongdocstatusDB, id).Error; err != nil {
				log.Panicln("CheckoutGongdocStatus : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongdocStatus.CheckoutPhaseOneInstance(&gongdocstatusDB)
			backRepo.BackRepoGongdocStatus.CheckoutPhaseTwoInstance(backRepo, &gongdocstatusDB)
		}
	}
}

// CopyBasicFieldsFromGongdocStatus
func (gongdocstatusDB *GongdocStatusDB) CopyBasicFieldsFromGongdocStatus(gongdocstatus *models.GongdocStatus) {
	// insertion point for fields commit

	gongdocstatusDB.Name_Data.String = gongdocstatus.Name
	gongdocstatusDB.Name_Data.Valid = true

	gongdocstatusDB.Status_Data.String = gongdocstatus.Status.ToString()
	gongdocstatusDB.Status_Data.Valid = true

	gongdocstatusDB.CommandCompletionDate_Data.String = gongdocstatus.CommandCompletionDate
	gongdocstatusDB.CommandCompletionDate_Data.Valid = true
}

// CopyBasicFieldsFromGongdocStatusWOP
func (gongdocstatusDB *GongdocStatusDB) CopyBasicFieldsFromGongdocStatusWOP(gongdocstatus *GongdocStatusWOP) {
	// insertion point for fields commit

	gongdocstatusDB.Name_Data.String = gongdocstatus.Name
	gongdocstatusDB.Name_Data.Valid = true

	gongdocstatusDB.Status_Data.String = gongdocstatus.Status.ToString()
	gongdocstatusDB.Status_Data.Valid = true

	gongdocstatusDB.CommandCompletionDate_Data.String = gongdocstatus.CommandCompletionDate
	gongdocstatusDB.CommandCompletionDate_Data.Valid = true
}

// CopyBasicFieldsToGongdocStatus
func (gongdocstatusDB *GongdocStatusDB) CopyBasicFieldsToGongdocStatus(gongdocstatus *models.GongdocStatus) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongdocstatus.Name = gongdocstatusDB.Name_Data.String
	gongdocstatus.Status.FromString(gongdocstatusDB.Status_Data.String)
	gongdocstatus.CommandCompletionDate = gongdocstatusDB.CommandCompletionDate_Data.String
}

// CopyBasicFieldsToGongdocStatusWOP
func (gongdocstatusDB *GongdocStatusDB) CopyBasicFieldsToGongdocStatusWOP(gongdocstatus *GongdocStatusWOP) {
	gongdocstatus.ID = int(gongdocstatusDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongdocstatus.Name = gongdocstatusDB.Name_Data.String
	gongdocstatus.Status.FromString(gongdocstatusDB.Status_Data.String)
	gongdocstatus.CommandCompletionDate = gongdocstatusDB.CommandCompletionDate_Data.String
}

// Backup generates a json file from a slice of all GongdocStatusDB instances in the backrepo
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongdocStatusDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongdocStatusDB, 0)
	for _, gongdocstatusDB := range *backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB {
		forBackup = append(forBackup, gongdocstatusDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json GongdocStatus ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json GongdocStatus file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongdocStatusDB instances in the backrepo
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongdocStatusDB, 0)
	for _, gongdocstatusDB := range *backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB {
		forBackup = append(forBackup, gongdocstatusDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongdocStatus")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongdocStatus_Fields, -1)
	for _, gongdocstatusDB := range forBackup {

		var gongdocstatusWOP GongdocStatusWOP
		gongdocstatusDB.CopyBasicFieldsToGongdocStatusWOP(&gongdocstatusWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongdocstatusWOP, -1)
	}
}

// RestoreXL from the "GongdocStatus" sheet all GongdocStatusDB instances
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongdocStatusid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongdocStatus"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongdocStatus.rowVisitorGongdocStatus)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) rowVisitorGongdocStatus(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongdocstatusWOP GongdocStatusWOP
		row.ReadStruct(&gongdocstatusWOP)

		// add the unmarshalled struct to the stage
		gongdocstatusDB := new(GongdocStatusDB)
		gongdocstatusDB.CopyBasicFieldsFromGongdocStatusWOP(&gongdocstatusWOP)

		gongdocstatusDB_ID_atBackupTime := gongdocstatusDB.ID
		gongdocstatusDB.ID = 0
		query := backRepoGongdocStatus.db.Create(gongdocstatusDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB)[gongdocstatusDB.ID] = gongdocstatusDB
		BackRepoGongdocStatusid_atBckpTime_newID[gongdocstatusDB_ID_atBackupTime] = gongdocstatusDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongdocStatusDB.json" in dirPath that stores an array
// of GongdocStatusDB and stores it in the database
// the map BackRepoGongdocStatusid_atBckpTime_newID is updated accordingly
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongdocStatusid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongdocStatusDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json GongdocStatus file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongdocStatusDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongdocStatusDBID_GongdocStatusDB
	for _, gongdocstatusDB := range forRestore {

		gongdocstatusDB_ID_atBackupTime := gongdocstatusDB.ID
		gongdocstatusDB.ID = 0
		query := backRepoGongdocStatus.db.Create(gongdocstatusDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB)[gongdocstatusDB.ID] = gongdocstatusDB
		BackRepoGongdocStatusid_atBckpTime_newID[gongdocstatusDB_ID_atBackupTime] = gongdocstatusDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json GongdocStatus file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongdocStatus>id_atBckpTime_newID
// to compute new index
func (backRepoGongdocStatus *BackRepoGongdocStatusStruct) RestorePhaseTwo() {

	for _, gongdocstatusDB := range *backRepoGongdocStatus.Map_GongdocStatusDBID_GongdocStatusDB {

		// next line of code is to avert unused variable compilation error
		_ = gongdocstatusDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoGongdocStatus.db.Model(gongdocstatusDB).Updates(*gongdocstatusDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongdocStatusid_atBckpTime_newID map[uint]uint