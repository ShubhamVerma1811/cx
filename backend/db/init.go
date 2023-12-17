package db

import (
	"cx/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB(db *gorm.DB) (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(sqlite.Open("cx.db"), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	db.AutoMigrate(&model.Link{})
	db.AutoMigrate(&model.User{})

	return db, nil
}
