package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/noclueps/fablab/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db struct {
	DB *gorm.DB
}

var Database Db

func getDSN() string {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv("CLEARDB_DATABASE_URL")
}

func ConnectDB() {
	dsn := getDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to db! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to db!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	db.AutoMigrate(&models.User{}, &models.Project{})

	Database = Db{DB: db}
}