package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wewebplus/wewebapi/api/auth"
	"github.com/wewebplus/wewebapi/api/models"
	"github.com/wewebplus/wewebapi/api/responses"
	"github.com/wewebplus/wewebapi/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.SysStf{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
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

	user := models.SysStf{}

	err = server.DB.Debug().Model(models.SysStf{}).Where("email = ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.SyStfPassword, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.SyStfId)
}
