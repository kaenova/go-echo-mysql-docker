package routes

import (
	"fmt"
	"go_echo_rest/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	fmt.Println("==== Initializing Server ====")
	e := echo.New()

	// Home
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//Pegawai
	e.GET("/pegawai", controllers.FetchPegawai)
	e.POST("/pegawai", controllers.InputPegawai)
	e.DELETE("/pegawai", controllers.DeletePegawai)
	e.PUT("/pegawai", controllers.UpdatePegawaiByName)

	// Start
	e.Logger.Fatal(e.Start(":1323"))

	return e
}
