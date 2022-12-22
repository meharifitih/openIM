package rpc

import (
	api "Open_IM/pkg/base_info"
	"Open_IM/pkg/common/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Config.Mysql.DBUserName, config.Config.Mysql.DBPassword, config.Config.Mysql.DBAddress[0], "openIM_v2")

	var db *gorm.DB
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		fmt.Println("Open failed ", err.Error(), dsn)
	}

	db.AutoMigrate(&api.Banner{}, &api.MiniApp{}, &api.UserStory{})

	return db
}
