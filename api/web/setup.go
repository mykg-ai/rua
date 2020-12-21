/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-11
 */
package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mykg.ai/rua/config"
	"mykg.ai/rua/web/ns_route"
	"net/http"
)

func Setup() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "healthy!")
	})

	ns_route.Setup(r)

	if config.Env != "local" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	_ = r.Run(fmt.Sprintf(":%s", config.Server.HttpPort))
}