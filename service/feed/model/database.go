package model

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"

	"github.com/chaossat/tiktak/migration"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var Db *gorm.DB
var Rdb *redis.Client

// 加载redis连接
func InitRedis() error {
	address := viper.GetString("datasource.redis.host")
	port := viper.GetString("datasource.redis.port")
	password := viper.GetString("datasource.redis.password")
	db_id := viper.GetInt("datasource.redis.DB")
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", address, port),
		Password: password,
		DB:       db_id,
	})
	_, err := Rdb.Ping().Result()
	if err != nil {
		log.Println("连接redis失败", err.Error())
		return err
	}
	return nil
}

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
	log.Println("连接mysql数据库:", args)
	if err != nil {
		log.Println("连接mysql失败", err)
		panic(err)
	}
	migration.SetAutoMigrate(tpdb)
	Db = tpdb
	return Db
}

//GetDB:返回db
func GetDB() *gorm.DB {
	return Db
}

//GetRDB:返回redisDB
func GetRDB() *redis.Client {
	return Rdb
}
