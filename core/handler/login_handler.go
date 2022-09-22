package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wewebplus/wewebapi/auth"
	"github.com/wewebplus/wewebapi/core/models"
	"github.com/wewebplus/wewebapi/core/types"
	"github.com/wewebplus/wewebapi/responses"
	"github.com/wewebplus/wewebapi/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := types.SysStf{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	models.Prepare(&user)
	err = models.Validate(&user, "login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(user.SyStfUsername, user.SyStfPassword)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func SignIn(username, password string) (string, error) {

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
