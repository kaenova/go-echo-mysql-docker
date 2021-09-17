package db

import (
	"database/sql"
	"fmt"
	"go_echo_rest/config"
	"strconv"
)

var (
	first_time bool = true
	create_db  bool = true
)

func InitSchema() *sql.DB {
	/*
		Check is DB available?
		if not, create Schema and the Table ✅
		if available, check schema => if table is no good, drop all table and create table
	*/

	var db *sql.DB
	conf := config.GetConfig()

	// Check DB Is Available
	for create_db {
		connectionString := conf.DB_Username + ":" + conf.DB_Pass + "@tcp(" + conf.DB_HOST + ":" + strconv.Itoa(conf.DB_PORT) + ")/" + conf.DB_Name
		fmt.Println("Trying to connect with", connectionString)
		db, err = sql.Open("mysql", connectionString)
		err = db.Ping()
		if err != nil {
			db = createSchema(db)
		} else {
			create_db = false
		}
	}
	fmt.Println("DB Connected")
	return db
}

func createSchema(db *sql.DB) *sql.DB {
	conf := config.GetConfig()
	if first_time {
		// Initializing Database
		fmt.Println("Won't connect to DB, Probably Schema not Initialize")
		fmt.Println("Creating Schema")
		// Connecting to DB
		connectionString := conf.DB_Username + ":" + conf.DB_Pass + "@tcp(" + conf.DB_HOST + ":" + strconv.Itoa(conf.DB_PORT) + ")/"
		db, err = sql.Open("mysql", connectionString)
		if err != nil {
			panic("Wont connect to DB, check your configuration")
		}
		_, err = db.Exec(fmt.Sprintf("CREATE SCHEMA `%s` ;", conf.DB_Name))
		if err != nil {
			panic("Wont create to Schema, check your Databases")
		}
		fmt.Println("Schema Created")
		first_time = false
	} else {
		// Cannot Initialize Database
		panic("Can't create database")
	}
	err = nil

	return db
}
