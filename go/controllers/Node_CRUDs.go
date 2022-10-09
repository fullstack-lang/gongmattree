// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongmattree/go/models"
	"github.com/fullstack-lang/gongmattree/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Node__dummysDeclaration__ models.Node
var __Node_time__dummyDeclaration time.Duration

// An NodeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNode updateNode deleteNode
type NodeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NodeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNode updateNode
type NodeInput struct {
	// The Node to submit or modify
	// in: body
	Node *orm.NodeAPI
}

// GetNodes
//
// swagger:route GET /nodes nodes getNodes
//
// Get all nodes
//
// Responses:
//    default: genericError
//        200: nodeDBsResponse
func GetNodes(c *gin.Context) {
	db := orm.BackRepo.BackRepoNode.GetDB()

	// source slice
	var nodeDBs []orm.NodeDB
	query := db.Find(&nodeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	nodeAPIs := make([]orm.NodeAPI, 0)

	// for each node, update fields from the database nullable fields
	for idx := range nodeDBs {
		nodeDB := &nodeDBs[idx]
		_ = nodeDB
		var nodeAPI orm.NodeAPI

		// insertion point for updating fields
		nodeAPI.ID = nodeDB.ID
		nodeDB.CopyBasicFieldsToNode(&nodeAPI.Node)
		nodeAPI.NodePointersEnconding = nodeDB.NodePointersEnconding
		nodeAPIs = append(nodeAPIs, nodeAPI)
	}

	c.JSON(http.StatusOK, nodeAPIs)
}

// PostNode
//
// swagger:route POST /nodes nodes postNode
//
// Creates a node
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: nodeDBResponse
func PostNode(c *gin.Context) {
	db := orm.BackRepo.BackRepoNode.GetDB()

	// Validate input
	var input orm.NodeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create node
	nodeDB := orm.NodeDB{}
	nodeDB.NodePointersEnconding = input.NodePointersEnconding
	nodeDB.CopyBasicFieldsFromNode(&input.Node)

	query := db.Create(&nodeDB)
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

	c.JSON(http.StatusOK, nodeDB)
}

// GetNode
//
// swagger:route GET /nodes/{ID} nodes getNode
//
// Gets the details for a node.
//
// Responses:
//    default: genericError
//        200: nodeDBResponse
func GetNode(c *gin.Context) {
	db := orm.BackRepo.BackRepoNode.GetDB()

	// Get nodeDB in DB
	var nodeDB orm.NodeDB
	if err := db.First(&nodeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var nodeAPI orm.NodeAPI
	nodeAPI.ID = nodeDB.ID
	nodeAPI.NodePointersEnconding = nodeDB.NodePointersEnconding
	nodeDB.CopyBasicFieldsToNode(&nodeAPI.Node)

	c.JSON(http.StatusOK, nodeAPI)
}

// UpdateNode
//
// swagger:route PATCH /nodes/{ID} nodes updateNode
//
// Update a node
//
// Responses:
//    default: genericError
//        200: nodeDBResponse
func UpdateNode(c *gin.Context) {
	db := orm.BackRepo.BackRepoNode.GetDB()

	// Get model if exist
	var nodeDB orm.NodeDB

	// fetch the node
	query := db.First(&nodeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.NodeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	nodeDB.CopyBasicFieldsFromNode(&input.Node)
	nodeDB.NodePointersEnconding = input.NodePointersEnconding

	query = db.Model(&nodeDB).Updates(nodeDB)
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

	// return status OK with the marshalling of the the nodeDB
	c.JSON(http.StatusOK, nodeDB)
}

// DeleteNode
//
// swagger:route DELETE /nodes/{ID} nodes deleteNode
//
// Delete a node
//
// Responses:
//    default: genericError
func DeleteNode(c *gin.Context) {
	db := orm.BackRepo.BackRepoNode.GetDB()

	// Get model if exist
	var nodeDB orm.NodeDB
	if err := db.First(&nodeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&nodeDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
