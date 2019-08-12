package database

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var once sync.Once

// type global
type singleton map[string]string

var (
	instance singleton
)

func DatabaseConnection() singleton {

	once.Do(func() { // <-- atomic, does not allow repeating
		gorm.Open("mysql", "root:@tcp(localhost:3306)/otpUsers?charset=utf8&parseTime=True&loc=Local")
		instance = make(singleton) // <-- thread safe

	})

	return instance
}
