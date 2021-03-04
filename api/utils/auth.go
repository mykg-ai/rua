package utils

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"mykg.ai/rua/config"
	"time"
)

var SecretKey = config.Config.JWT.SecretKey

// 加密密码
func EncryptPassword(password string) string {
	const cost = 10
	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(encryptPassword)
}

// 生成JWT
func GenerateJWT(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// 使用一个密钥字符串对 Token 进行签名
	t, _ := token.SignedString([]byte(SecretKey))
	return t
}
