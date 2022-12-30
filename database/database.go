package database

import (
	"fmt"
	"poke-api-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DbConfig struct {
	host     string
	port     uint
	dbname   string
	user     string
	password string
}

var config = DbConfig{
	host:     "localhost",
	port:     5432,
	dbname:   "poke-api",
	user:     "poke-api",
	password: "poke-api",
}

var Dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow", config.host, config.user, config.password, config.dbname, config.port)


type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// connectDb
func ConnectDb() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	if err = db.AutoMigrate(&models.Pokemon{}); err !=nil {
		log.Fatal(err.Error())
	}

	DB = Dbinstance{
		Db: db,
	}
}