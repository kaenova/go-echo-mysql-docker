package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	fmt.Println("==== Initializing DB ====")
	db = InitSchema()
	CheckTable(db)
	fmt.Println("Connected to Database")
}

func CreateCon() *sql.DB {
	return db
}
