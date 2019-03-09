package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetDataBase() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}

func Migrate() {
	db, _ := GetDataBase()
	db.AutoMigrate(&City{})
	defer db.Close()
}
