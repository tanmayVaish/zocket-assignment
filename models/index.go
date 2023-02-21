package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB_Host     string
	DB_Port     string
	DB_User     string
	DB_Password string
	DB_Name     string
	DB_SSLMode  string
}

var DB *gorm.DB

func InitDB(cfg Config) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DB_Host, cfg.DB_User, cfg.DB_Password, cfg.DB_Name, cfg.DB_Port, cfg.DB_SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&Book{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db
}
