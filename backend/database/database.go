package database

import (
	"task-management-backend/config"
        "github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Task struct {
    gorm.Model
    Title       string
    Description string
    Status      string
}

func Init() (*gorm.DB, error) {
	dbConfig := config.LoadConfig()

	//connection string

	connectionString := "host=" + dbConfig.Host +
	"port=" + dbConfig.Port + 
	"user=" + dbConfig.Username + 
	"dbname=" + dbConfig.Name + 
	"password=" + dbConfig.Password + 
	"sslmode=disable"

	//connecting the db
	db, err := gorm.Open(dbConfig.Dialect, connectionString)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	err = db.AutoMigrate(&Task{}).Error
	if err != nil {
		return nil, err
	}

	DB = db

	return db, nil
}
