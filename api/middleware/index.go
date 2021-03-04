package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mykg.ai/rua/config"
	"mykg.ai/rua/domain/entity"
	R "mykg.ai/rua/err/r"
	Recode "mykg.ai/rua/err/recode"
	"mykg.ai/rua/utils"
	"strings"
)

func JWTAuthenticatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 提取 JWT
		jwtStr := c.Request.Header.Get("Authorization")
		// 校验 JWT
		if jwtStr == "" {
			R.Unauthorized(c, Recode.TOKEN_NOT_EXIST.Code, Recode.TOKEN_NOT_EXIST.Msg)
			c.Abort()
			return
		}
		checkToken := strings.Split(jwtStr, " ")
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			R.Unauthorized(c, Recode.TOKEN_TYPE_ERROR.Code, Recode.TOKEN_TYPE_ERROR.Msg)
			c.Abort()
			return
		}
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(checkToken[1], claims, func(*jwt.Token) (interface{}, error) {
			return []byte(utils.SecretKey), nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					R.Unauthorized(c, Recode.TOKEN_EXPIRE.Code, Recode.TOKEN_EXPIRE.Msg)
					c.Abort()
					return
				} else {
					R.Unauthorized(c, Recode.TOKEN_ERROR.Code, Recode.TOKEN_ERROR.Msg)
					c.Abort()
					return
				}
			}
		}
		user := entity.User{}
		config.DB.Where("username = ?", claims["username"].(string)).Omit("password").First(&user)
		if user.ID == 0 {
			R.Unauthorized(c, Recode.NOT_FOUND.Code, fmt.Sprintf("%s : %s", Recode.NOT_FOUND, claims["username"].(string)))
			c.Abort()
			return
		}
		c.Set("user", user)
	}
}

type AuthAPIFunc func(c *gin.Context, user *entity.User)

func Authenticator(f AuthAPIFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, exists := c.Get("user")
		if !exists {
			R.Unauthorized(c, Recode.TOKEN_NOT_EXIST.Code, Recode.TOKEN_NOT_EXIST.Msg)
			c.Abort()
			return
		}
		var user entity.User
		user = u.(entity.User)
		f(c, &user)
	}
}
