package db

import (
	"demo-clean/entitiies"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, func()) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(entitiies.UserModel{})

	sqlDB, _ := db.DB()

	tearDown := func() {
		defer sqlDB.Close()
	}

	return db, tearDown

}
