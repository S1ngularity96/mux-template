package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

)


var DbCon *gorm.DB

func MigrateData(){
	DbCon.AutoMigrate(&Note{})
}

func OpenDatabase() {
	db, err := gorm.Open("sqlite3", "test.db")
	DbCon = db
	if err != nil {
		panic("failed to connect database")
	}
}
func CloseDatabase() {
	if DbCon != nil{
		DbCon.Close()
	}
}
