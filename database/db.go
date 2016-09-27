package database

import (
	"fmt"
	"os"

	"github.com/LastZactionHero/contact_us/models"
	"github.com/jinzhu/gorm"
)

// DB database instance
var DB *gorm.DB

// DBConnect connect to database
func DBConnect() *gorm.DB {
	dbUser := os.Getenv("CONTACT_US_DB_USER")
	dbPass := os.Getenv("CONTACT_US_DB_PASS")
	dbName := os.Getenv("CONTACT_US_DB_NAME")
	connectStr := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)
	dbc, err := gorm.Open("mysql", connectStr)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	DB = dbc
	return DB
}

// DBInit initialize database
func DBInit() {
	DB.AutoMigrate(&models.Contact{})
	DB.AutoMigrate(&models.Skill{})
	DB.AutoMigrate(&models.Contractor{})
}
