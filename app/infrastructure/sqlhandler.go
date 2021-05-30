package infrastructure

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"go_http_api_template/interfaces/database"
)

type SqlHandler struct {
	db  *gorm.DB
	err error
}

func NewSqlHandler() database.SqlHandler {
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
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		panic(err.Error)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}
