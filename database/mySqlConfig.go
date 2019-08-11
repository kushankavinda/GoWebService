package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/webAPi/models"
)

func TestingDb() {
	fmt.Print("start db operations")
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/otpUsers?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	productDataModal := models.Product{}
	db.AutoMigrate(productDataModal)

	// Create

	db.Create(models.Product{Code: "L1212", Price: 1000})

	// Read
	var product models.Product
	//db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
	fmt.Print("db operarions done")
}
func init() {
	fmt.Print("labbe banis")
}
