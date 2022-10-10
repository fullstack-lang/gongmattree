package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongmattree/go/orm"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongmattree/go")
	{ // insertion point for registrations
		v1.GET("/v1/nodes", GetNodes)
		v1.GET("/v1/nodes/:id", GetNode)
		v1.POST("/v1/nodes", PostNode)
		v1.PATCH("/v1/nodes/:id", UpdateNode)
		v1.PUT("/v1/nodes/:id", UpdateNode)
		v1.DELETE("/v1/nodes/:id", DeleteNode)

		v1.GET("/v1/trees", GetTrees)
		v1.GET("/v1/trees/:id", GetTree)
		v1.POST("/v1/trees", PostTree)
		v1.PATCH("/v1/trees/:id", UpdateTree)
		v1.PUT("/v1/trees/:id", UpdateTree)
		v1.DELETE("/v1/trees/:id", DeleteTree)

		v1.GET("/commitfrombacknb", GetLastCommitFromBackNb)
		v1.GET("/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func GetLastCommitFromBackNb(c *gin.Context) {
	res := orm.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
