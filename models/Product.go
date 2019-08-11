package models

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
//	ID        uint
 //   CreatedAt time.Time
	Code  string
	Price uint
}
