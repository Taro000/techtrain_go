package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"techtrain_go_api/models"
)


var (
	db  *gorm.DB
	err error
)

const (
	// DBMS
	Dialect = "mysql"

	// User Name
	DBUser = "root"

	// Password
	DBPass = "gogocyber"

	// Protocol
	DBProtocol = "tcp(localhost:3307)"

	// DB Name
	DBName = "techtrain_game"
)


// Initialize DB from main()
func Init() {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err = gorm.Open(Dialect, connect+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	// db.SingularTable(true)
}


// Called in models
func GetDB() *gorm.DB {
	return db
}


// Close DB
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration()  {
	db.AutoMigrate(&models.User{}, &models.UserCharacter{}, &models.Character{})
}