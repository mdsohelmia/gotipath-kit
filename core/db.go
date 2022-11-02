package core

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(env Env, logger Logger) Database {
	username := env.DB_USERNAME
	password := env.DB_PASSWORD
	host := env.DB_HOST
	port := env.DB_PORT
	dbname := env.DB_DATABASE
	fmt.Println(password)

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	fmt.Println(url)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	logger.Info("Database connection established")

	return Database{
		DB: db,
	}
}
