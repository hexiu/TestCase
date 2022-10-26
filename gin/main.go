/*
 * @Date: 2021-10-18 15:46:02
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-10-18 16:05:21
 * @FilePath: /TestCase/gin/main.go
 */
package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/hi", hi)

}

// swagger:parameters Req
type Req struct {
	A json.RawMessage `json:"a"`
	B json.RawMessage `json:"b"`
}

// swagger:route Post /hi req Req
//
// Struct Req filtered by some parameters.
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//     Consumes:
//     - application/json
//     - application/x-protobuf
//
//     Produces:
//     - application/json
//     - application/x-protobuf
//
//     Schemes: http, https, ws, wss
//
//     Deprecated: true
//
//     Security:
//       api_key:
//       oauth: read, write
//
//     Responses:
//       default: genericError
//       200: someResponse
//       422: validationError
func hi(c *gin.Context) {
	req := &Req{}
	c.BindJSON(req)

	c.JSON(200, req)
}
