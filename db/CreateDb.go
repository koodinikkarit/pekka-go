package pekka

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func CreateDb(
	dbUser string,
	dbPass string,
	dbIp string,
	dbPort string,
	dbName string,
) *gorm.DB {
	db, err := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbIp+":"+dbPort+")/"+dbName)
	if err != nil {
		fmt.Println("Creating database connection failed")
	}

	db.AutoMigrate(&Button{})
	db.AutoMigrate(&EventTimer{})
	db.AutoMigrate(&Executor{})
	db.AutoMigrate(&ExecutorAction{})
	db.AutoMigrate(&Pentti{})
	db.AutoMigrate(&WeeklyTimer{})

	return db
}
