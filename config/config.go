package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


type Configuration struct{
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT string
	DB_HOST string
	DB_NAME string
}

var db *gorm.DB
var err error

func GetConfig() Configuration{
	configuration := Configuration{}
	gonfig.GetConf("config/config.json", &configuration)
	return configuration
}
//uname:pswd@tcp(<ip>:<port>)/<dbname>?charset=utf8mb4&parseTime=True&loc=Local
func Init(){
	conf := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func GetDB() *gorm.DB{
	return db
}



