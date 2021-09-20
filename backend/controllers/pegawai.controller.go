package controllers

import (
	"go_echo_rest/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Get
func FetchPegawai(c echo.Context) error {
	var (
		id     string          = ""
		nama   string          = ""
		result models.Response = models.CreateResponse()
		err    error           = nil
	)

	id = c.QueryParam("id")
	nama = c.QueryParam("nama")

	// if no query
	if id != "" || strings.TrimSpace(nama) != "" {
		result, err = models.FetchPegawaiIDOrNama(id, nama)
	} else {
		result, err = models.FetchAllPegawai()
	}
	if err != nil {
		result.Message = err.Error()
		return c.JSON(result.Status, result)
	}

	return c.JSON(http.StatusOK, result)
}

// Post
func InputPegawai(c echo.Context) error {
	var (
		nama    string          = c.FormValue("nama")
		alamat  string          = c.FormValue("alamat")
		telepon string          = c.FormValue("telepon")
		result  models.Response = models.CreateResponse()
	)

	if len(nama) > 250 {
		result.Status = http.StatusBadRequest
		result.Message = "Nama terlalu panjang"
		return c.JSON(result.Status, result)
	}
	if len(telepon) > 15 {
		result.Status = http.StatusBadRequest
		result.Message = "Nomor Telepon terlalu panjang"
		return c.JSON(result.Status, result)
	}

	result, err := models.StorePegawai(nama, alamat, telepon)
	if err != nil {
		result.Status = http.StatusInternalServerError
		result.Message = err.Error()
		return c.JSON(result.Status, result)
	}
	return c.JSON(result.Status, result)
}

// Delete
func DeletePegawai(c echo.Context) error {
	var (
		nama    string          = c.FormValue("nama")
		telepon string          = c.FormValue("telepon")
		result  models.Response = models.CreateResponse()
	)

	if len(nama) > 250 {
		result.Message = "Nama terlalu panjang"
		result.Status = http.StatusBadRequest
		return c.JSON(result.Status, result)
	}
	if len(telepon) > 15 {
		result.Message = "Nomor Telepon panjang"
		result.Status = http.StatusBadRequest
		return c.JSON(result.Status, result)
	}

	result, err := models.DeletePegawai(nama, telepon)
	if err != nil {
		result.Status = http.StatusInternalServerError
		result.Message = err.Error()
		return c.JSON(result.Status, result)
	}
	return c.JSON(http.StatusOK, result)
}

// Put
func UpdatePegawaiByName(c echo.Context) error {
	// Ini hanya by Name, jadi kemungkinan masih bisa banyak bug

	var (
		nama    string          = c.FormValue("nama")
		alamat  string          = c.FormValue("alamat")
		telepon string          = c.FormValue("telepon")
		result  models.Response = models.CreateResponse()
	)

	if len(strings.TrimSpace(alamat))+len(strings.TrimSpace(telepon)) == 0 {
		result.Message = "Tidak boleh kosong"
		result.Status = http.StatusBadRequest
		return c.JSON(result.Status, result)
	}

	if len(strings.TrimSpace(nama)) > 250 {
		result.Message = "Nama terlalu panjang atau todal valid"
		result.Status = http.StatusBadRequest
		return c.JSON(result.Status, result)
	}

	if len(strings.TrimSpace(telepon)) > 15 {
		result.Message = "Nomor Telepon panjang atau tidak valid"
		result.Status = http.StatusBadRequest
		return c.JSON(result.Status, result)
	}

	result, err := models.UpdatePegawai(nama, alamat, telepon)
	if err != nil {
		result.Status = http.StatusInternalServerError
		result.Message = err.Error()
		return c.JSON(result.Status, result)
	}
	return c.JSON(result.Status, result)
}
