package config

import (
	"fmt"
	"log"

	"github.com/zekhoi/golang-todo/app/entities"
	"github.com/zekhoi/golang-todo/constant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	// Retrieve database connection details from environment variables
	host := constant.MYSQL_HOST
	port := constant.MYSQL_PORT
	user := constant.MYSQL_USER
	password := constant.MYSQL_PASSWORD
	dbname := constant.MYSQL_DBNAME

	// Construct the database connection string

	// Postgre
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=true", user, password, host, port, dbname)
	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return err
	}

	err = db.AutoMigrate(&entities.Activity{}, &entities.ToDo{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Assign the database connection to the global variable
	DB = db

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
