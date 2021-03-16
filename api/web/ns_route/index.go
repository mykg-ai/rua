/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-11
 */
package ns_route

import (
	"github.com/gin-gonic/gin"
	"mykg.ai/rua/domain/entity"
	"mykg.ai/rua/domain/enum"
	ERR "mykg.ai/rua/err"
	R "mykg.ai/rua/err/r"
	Recode "mykg.ai/rua/err/recode"
	"mykg.ai/rua/middleware"
	"mykg.ai/rua/service/ns_service"
	"strconv"
)

func createNamespace() gin.HandlerFunc {
	return middleware.Authenticator(func(c *gin.Context, user *entity.User) {
		if user.Identity == enum.ADMIN {
			var namespace entity.Namespace
			namespace.Name = c.PostForm("name")
			namespace.Creator = user.ID
			ns, err := ns_service.CreateNamespace(&namespace)
			if err != nil {
				ERR.Panic(Recode.NAMESPACE_ALREADY_EXIST, "")
			}
			R.Ok(c, ns)
		} else {
			R.Unauthorized(c, Recode.USER_UNAUTHORIZED.Code, Recode.USER_UNAUTHORIZED.Msg)
		}
	})
}

func findNamespaces() gin.HandlerFunc {
	return middleware.Authenticator(func(c *gin.Context, user *entity.User) {
		if user.Identity == enum.ADMIN {
			namespaces, err := ns_service.FindNamespaces()
			if err != nil {
				ERR.Panic(Recode.ERROR, err.Error())
			}
			R.Ok(c, namespaces)
		} else {
			R.Unauthorized(c, Recode.USER_UNAUTHORIZED.Code, Recode.USER_UNAUTHORIZED.Msg)
		}
	})
}

func findNamespace(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	namespace, err := ns_service.FindNamespace(uint64(id))
	if err != nil {
		ERR.Panic(Recode.NOT_FOUND, "namespace")
	}
	R.Ok(c, namespace)
}

func deleteNamespace() gin.HandlerFunc {
	return middleware.Authenticator(func(c *gin.Context, user *entity.User) {
		if user.Identity == enum.ADMIN {
			id, _ := strconv.Atoi(c.Param("id"))
			err := ns_service.DeleteNamespace(uint64(id))
			if err != nil {
				ERR.Panic(Recode.ERROR, err.Error())
			}
		} else {
			R.Unauthorized(c, Recode.USER_UNAUTHORIZED.Code, Recode.USER_UNAUTHORIZED.Msg)
		}
	})
}

func Setup(e *gin.Engine) {
	g := e.Group("/namespace", middleware.JWTAuthenticatorMiddleware())
	{
		g.POST("/", createNamespace())
		g.GET("/", findNamespaces())
		g.GET("/:id", findNamespace)
		g.DELETE("/:id", deleteNamespace())
	}
}
