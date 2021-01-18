/**
 * Author: Goddy <goddy@mykg.ai> 2021-01-15
 */
package R

import (
	"github.com/gin-gonic/gin"
	Recode "mykg.ai/rua/err/recode"
	"net/http"
)

type R struct {
	Data interface{} `json:"data,omitempty"`
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, R{
		Data: data,
		Code: Recode.SUCCESS.Code,
		Msg:  Recode.SUCCESS.Msg,
	})
}

func Error(c *gin.Context, code uint, msg string) {
	c.JSON(http.StatusOK, R{
		Data: nil,
		Code: code,
		Msg:  msg,
	})
}
