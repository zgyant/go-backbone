package configs

import (
	"go-backbone/src/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {
	var err error
	dsn := utils.Bt(`host=${env("PG_HOST")} user=${env("PG_USER")} password=${env("PG_PASSWORD")} dbname=${env("PG_DB")} port=${env("PG_PORT")} sslmode=disable`)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("DB Connected!")

	return nil
}

// GetDB returns the GORM database instance
func GetDB() *gorm.DB {
	return db
}
