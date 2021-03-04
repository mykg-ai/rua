/**
 * Author: Honghui <honghui@mykg.ai> 2021-03-01
 */
package user_route

import (
	"github.com/gin-gonic/gin"
	"mykg.ai/rua/domain/entity"
	"mykg.ai/rua/domain/enum"
	"mykg.ai/rua/err"
	R "mykg.ai/rua/err/r"
	Recode "mykg.ai/rua/err/recode"
	"mykg.ai/rua/middleware"
	"mykg.ai/rua/service/user_service"
	"strconv"
)

func login(c *gin.Context) {
	token := user_service.Login(c.PostForm("username"), c.PostForm("password"))
	R.Ok(c, gin.H{
		"token": token,
	})
}

func createUser() gin.HandlerFunc {
	return middleware.Authenticator(func(c *gin.Context, user *entity.User) {
		if user.Identity == enum.ADMIN {
			var user entity.User
			_ = c.ShouldBindJSON(&user)
			user, err := user_service.CreateUser(&user)
			if err != nil {
				ERR.Panic(Recode.USERNAME_ALREADY_EXIST, "")
			}
			R.Ok(c, user)
		} else {
			R.Unauthorized(c, Recode.USER_UNAUTHORIZED.Code, Recode.USER_UNAUTHORIZED.Msg)
		}
	})
}

func findUsers() gin.HandlerFunc {
	return middleware.Authenticator(func(c *gin.Context, user *entity.User) {
		if user.Identity == enum.ADMIN {
			users, err := user_service.FindUsers()
			if err != nil {
				ERR.Panic(Recode.ERROR, err.Error())
			}
			R.Ok(c, users)
		} else {
			R.Unauthorized(c, Recode.USER_UNAUTHORIZED.Code, Recode.USER_UNAUTHORIZED.Msg)
		}
	})
}

func findUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := user_service.FindUser(uint64(id))
	if err != nil {
		ERR.Panic(Recode.NOT_FOUND, "user")
	}
	R.Ok(c, user)
}

func deleteUser() gin.HandlerFunc {
	return middleware.Authenticator(func(c *gin.Context, user *entity.User) {
		if user.Identity == enum.ADMIN {
			id, _ := strconv.Atoi(c.Param("id"))
			err := user_service.DeleteUser(uint64(id))
			if err != nil {
				ERR.Panic(Recode.ERROR, err.Error())
			}
		} else {
			R.Unauthorized(c, Recode.USER_UNAUTHORIZED.Code, Recode.USER_UNAUTHORIZED.Msg)
		}
	})
}

func updatePassword(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := user_service.UpdatePassword(uint64(id), c.PostForm("password"))
	if err != nil {
		ERR.Panic(Recode.ERROR, err.Error())
	}
}

func Setup(e *gin.Engine) {
	g := e.Group("/user")
	{
		g.POST("/login", login)
		g.POST("/", middleware.JWTAuthenticatorMiddleware(), createUser())
		g.GET("/", middleware.JWTAuthenticatorMiddleware(), findUsers())
		g.GET("/:id", middleware.JWTAuthenticatorMiddleware(), findUser)
		g.DELETE("/:id", middleware.JWTAuthenticatorMiddleware(), deleteUser())
		g.PATCH("/:id/password", middleware.JWTAuthenticatorMiddleware(), updatePassword)
	}
}
