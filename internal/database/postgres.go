package database

import (
	"fmt"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func NewConnectPostgres() *gorm.DB {
	dbParams := fmt.Sprint("host=localhost password=sin user=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Dushanbe")
	conn, err := gorm.Open(postgresDriver.Open(dbParams), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("postgres conn success")
	return conn
}
