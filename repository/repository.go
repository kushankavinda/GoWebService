package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/webAPi/models"
)

func Create(any interface{}) error {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/otpUsers?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	productDataModal := models.User{}
	db.AutoMigrate(productDataModal)

	// Create

	db.Create(any)
	fmt.Println("repository")
	fmt.Print(any)

	return nil
}
