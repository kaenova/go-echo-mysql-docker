package db

import (
	"database/sql"
	"fmt"
	"go_echo_rest/config"
)

func InitTable(db *sql.DB) *sql.DB {
	fmt.Println("Creating Table")
	// Creating table

	// Create Pegawai table
	_, err := db.Exec(fmt.Sprintf("CREATE TABLE `%s`.`Pegawai` (`idPegawai` INT NOT NULL AUTO_INCREMENT,`Nama` VARCHAR(250) NOT NULL, `Alamat` MEDIUMTEXT NULL, `Telepon` VARCHAR(15) NOT NULL, PRIMARY KEY (`idPegawai`), UNIQUE INDEX `idPegawai_UNIQUE` (`idPegawai` ASC) VISIBLE);", config.GetConfig().DB_Name))
	if err != nil {
		fmt.Println(err)
		panic("Cannot create table Pegawai")
	}

	fmt.Println("Table Created")
	return db
}

func CheckTable(db *sql.DB) {
	// Checking Table, if Table not available or wrong, go InitTable()
	var valid bool = false

	// Making map of Table Needed in Schema
	var arr_table []string = config.GetConfig().DB_Table
	var map_table = make(map[string]bool)
	for i := 0; i < len(arr_table); i++ {
		map_table[arr_table[i]] = true
	}

	for !valid {
		var counter int = 0
		// Get Tables
		rows, err := db.Query("SHOW TABLES")
		if err != nil {
			fmt.Println(err)
			panic("Cannot query to check table")
		}

		// Scanning
		var temp_table string
		for rows.Next() {
			counter++
			rows.Scan(&temp_table)
			if map_table[temp_table] {
				valid = true
			} else {
				valid = false
			}
		}

		if counter < 1 {
			valid = false
		}

		if !valid {
			db = InitTable(db)
		}
	}

	fmt.Println("Table is OK!")
}
