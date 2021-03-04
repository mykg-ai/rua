/**
 * Author: Honghui <honghui@mykg.ai> 2021-03-01
 */
package user_service

import (
	"golang.org/x/crypto/bcrypt"
	"mykg.ai/rua/config"
	"mykg.ai/rua/domain/entity"
	ERR "mykg.ai/rua/err"
	Recode "mykg.ai/rua/err/recode"
	"mykg.ai/rua/utils"
)

func Login(username string, password string) string {
	if username == "" || password == "" {
		ERR.Panic(Recode.USERNAME_OR_PASSWORD_IS_EMPTY, "")
	}
	user := entity.User{}
	config.DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		ERR.Panic(Recode.NOT_FOUND, "user")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		ERR.Panic(Recode.WRONG_PASSWORD, "")
	}
	return utils.GenerateJWT(username)
}

func CreateUser(user *entity.User) (entity.User, error) {
	user.Password = utils.EncryptPassword(user.Password)
	err := config.DB.Create(&user).Error
	if err != nil {
		return entity.User{}, err
	} else {
		user.Password = ""
	}
	return *user, nil
}

func FindUsers() ([]entity.User, error) {
	var users []entity.User
	err := config.DB.Omit("password").Find(&users).Error
	if err != nil {
		return []entity.User{}, err
	}
	return users, nil
}

func FindUser(id uint64) (entity.User, error) {
	var user entity.User
	err := config.DB.Where("id = ?", id).Omit("password").First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func DeleteUser(id uint64) error {
	var user entity.User
	err := config.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdatePassword(id uint64, password string) error {
	var user entity.User
	err := config.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	user.Password = utils.EncryptPassword(password)
	err = config.DB.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
