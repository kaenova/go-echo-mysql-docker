package db

import (
	"database/sql"
	"fmt"
	"go_echo_rest/config"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	fmt.Println("==== Initializing DB ====")
	db = InitSchema()
	CheckTable(db)
	conf := config.GetConfig()
	connectionString := conf.DB_Username + ":" + conf.DB_Pass + "@tcp(" + conf.DB_HOST + ":" + strconv.Itoa(conf.DB_PORT) + ")/" + conf.DB_Name
	fmt.Println("Trying to connect with", connectionString)
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Database")
}

func CreateCon() *sql.DB {
	return db
}
