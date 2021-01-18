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

func connectDB() {
	var a = Config
	fmt.Println(a)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			Config.DB.User,
			Config.DB.Password,
			Config.DB.Host,
			Config.DB.Port,
			Config.DB.Database),
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

func migrateDomains()  {
	_ = DB.AutoMigrate(
		/** relationship */
		&r.FolderTag{},
		&r.LinkTag{},
		&r.UserNamespace{},

		/** domain */
		&entity.Folder{},
		&entity.FolderView{},
		&entity.Link{},
		&entity.Namespace{},
		&entity.OpLog{},
		&entity.Tag{},
		&entity.UsageLog{},
		&entity.User{},

		)
}
