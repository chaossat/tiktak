package migration

import (
	"github.com/chaossat/tiktak/model"
	"github.com/jinzhu/gorm"
)

//SetAutoMigrate: 根据结构体自动建表
func SetAutoMigrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Video{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Comment{})
}
