package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/wewebplus/wewebapi/models"
	"github.com/wewebplus/wewebapi/responses"
	"github.com/wewebplus/wewebapi/types"
)

func (server *Server) CreateUser(c echo.Context) error {
	user := types.SysStf{}
	if err := c.Bind(&user); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, responses.ParseStatus("REQ_ERR", ""))
	}
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ParseStatus("REQ_INVALID", err.Error()))
	}
	models.Prepare(&user)
	err := models.Validate(&user, "")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ParseStatus("REQ_INVALID", err.Error()))
	}

	data, err := models.SaveUser(server.DB, &user)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, responses.ParseStatus("NOT_ACCEPTED", err.Error()))
	}
	return c.JSON(http.StatusCreated, data)
}

func (server *Server) GetUsers(c echo.Context) error {

	user := types.SysStf{}
	if err := c.Bind(&user); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, responses.ParseStatus("REQ_ERR", ""))
	}
	data, err := models.FindAllUsers(server.DB, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ParseStatus("NOT_ACCEPTED", err.Error()))
	}
	return c.JSON(http.StatusOK, data)
}

func (server *Server) GetUser(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ParseStatus("REQ_INVALID", "ID invalid"))
	}
	user := types.SysStf{}
	data, err := models.FindUserByID(&user, server.DB, uint32(uid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ParseStatus("NOT_ACCEPTED", err.Error()))
	}
	return c.JSON(http.StatusOK, data)
}

func (server *Server) UpdateUser(c echo.Context) error {

	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ParseStatus("REQ_INVALID", "ID invalid"))
	}

	user := types.SysStf{}
	if err := c.Bind(&user); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, responses.ParseStatus("REQ_ERR", "Có lỗi xảy ra, vui lòng kiểm tra lại thông tin"))
	}
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ParseStatus("REQ_INVALID", err.Error()))
	}
	models.Prepare(&user)
	err = models.Validate(&user, "update")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, responses.ParseStatus("REQ_INVALID", err.Error()))
	}
	data, err := models.UpdateAUser(server.DB, &user, uint32(uid))
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, responses.ParseStatus("NOT_ACCEPTED", err.Error()))
	}
	return c.JSON(http.StatusOK, data)
}

func (server *Server) DeleteUser(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("id"))
	user := types.SysStf{}
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ParseStatus("REQ_INVALID", "ID invalid"))
	}
	data, err := models.DeleteAUser(server.DB, &user, uint32(uid))
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, responses.ParseStatus("NOT_ACCEPTED", err.Error()))
	}
	return c.JSON(http.StatusOK, data)
}
