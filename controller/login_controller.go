package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/wewebplus/wewebapi/auth"
	"github.com/wewebplus/wewebapi/models"
	"github.com/wewebplus/wewebapi/responses"
	"github.com/wewebplus/wewebapi/types"
	"github.com/wewebplus/wewebapi/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(c echo.Context) error {
	user := types.SysStf{}
	if err := c.Bind(&user); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, responses.ParseStatus("REQ_ERR", ""))
	}

	models.Prepare(&user)
	err := models.Validate(&user, "login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.SyStfUsername, user.SyStfPassword)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(username, password string) (string, error) {

	var err error

	user := types.SysStf{}

	err = server.DB.Debug().Model(types.SysStf{}).Where("email = ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.SyStfPassword, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.SyStfId)
}
