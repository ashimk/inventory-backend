package main

import (
	"flag"
	"fmt"
	"inventory/common"
	"inventory/service"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	flag.Parse()
	config, err := common.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DbUser, config.DbPass, config.DbUrl, config.DbName)
	DB, err := gorm.Open("mysql", url)
	DB = DB.LogMode(true)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	DB.AutoMigrate(service.Inventory{})
}
