package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"lema/src/models"
	"log"
	"time"
)

var SqliteDB *gorm.DB

func ConnectDB() {

	DBName := viper.GetString("DB_NAME")

	var dbError error
	SqliteDB, dbError = gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	if dbError != nil {
		log.Fatalf("Failed to connect to SQLite database: %v", dbError)
	}

	sqlDB, poolError := SqliteDB.DB()
	if poolError != nil {
		log.Fatalf("Error creating database connection pool: %v", poolError)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Connected to SQLite Database")

	err1 := SqliteDB.AutoMigrate(
		&models.Address{},
		&models.User{},
		&models.Post{},
	)

	if err1 != nil {
		log.Fatalf("Error migrating SQLite DB: %v", err1)
	} else {
		fmt.Println("SQLite DB successfully migrated")
	}

	//seed users
	SeedUsers()

}
