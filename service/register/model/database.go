package model

import (
	"fmt"
	"log"
	"time"

	"github.com/chaossat/tiktak/migration"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var Db *gorm.DB

//TODO:支持redis
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := viper.GetString("datasource.mysql.host")
	port := viper.GetString("datasource.mysql.port")
	database := viper.GetString("datasource.mysql.database")
	username := viper.GetString("datasource.mysql.username")
	password := viper.GetString("datasource.mysql.password")
	charset := viper.GetString("datasource.mysql.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	tpdb, err := gorm.Open(driverName, args)
	if err != nil {
		log.Println(err.Error())
		log.Println(args)
		panic(err)
	}
	migration.SetAutoMigrate(tpdb)
	Db = tpdb
	Db.DB().SetMaxIdleConns(viper.GetInt("datasource.mysql.maxIdle"))
	Db.DB().SetMaxOpenConns(viper.GetInt("datasource.mysql.maxOpen"))
	Db.DB().SetConnMaxLifetime(time.Duration(viper.GetInt("datasource.mysql.maxLifetime")))
	return Db
}

//GetDB:返回db
func GetDB() *gorm.DB {
	return Db
}
