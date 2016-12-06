package database

import (
	"fmt"

	"github.com/LastZactionHero/contact_us/models"
	"github.com/jinzhu/gorm"
)

// DB database instance
var DB *gorm.DB

// DBConnect connect to database
func DBConnect() *gorm.DB {
	fmt.Println("Connecting to db")
	dbc, err := gorm.Open("postgres", "host=db user=postgres dbname=postgres sslmode=disable password=")
	fmt.Println("Connected!")

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
