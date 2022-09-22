package modeltests

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"github.com/wewebplus/wewebapi/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllUsers(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatalf("Error refreshing user table %v\n", err)
	}

	err = seedUsers()
	if err != nil {
		log.Fatalf("Error seeding user table %v\n", err)
	}

	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatalf("Error user refreshing table %v\n", err)
	}
	newUser := models.SysStf{
		SyStfId:       1,
		SyStfEmail:    "test@gmail.com",
		SyStfFnameeng: "test",
		SyStfPassword: "password",
	}
	savedUser, err := newUser.SaveUser(server.DB)
	if err != nil {
		t.Errorf("Error while saving a user: %v\n", err)
		return
	}
	assert.Equal(t, newUser.SyStfId, savedUser.SyStfId)
	assert.Equal(t, newUser.SyStfEmail, savedUser.SyStfEmail)
	assert.Equal(t, newUser.SyStfFnameeng, savedUser.SyStfFnameeng)
}

func TestGetUserByID(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatalf("Error user refreshing table %v\n", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	foundUser, err := userInstance.FindUserByID(server.DB, user.SyStfId)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.SyStfId, user.SyStfId)
	assert.Equal(t, foundUser.SyStfEmail, user.SyStfEmail)
	assert.Equal(t, foundUser.SyStfFnameeng, user.SyStfFnameeng)
}

func TestUpdateAUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	userUpdate := models.SysStf{
		SyStfId:       1,
		SyStfFnameeng: "modiUpdate",
		SyStfEmail:    "modiupdate@gmail.com",
		SyStfPassword: "password",
	}
	updatedUser, err := userUpdate.UpdateAUser(server.DB, user.SyStfId)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedUser.SyStfId, userUpdate.SyStfId)
	assert.Equal(t, updatedUser.SyStfEmail, userUpdate.SyStfEmail)
	assert.Equal(t, updatedUser.SyStfFnameeng, userUpdate.SyStfFnameeng)
}

func TestDeleteAUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()

	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	isDeleted, err := userInstance.DeleteAUser(server.DB, user.SyStfId)
	if err != nil {
		t.Errorf("this is the error deleting the user: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}
