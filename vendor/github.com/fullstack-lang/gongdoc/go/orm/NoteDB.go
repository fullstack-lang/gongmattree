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
var dummy_Note_sql sql.NullBool
var dummy_Note_time time.Duration
var dummy_Note_sort sort.Float64Slice

// NoteAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model noteAPI
type NoteAPI struct {
	gorm.Model

	models.Note

	// encoding of pointers
	NotePointersEnconding
}

// NotePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type NotePointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field Classdiagram{}.Notes []*Note
	Classdiagram_NotesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Classdiagram_NotesDBID_Index sql.NullInt64
}

// NoteDB describes a note in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model noteDB
type NoteDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field noteDB.Name
	Name_Data sql.NullString

	// Declation for basic field noteDB.Body
	Body_Data sql.NullString

	// Declation for basic field noteDB.X
	X_Data sql.NullFloat64

	// Declation for basic field noteDB.Y
	Y_Data sql.NullFloat64

	// Declation for basic field noteDB.Width
	Width_Data sql.NullFloat64

	// Declation for basic field noteDB.Heigth
	Heigth_Data sql.NullFloat64

	// Declation for basic field noteDB.Matched
	// provide the sql storage for the boolan
	Matched_Data sql.NullBool
	// encoding of pointers
	NotePointersEnconding
}

// NoteDBs arrays noteDBs
// swagger:response noteDBsResponse
type NoteDBs []NoteDB

// NoteDBResponse provides response
// swagger:response noteDBResponse
type NoteDBResponse struct {
	NoteDB
}

// NoteWOP is a Note without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type NoteWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Body string `xlsx:"2"`

	X float64 `xlsx:"3"`

	Y float64 `xlsx:"4"`

	Width float64 `xlsx:"5"`

	Heigth float64 `xlsx:"6"`

	Matched bool `xlsx:"7"`
	// insertion for WOP pointer fields
}

var Note_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Body",
	"X",
	"Y",
	"Width",
	"Heigth",
	"Matched",
}

type BackRepoNoteStruct struct {
	// stores NoteDB according to their gorm ID
	Map_NoteDBID_NoteDB *map[uint]*NoteDB

	// stores NoteDB ID according to Note address
	Map_NotePtr_NoteDBID *map[*models.Note]uint

	// stores Note according to their gorm ID
	Map_NoteDBID_NotePtr *map[uint]*models.Note

	db *gorm.DB
}

func (backRepoNote *BackRepoNoteStruct) GetDB() *gorm.DB {
	return backRepoNote.db
}

// GetNoteDBFromNotePtr is a handy function to access the back repo instance from the stage instance
func (backRepoNote *BackRepoNoteStruct) GetNoteDBFromNotePtr(note *models.Note) (noteDB *NoteDB) {
	id := (*backRepoNote.Map_NotePtr_NoteDBID)[note]
	noteDB = (*backRepoNote.Map_NoteDBID_NoteDB)[id]
	return
}

// BackRepoNote.Init set up the BackRepo of the Note
func (backRepoNote *BackRepoNoteStruct) Init(db *gorm.DB) (Error error) {

	if backRepoNote.Map_NoteDBID_NotePtr != nil {
		err := errors.New("In Init, backRepoNote.Map_NoteDBID_NotePtr should be nil")
		return err
	}

	if backRepoNote.Map_NoteDBID_NoteDB != nil {
		err := errors.New("In Init, backRepoNote.Map_NoteDBID_NoteDB should be nil")
		return err
	}

	if backRepoNote.Map_NotePtr_NoteDBID != nil {
		err := errors.New("In Init, backRepoNote.Map_NotePtr_NoteDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Note, 0)
	backRepoNote.Map_NoteDBID_NotePtr = &tmp

	tmpDB := make(map[uint]*NoteDB, 0)
	backRepoNote.Map_NoteDBID_NoteDB = &tmpDB

	tmpID := make(map[*models.Note]uint, 0)
	backRepoNote.Map_NotePtr_NoteDBID = &tmpID

	backRepoNote.db = db
	return
}

// BackRepoNote.CommitPhaseOne commits all staged instances of Note to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoNote *BackRepoNoteStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for note := range stage.Notes {
		backRepoNote.CommitPhaseOneInstance(note)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, note := range *backRepoNote.Map_NoteDBID_NotePtr {
		if _, ok := stage.Notes[note]; !ok {
			backRepoNote.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoNote.CommitDeleteInstance commits deletion of Note to the BackRepo
func (backRepoNote *BackRepoNoteStruct) CommitDeleteInstance(id uint) (Error error) {

	note := (*backRepoNote.Map_NoteDBID_NotePtr)[id]

	// note is not staged anymore, remove noteDB
	noteDB := (*backRepoNote.Map_NoteDBID_NoteDB)[id]
	query := backRepoNote.db.Unscoped().Delete(&noteDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoNote.Map_NotePtr_NoteDBID), note)
	delete((*backRepoNote.Map_NoteDBID_NotePtr), id)
	delete((*backRepoNote.Map_NoteDBID_NoteDB), id)

	return
}

// BackRepoNote.CommitPhaseOneInstance commits note staged instances of Note to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoNote *BackRepoNoteStruct) CommitPhaseOneInstance(note *models.Note) (Error error) {

	// check if the note is not commited yet
	if _, ok := (*backRepoNote.Map_NotePtr_NoteDBID)[note]; ok {
		return
	}

	// initiate note
	var noteDB NoteDB
	noteDB.CopyBasicFieldsFromNote(note)

	query := backRepoNote.db.Create(&noteDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoNote.Map_NotePtr_NoteDBID)[note] = noteDB.ID
	(*backRepoNote.Map_NoteDBID_NotePtr)[noteDB.ID] = note
	(*backRepoNote.Map_NoteDBID_NoteDB)[noteDB.ID] = &noteDB

	return
}

// BackRepoNote.CommitPhaseTwo commits all staged instances of Note to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNote *BackRepoNoteStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, note := range *backRepoNote.Map_NoteDBID_NotePtr {
		backRepoNote.CommitPhaseTwoInstance(backRepo, idx, note)
	}

	return
}

// BackRepoNote.CommitPhaseTwoInstance commits {{structname }} of models.Note to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNote *BackRepoNoteStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, note *models.Note) (Error error) {

	// fetch matching noteDB
	if noteDB, ok := (*backRepoNote.Map_NoteDBID_NoteDB)[idx]; ok {

		noteDB.CopyBasicFieldsFromNote(note)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoNote.db.Save(&noteDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Note intance %s", note.Name))
		return err
	}

	return
}

// BackRepoNote.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoNote *BackRepoNoteStruct) CheckoutPhaseOne() (Error error) {

	noteDBArray := make([]NoteDB, 0)
	query := backRepoNote.db.Find(&noteDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	noteInstancesToBeRemovedFromTheStage := make(map[*models.Note]any)
	for key, value := range models.Stage.Notes {
		noteInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, noteDB := range noteDBArray {
		backRepoNote.CheckoutPhaseOneInstance(&noteDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		note, ok := (*backRepoNote.Map_NoteDBID_NotePtr)[noteDB.ID]
		if ok {
			delete(noteInstancesToBeRemovedFromTheStage, note)
		}
	}

	// remove from stage and back repo's 3 maps all notes that are not in the checkout
	for note := range noteInstancesToBeRemovedFromTheStage {
		note.Unstage()

		// remove instance from the back repo 3 maps
		noteID := (*backRepoNote.Map_NotePtr_NoteDBID)[note]
		delete((*backRepoNote.Map_NotePtr_NoteDBID), note)
		delete((*backRepoNote.Map_NoteDBID_NoteDB), noteID)
		delete((*backRepoNote.Map_NoteDBID_NotePtr), noteID)
	}

	return
}

// CheckoutPhaseOneInstance takes a noteDB that has been found in the DB, updates the backRepo and stages the
// models version of the noteDB
func (backRepoNote *BackRepoNoteStruct) CheckoutPhaseOneInstance(noteDB *NoteDB) (Error error) {

	note, ok := (*backRepoNote.Map_NoteDBID_NotePtr)[noteDB.ID]
	if !ok {
		note = new(models.Note)

		(*backRepoNote.Map_NoteDBID_NotePtr)[noteDB.ID] = note
		(*backRepoNote.Map_NotePtr_NoteDBID)[note] = noteDB.ID

		// append model store with the new element
		note.Name = noteDB.Name_Data.String
		note.Stage()
	}
	noteDB.CopyBasicFieldsToNote(note)

	// preserve pointer to noteDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_NoteDBID_NoteDB)[noteDB hold variable pointers
	noteDB_Data := *noteDB
	preservedPtrToNote := &noteDB_Data
	(*backRepoNote.Map_NoteDBID_NoteDB)[noteDB.ID] = preservedPtrToNote

	return
}

// BackRepoNote.CheckoutPhaseTwo Checkouts all staged instances of Note to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNote *BackRepoNoteStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, noteDB := range *backRepoNote.Map_NoteDBID_NoteDB {
		backRepoNote.CheckoutPhaseTwoInstance(backRepo, noteDB)
	}
	return
}

// BackRepoNote.CheckoutPhaseTwoInstance Checkouts staged instances of Note to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNote *BackRepoNoteStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, noteDB *NoteDB) (Error error) {

	note := (*backRepoNote.Map_NoteDBID_NotePtr)[noteDB.ID]
	_ = note // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitNote allows commit of a single note (if already staged)
func (backRepo *BackRepoStruct) CommitNote(note *models.Note) {
	backRepo.BackRepoNote.CommitPhaseOneInstance(note)
	if id, ok := (*backRepo.BackRepoNote.Map_NotePtr_NoteDBID)[note]; ok {
		backRepo.BackRepoNote.CommitPhaseTwoInstance(backRepo, id, note)
	}
}

// CommitNote allows checkout of a single note (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutNote(note *models.Note) {
	// check if the note is staged
	if _, ok := (*backRepo.BackRepoNote.Map_NotePtr_NoteDBID)[note]; ok {

		if id, ok := (*backRepo.BackRepoNote.Map_NotePtr_NoteDBID)[note]; ok {
			var noteDB NoteDB
			noteDB.ID = id

			if err := backRepo.BackRepoNote.db.First(&noteDB, id).Error; err != nil {
				log.Panicln("CheckoutNote : Problem with getting object with id:", id)
			}
			backRepo.BackRepoNote.CheckoutPhaseOneInstance(&noteDB)
			backRepo.BackRepoNote.CheckoutPhaseTwoInstance(backRepo, &noteDB)
		}
	}
}

// CopyBasicFieldsFromNote
func (noteDB *NoteDB) CopyBasicFieldsFromNote(note *models.Note) {
	// insertion point for fields commit

	noteDB.Name_Data.String = note.Name
	noteDB.Name_Data.Valid = true

	noteDB.Body_Data.String = note.Body
	noteDB.Body_Data.Valid = true

	noteDB.X_Data.Float64 = note.X
	noteDB.X_Data.Valid = true

	noteDB.Y_Data.Float64 = note.Y
	noteDB.Y_Data.Valid = true

	noteDB.Width_Data.Float64 = note.Width
	noteDB.Width_Data.Valid = true

	noteDB.Heigth_Data.Float64 = note.Heigth
	noteDB.Heigth_Data.Valid = true

	noteDB.Matched_Data.Bool = note.Matched
	noteDB.Matched_Data.Valid = true
}

// CopyBasicFieldsFromNoteWOP
func (noteDB *NoteDB) CopyBasicFieldsFromNoteWOP(note *NoteWOP) {
	// insertion point for fields commit

	noteDB.Name_Data.String = note.Name
	noteDB.Name_Data.Valid = true

	noteDB.Body_Data.String = note.Body
	noteDB.Body_Data.Valid = true

	noteDB.X_Data.Float64 = note.X
	noteDB.X_Data.Valid = true

	noteDB.Y_Data.Float64 = note.Y
	noteDB.Y_Data.Valid = true

	noteDB.Width_Data.Float64 = note.Width
	noteDB.Width_Data.Valid = true

	noteDB.Heigth_Data.Float64 = note.Heigth
	noteDB.Heigth_Data.Valid = true

	noteDB.Matched_Data.Bool = note.Matched
	noteDB.Matched_Data.Valid = true
}

// CopyBasicFieldsToNote
func (noteDB *NoteDB) CopyBasicFieldsToNote(note *models.Note) {
	// insertion point for checkout of basic fields (back repo to stage)
	note.Name = noteDB.Name_Data.String
	note.Body = noteDB.Body_Data.String
	note.X = noteDB.X_Data.Float64
	note.Y = noteDB.Y_Data.Float64
	note.Width = noteDB.Width_Data.Float64
	note.Heigth = noteDB.Heigth_Data.Float64
	note.Matched = noteDB.Matched_Data.Bool
}

// CopyBasicFieldsToNoteWOP
func (noteDB *NoteDB) CopyBasicFieldsToNoteWOP(note *NoteWOP) {
	note.ID = int(noteDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	note.Name = noteDB.Name_Data.String
	note.Body = noteDB.Body_Data.String
	note.X = noteDB.X_Data.Float64
	note.Y = noteDB.Y_Data.Float64
	note.Width = noteDB.Width_Data.Float64
	note.Heigth = noteDB.Heigth_Data.Float64
	note.Matched = noteDB.Matched_Data.Bool
}

// Backup generates a json file from a slice of all NoteDB instances in the backrepo
func (backRepoNote *BackRepoNoteStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "NoteDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*NoteDB, 0)
	for _, noteDB := range *backRepoNote.Map_NoteDBID_NoteDB {
		forBackup = append(forBackup, noteDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Note ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Note file", err.Error())
	}
}

// Backup generates a json file from a slice of all NoteDB instances in the backrepo
func (backRepoNote *BackRepoNoteStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*NoteDB, 0)
	for _, noteDB := range *backRepoNote.Map_NoteDBID_NoteDB {
		forBackup = append(forBackup, noteDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Note")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Note_Fields, -1)
	for _, noteDB := range forBackup {

		var noteWOP NoteWOP
		noteDB.CopyBasicFieldsToNoteWOP(&noteWOP)

		row := sh.AddRow()
		row.WriteStruct(&noteWOP, -1)
	}
}

// RestoreXL from the "Note" sheet all NoteDB instances
func (backRepoNote *BackRepoNoteStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoNoteid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Note"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoNote.rowVisitorNote)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoNote *BackRepoNoteStruct) rowVisitorNote(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var noteWOP NoteWOP
		row.ReadStruct(&noteWOP)

		// add the unmarshalled struct to the stage
		noteDB := new(NoteDB)
		noteDB.CopyBasicFieldsFromNoteWOP(&noteWOP)

		noteDB_ID_atBackupTime := noteDB.ID
		noteDB.ID = 0
		query := backRepoNote.db.Create(noteDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoNote.Map_NoteDBID_NoteDB)[noteDB.ID] = noteDB
		BackRepoNoteid_atBckpTime_newID[noteDB_ID_atBackupTime] = noteDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "NoteDB.json" in dirPath that stores an array
// of NoteDB and stores it in the database
// the map BackRepoNoteid_atBckpTime_newID is updated accordingly
func (backRepoNote *BackRepoNoteStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoNoteid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "NoteDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Note file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*NoteDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_NoteDBID_NoteDB
	for _, noteDB := range forRestore {

		noteDB_ID_atBackupTime := noteDB.ID
		noteDB.ID = 0
		query := backRepoNote.db.Create(noteDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoNote.Map_NoteDBID_NoteDB)[noteDB.ID] = noteDB
		BackRepoNoteid_atBckpTime_newID[noteDB_ID_atBackupTime] = noteDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Note file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Note>id_atBckpTime_newID
// to compute new index
func (backRepoNote *BackRepoNoteStruct) RestorePhaseTwo() {

	for _, noteDB := range *backRepoNote.Map_NoteDBID_NoteDB {

		// next line of code is to avert unused variable compilation error
		_ = noteDB

		// insertion point for reindexing pointers encoding
		// This reindex note.Notes
		if noteDB.Classdiagram_NotesDBID.Int64 != 0 {
			noteDB.Classdiagram_NotesDBID.Int64 =
				int64(BackRepoClassdiagramid_atBckpTime_newID[uint(noteDB.Classdiagram_NotesDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoNote.db.Model(noteDB).Updates(*noteDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoNoteid_atBckpTime_newID map[uint]uint
