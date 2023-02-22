package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// StaticDatabaseInstance : database pointer
type StaticDatabaseInstance struct {
	Db *gorm.DB
}

// Database : object instance
var Database StaticDatabaseInstance

const DNS string = "bbce14802e2bf7:765f916b@tcp(eu-cdbr-west-03.cleardb.net:3306)/heroku_9a62a0aad140d17?charset=utf8mb4&parseTime=True&loc=Local"

// ConnectToDatabase : Connect to the database
func ConnectToDatabase() {
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Successfully connected to the database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	// Launch the database migration
	RunMigration(db)

	Database = StaticDatabaseInstance{Db: db}
}
