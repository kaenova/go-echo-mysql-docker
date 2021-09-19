package models

import (
	"errors"
	"fmt"
	"go_echo_rest/db"
	"net/http"
	"strconv"
	"strings"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

// GET
func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var arrObj []Pegawai
	var res Response
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM Pegawai"
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}
		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj
	return res, nil
}

func FetchPegawaiID(id_string string) (Response, error) {
	var (
		res    Response
		obj    Pegawai
		arrObj []Pegawai
	)
	_, err := strconv.Atoi(id_string)
	if err != nil {
		return res, errors.New("ID Tidak Valid")
	}

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM Pegawai WHERE idPegawai = " + id_string

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	if arrObj == nil {
		res.Status = http.StatusBadRequest
		res.Message = "Data tidak tersedia"
		res.Data = nil
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj
	return res, nil
}

// POST
func StorePegawai(nama, alamat, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement :=
		fmt.Sprintf(`INSERT INTO Pegawai (Nama, Alamat, Telepon) VALUES ("%s", "%s", "%s")`,
			nama, alamat, telepon)

	result, err := con.Exec(sqlStatement)
	if err != nil {
		return res, err
	}
	if check, err_temp := result.RowsAffected(); check == 0 || err_temp != nil {
		return res, errors.New("Gagal memasukkan Pegawai")
	}

	var obj = struct {
		Nama    string `json:"nama"`
		Alamat  string `json:"alamat"`
		Telepon string `json:"telepon"`
	}{}

	obj.Nama = nama
	obj.Alamat = alamat
	obj.Telepon = telepon

	res.Status = http.StatusOK
	res.Message = "Data Inputted"
	res.Data = obj
	return res, nil
}

// Delete
func DeletePegawai(nama, telepon string) (Response, error) {
	var res Response

	sqlStatement := `DELETE FROM Pegawai WHERE Nama="` + nama + `" AND Telepon="` +
		telepon + `"`
	con := db.CreateCon()
	result, err := con.Exec(sqlStatement)
	if err != nil {
		res.Status = http.StatusInternalServerError
		return res, err
	}
	if check, err_temp := result.RowsAffected(); check == 0 || err_temp != nil {
		return res, errors.New("Nama atau Telepon tidak valid.")
	}

	var obj = struct {
		Nama    string `json:"nama"`
		Telepon string `json:"telepon"`
	}{}

	obj.Nama = nama
	obj.Telepon = telepon

	res.Data = obj
	res.Message = "Success"
	res.Status = http.StatusOK

	return res, nil
}

// PUT
func UpdatePegawai(nama, alamat, telepon string) (Response, error) {
	var (
		res          Response
		check_alamat bool = false
	)

	var obj = struct {
		Nama    string `json:"nama"`
		Alamat  string `json:"alamat"`
		Telepon string `json:"telepon"`
	}{}
	obj.Nama = nama

	sqlStatement := "UPDATE Pegawai SET "

	if len(strings.TrimSpace(alamat)) != 0 {
		sqlStatement = sqlStatement + fmt.Sprintf(`Alamat = "%s" `, alamat)
		check_alamat = true
		obj.Alamat = alamat
	}

	if len(strings.TrimSpace(telepon)) != 0 {
		if check_alamat {
			sqlStatement = sqlStatement + ", "
		}
		sqlStatement = sqlStatement + fmt.Sprintf(`Telepon = "%s" `, telepon)
		obj.Telepon = telepon
	}

	sqlStatement = sqlStatement + fmt.Sprintf(`WHERE Nama = "%s" `, nama)

	con := db.CreateCon()
	result, err := con.Exec(sqlStatement)
	if err != nil {
		res.Status = http.StatusInternalServerError
		return res, err
	}
	if check, err_temp := result.RowsAffected(); check == 0 || err_temp != nil {
		res.Status = http.StatusBadRequest
		return res, errors.New("Nama atau Telepon tidak tersedia di penyimpanan data.")
	}

	res.Data = obj
	res.Message = "Success"
	res.Status = http.StatusOK

	return res, nil

}
