package database

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USERNAME")
	pass := viper.GetString("DB_PASSWORD")
	dbname := viper.GetString("DB_DATABASE")
	port := viper.GetString("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", host, user, pass, dbname, port)
	conf := &gorm.Config{
		SkipDefaultTransaction: true,
	}
	if os.Getenv("ENV") != "dev" {
		conf.Logger = logger.Default.LogMode(logger.Silent)
	}
	db, err := gorm.Open(postgres.Open(dsn), conf)
	if err != nil {
		panic(err)
	}
	DB = db
}

func Close() {
	c, _ := DB.DB()
	c.Close()
}
