package infrastructure

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	connectTemplate := "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	connect := fmt.Sprintf(
		connectTemplate,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// connect to DB
	db, err = gorm.Open("mysql", connect)

	if err != nil {
		panic(err)
	}

	autoMigration(db)
}

// Db is called in models
func Db() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
