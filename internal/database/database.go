package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type single struct {
	Db *gorm.DB
}

var (
	lock           = &sync.Mutex{}
	singleInstance *single
)

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		log.Println("Creating single instance now")
		singleInstance = &single{
			Db: Connection(),
		}
	} else {
		log.Println("Single instance is already created")
	}
	return singleInstance
}

func Connection() *gorm.DB {
	e := godotenv.Load()
	if e != nil {
		log.Println("Fail loading .env file\n", e)
	}
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s",
		os.Getenv("db_host"),
		os.Getenv("db_user"),
		os.Getenv("db_name"),
		os.Getenv("db_port"),
		os.Getenv("db_password"))
	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Println("Failed database connection\n", err)
	}
	return conn
}
