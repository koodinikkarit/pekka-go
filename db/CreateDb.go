package pekka

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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
		fmt.Println("Creating database connection failed", dbUser+":"+dbPass+"@tcp("+dbIp+":"+dbPort+")/"+dbName)
	}

	db.AutoMigrate(&Button{})
	db.AutoMigrate(&EventTimer{})
	db.AutoMigrate(&Executor{})
	db.AutoMigrate(&ExecutorAction{})
	db.AutoMigrate(&Pentti{})
	db.AutoMigrate(&WeeklyTimer{})

	return db
}
