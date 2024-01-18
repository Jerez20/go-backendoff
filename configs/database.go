package configs

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "sqlserver://prueba:prueba@127.0.0.1?database=PendaAsesys"
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println(err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Println(err)
		return
	}
}
