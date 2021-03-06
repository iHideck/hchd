package model

import (
	"log"

	"github.com/jasonsmithj/hchd/internal/adapter/configuration"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func mysqlConnect() *gorm.DB {
	db, err := gorm.Open("mysql", configuration.GetMySQLUri())
	if err != nil {
		log.Printf("db connect failed %s", err)
	}
	return db
}

func MySQLHealthcheck() error {
	c := mysqlConnect()
	err := c.DB().Ping()

	defer c.Close()

	return err
}
