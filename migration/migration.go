package migration

import (
	"github.com/chaossat/tiktak/model"
	"github.com/jinzhu/gorm"
)

//TODO：更新其他需要储存在mysql的结构体
func SetAutoMigrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User_info{})
	//.........
}
