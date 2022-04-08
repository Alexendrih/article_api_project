package database

import (
	"article_api/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbIntance struct {
	Db *gorm.DB
}

var Database DbIntance

func ConnectionDb() {
	db, err := gorm.Open(postgres.Open("article.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Printf("Connect to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running")

	db.AutoMigrate(&models.User{}, &models.Article{}, &models.Userole{}, &models.Audit{})

	Database = DbIntance{Db: db}

}
