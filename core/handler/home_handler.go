package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(c echo.Context) {
	c.JSON(http.StatusOK, "Welcome To This Awesome API")
}
