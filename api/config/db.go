/**
 * Author: Goddy <goddy@mykg.ai> 2020-12-10
 */
package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mykg.ai/rua/domain/entity"
	r "mykg.ai/rua/domain/entity/relation"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Database),
		DefaultStringSize: 255,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex: true,
		DontSupportRenameColumn: true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db
}

func MigrateDomains()  {
	/** relationship */
	DB.AutoMigrate(&r.LinkTag{})
	DB.AutoMigrate(&r.UserNamespace{})

	/** domain */
	DB.AutoMigrate(&entity.Link{})
	DB.AutoMigrate(&entity.Namespace{})
	DB.AutoMigrate(&entity.Tag{})
	DB.AutoMigrate(&entity.User{})

	/** log */
	DB.AutoMigrate(&entity.OpLog{})
	DB.AutoMigrate(&entity.UsageLog{})
}
