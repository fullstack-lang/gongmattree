// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __PointerToGongStructField__dummysDeclaration__ models.PointerToGongStructField
var __PointerToGongStructField_time__dummyDeclaration time.Duration

// An PointerToGongStructFieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPointerToGongStructField updatePointerToGongStructField deletePointerToGongStructField
type PointerToGongStructFieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PointerToGongStructFieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPointerToGongStructField updatePointerToGongStructField
type PointerToGongStructFieldInput struct {
	// The PointerToGongStructField to submit or modify
	// in: body
	PointerToGongStructField *orm.PointerToGongStructFieldAPI
}

// GetPointerToGongStructFields
//
// swagger:route GET /pointertogongstructfields pointertogongstructfields getPointerToGongStructFields
//
// Get all pointertogongstructfields
//
// Responses:
//    default: genericError
//        200: pointertogongstructfieldDBsResponse
func GetPointerToGongStructFields(c *gin.Context) {
	db := orm.BackRepo.BackRepoPointerToGongStructField.GetDB()

	// source slice
	var pointertogongstructfieldDBs []orm.PointerToGongStructFieldDB
	query := db.Find(&pointertogongstructfieldDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pointertogongstructfieldAPIs := make([]orm.PointerToGongStructFieldAPI, 0)

	// for each pointertogongstructfield, update fields from the database nullable fields
	for idx := range pointertogongstructfieldDBs {
		pointertogongstructfieldDB := &pointertogongstructfieldDBs[idx]
		_ = pointertogongstructfieldDB
		var pointertogongstructfieldAPI orm.PointerToGongStructFieldAPI

		// insertion point for updating fields
		pointertogongstructfieldAPI.ID = pointertogongstructfieldDB.ID
		pointertogongstructfieldDB.CopyBasicFieldsToPointerToGongStructField(&pointertogongstructfieldAPI.PointerToGongStructField)
		pointertogongstructfieldAPI.PointerToGongStructFieldPointersEnconding = pointertogongstructfieldDB.PointerToGongStructFieldPointersEnconding
		pointertogongstructfieldAPIs = append(pointertogongstructfieldAPIs, pointertogongstructfieldAPI)
	}

	c.JSON(http.StatusOK, pointertogongstructfieldAPIs)
}

// PostPointerToGongStructField
//
// swagger:route POST /pointertogongstructfields pointertogongstructfields postPointerToGongStructField
//
// Creates a pointertogongstructfield
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: pointertogongstructfieldDBResponse
func PostPointerToGongStructField(c *gin.Context) {
	db := orm.BackRepo.BackRepoPointerToGongStructField.GetDB()

	// Validate input
	var input orm.PointerToGongStructFieldAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create pointertogongstructfield
	pointertogongstructfieldDB := orm.PointerToGongStructFieldDB{}
	pointertogongstructfieldDB.PointerToGongStructFieldPointersEnconding = input.PointerToGongStructFieldPointersEnconding
	pointertogongstructfieldDB.CopyBasicFieldsFromPointerToGongStructField(&input.PointerToGongStructField)

	query := db.Create(&pointertogongstructfieldDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pointertogongstructfieldDB)
}

// GetPointerToGongStructField
//
// swagger:route GET /pointertogongstructfields/{ID} pointertogongstructfields getPointerToGongStructField
//
// Gets the details for a pointertogongstructfield.
//
// Responses:
//    default: genericError
//        200: pointertogongstructfieldDBResponse
func GetPointerToGongStructField(c *gin.Context) {
	db := orm.BackRepo.BackRepoPointerToGongStructField.GetDB()

	// Get pointertogongstructfieldDB in DB
	var pointertogongstructfieldDB orm.PointerToGongStructFieldDB
	if err := db.First(&pointertogongstructfieldDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pointertogongstructfieldAPI orm.PointerToGongStructFieldAPI
	pointertogongstructfieldAPI.ID = pointertogongstructfieldDB.ID
	pointertogongstructfieldAPI.PointerToGongStructFieldPointersEnconding = pointertogongstructfieldDB.PointerToGongStructFieldPointersEnconding
	pointertogongstructfieldDB.CopyBasicFieldsToPointerToGongStructField(&pointertogongstructfieldAPI.PointerToGongStructField)

	c.JSON(http.StatusOK, pointertogongstructfieldAPI)
}

// UpdatePointerToGongStructField
//
// swagger:route PATCH /pointertogongstructfields/{ID} pointertogongstructfields updatePointerToGongStructField
//
// Update a pointertogongstructfield
//
// Responses:
//    default: genericError
//        200: pointertogongstructfieldDBResponse
func UpdatePointerToGongStructField(c *gin.Context) {
	db := orm.BackRepo.BackRepoPointerToGongStructField.GetDB()

	// Get model if exist
	var pointertogongstructfieldDB orm.PointerToGongStructFieldDB

	// fetch the pointertogongstructfield
	query := db.First(&pointertogongstructfieldDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.PointerToGongStructFieldAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	pointertogongstructfieldDB.CopyBasicFieldsFromPointerToGongStructField(&input.PointerToGongStructField)
	pointertogongstructfieldDB.PointerToGongStructFieldPointersEnconding = input.PointerToGongStructFieldPointersEnconding

	query = db.Model(&pointertogongstructfieldDB).Updates(pointertogongstructfieldDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pointertogongstructfieldDB
	c.JSON(http.StatusOK, pointertogongstructfieldDB)
}

// DeletePointerToGongStructField
//
// swagger:route DELETE /pointertogongstructfields/{ID} pointertogongstructfields deletePointerToGongStructField
//
// Delete a pointertogongstructfield
//
// Responses:
//    default: genericError
func DeletePointerToGongStructField(c *gin.Context) {
	db := orm.BackRepo.BackRepoPointerToGongStructField.GetDB()

	// Get model if exist
	var pointertogongstructfieldDB orm.PointerToGongStructFieldDB
	if err := db.First(&pointertogongstructfieldDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&pointertogongstructfieldDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
