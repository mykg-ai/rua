/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-11
 */
package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mykg.ai/rua/config"
	ERR "mykg.ai/rua/err"
	R "mykg.ai/rua/err/r"
	"mykg.ai/rua/web/ns_route"
	"mykg.ai/rua/web/user_route"
)

func Setup() {
	r := gin.Default()
	r.Use(ERR.Recover)

	r.GET("/", func(c *gin.Context) {
		R.Ok(c, nil)
	})

	ns_route.Setup(r)
	user_route.Setup(r)

	if config.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	_ = r.Run(fmt.Sprintf(":%s", config.Config.Server.Port))
}
