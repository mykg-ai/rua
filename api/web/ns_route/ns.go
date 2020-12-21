/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-11
 */
package ns_route

import (
	"github.com/gin-gonic/gin"
	"mykg.ai/rua/service/ns_service"
	"net/http"
)

func createNs(c * gin.Context) {
	ns, err := ns_service.CreateNs(c.PostForm("name"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": ns,
	})
}

func getAllNs(c * gin.Context) {
	//ns_service.
}

func Setup(e *gin.Engine) {
	g := e.Group("/ns")
	{
		g.POST("", createNs)
	}
}
