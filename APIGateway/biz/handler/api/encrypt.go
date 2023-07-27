// Code generated by hertz generator.

package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	//api "hello/biz/model/api"
)

// Encrypt .
// @router /encrypt [POST]
func Encrypt(ctx context.Context, c *app.RequestContext) {
	var requestURL string = "http://example.com/life/client/11?vint64=1&items=item0,item1,item2"
	var IDLPATH string = "idl/encrypt.thrift"
	var jsonData map[string]interface{}

	response := c.GetRawData()

	err := json.Unmarshal(response, &jsonData)

	if err != nil {
		fmt.Println("Error", err)
		c.String(consts.StatusBadRequest, "post request fail due to wrong type")
		return
	}

	fmt.Println(jsonData)

	responseFromRPC, err := thriftsend(IDLPATH, "encrypt", jsonData, requestURL, ctx)

	if err != nil {
		fmt.Println(err)
		c.String(consts.StatusBadRequest, "thrift call failed")
		return
	}

	fmt.Println("Success")

	c.JSON(consts.StatusOK, responseFromRPC)
}
