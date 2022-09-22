package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"github.com/wewebplus/wewebapi/core/types"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
func MD5PasswordHash(password string) (string, error) {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hex.EncodeToString(hasher.Sum(nil))
	md5Value := hex.EncodeToString(hasher.Sum(nil))
	return md5Value, nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func BeforeSave(user *types.SysStf) error {
	hashedPassword, err := MD5PasswordHash(user.SyStfPassword)
	if err != nil {
		return err
	}
	user.SyStfPassword = string(hashedPassword)
	return nil
}

func Prepare(user *types.SysStf) {
	user.SyStfId = 0
	user.SyStfPrefix = html.EscapeString(strings.TrimSpace(user.SyStfPrefix))
	user.SyStfGender = html.EscapeString(strings.TrimSpace(user.SyStfGender))
	user.SyStfFnameeng = html.EscapeString(strings.TrimSpace(user.SyStfFnameeng))
	user.SyStfFnamethai = html.EscapeString(strings.TrimSpace(user.SyStfFnamethai))
	user.SyStfLnameeng = html.EscapeString(strings.TrimSpace(user.SyStfLnameeng))
	user.SyStfLnamethai = html.EscapeString(strings.TrimSpace(user.SyStfLnamethai))
	user.SyStfUsername = html.EscapeString(strings.TrimSpace(user.SyStfUsername))
	user.SyStfPassword = html.EscapeString(strings.TrimSpace(user.SyStfPassword))
	user.SyStfAddress = html.EscapeString(strings.TrimSpace(user.SyStfAddress))
	user.SyStfTelephone = html.EscapeString(strings.TrimSpace(user.SyStfTelephone))
	user.SyStfMobile = html.EscapeString(strings.TrimSpace(user.SyStfMobile))
	user.SyStfEmail = html.EscapeString(strings.TrimSpace(user.SyStfEmail))
	user.SyStfOther = html.EscapeString(strings.TrimSpace(user.SyStfOther))
	user.SyStfPicture = html.EscapeString(strings.TrimSpace(user.SyStfPicture))
	user.SyStfCreby = html.EscapeString(strings.TrimSpace(user.SyStfCreby))
	user.SyStfCredate = time.Now()
	user.SyStfLastdate = time.Now()
	user.SyStfStatus = html.EscapeString(strings.TrimSpace(user.SyStfStatus))
	user.SyStfLogdate = time.Now()
	user.SyStfMasterkey = html.EscapeString(strings.TrimSpace(user.SyStfMasterkey))
	user.SyStfPosition = html.EscapeString(strings.TrimSpace(user.SyStfPosition))
	user.SyStfPart = html.EscapeString(strings.TrimSpace(user.SyStfPart))
	user.SyStfPositionuser = html.EscapeString(strings.TrimSpace(user.SyStfPositionuser))
}

func Validate(user *types.SysStf, action string) error {
	switch strings.ToLower(action) {
	case "update":
		if user.SyStfFnameeng == "" {
			return errors.New("Required Nickname")
		}
		if user.SyStfPassword == "" {
			return errors.New("Required Password")
		}
		if user.SyStfEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.SyStfEmail); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if user.SyStfPassword == "" {
			return errors.New("Required Password")
		}
		if user.SyStfEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.SyStfEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if user.SyStfFnameeng == "" {
			return errors.New("Required Nickname")
		}
		if user.SyStfPassword == "" {
			return errors.New("Required Password")
		}
		if user.SyStfEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(user.SyStfEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func SaveUser(db *gorm.DB, user *types.SysStf) (*types.SysStf, error) {
	err := db.Table("sy_stf").Create(&user).Error
	if err != nil {
		return &types.SysStf{}, err
	}
	return user, nil
}

func FindAllUsers(db *gorm.DB, user *types.SysStf) (*[]types.SysStf, error) {
	var err error
	users := []types.SysStf{}
	err = db.Debug().Model(&types.SysStf{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]types.SysStf{}, err
	}
	return &users, err
}

func FindUserByID(user *types.SysStf, db *gorm.DB, uid uint32) (*types.SysStf, error) {
	err := db.Debug().Model(types.SysStf{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &types.SysStf{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &types.SysStf{}, errors.New("User Not Found")
	}
	return user, err
}

func UpdateAUser(db *gorm.DB, user *types.SysStf, uid uint32) (*types.SysStf, error) {

	// To hash the password
	err := BeforeSave(user)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&types.SysStf{}).Where("id = ?", uid).Take(&types.SysStf{}).UpdateColumns(
		map[string]interface{}{
			"password":  user.SyStfPassword,
			"nickname":  user.SyStfFnameeng,
			"email":     user.SyStfEmail,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &types.SysStf{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&types.SysStf{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &types.SysStf{}, err
	}
	return user, nil
}

func DeleteAUser(db *gorm.DB, user *types.SysStf, uid uint32) (int64, error) {

	db = db.Debug().Model(&types.SysStf{}).Where("id = ?", uid).Take(&types.SysStf{}).Delete(&types.SysStf{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
