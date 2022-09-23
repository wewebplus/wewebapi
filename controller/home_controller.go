package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func (server *Server) Home(c echo.Context) {
	c.JSON(http.StatusOK, "Welcome To This Awesome API")
}
