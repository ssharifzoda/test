package database

import (
	"fmt"
	"github.com/spf13/viper"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func NewConnectPostgres() *gorm.DB {
	Host := viper.GetString("db.host")
	Port := viper.GetUint16("db.port")
	Username := viper.GetString("db.username")
	Password := os.Getenv("DB_PASSWORD")
	DBName := viper.GetString("db.dbname")
	dbParams := fmt.Sprintf("host=%s password=%s user=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Dushanbe",
		Host, Password, Username, DBName, Port)
	conn, err := gorm.Open(postgresDriver.Open(dbParams), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("postgres conn success")
	return conn
}
