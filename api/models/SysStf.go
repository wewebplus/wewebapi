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
	"golang.org/x/crypto/bcrypt"
)

type SysStf struct {
	SyStfId           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	SyStfPrefix       string    `gorm:"size:255;not null;" json:"prefix"`
	SyStfGender       string    `gorm:"size:100;not null;" json:"gender"`
	SyStfFnameeng     string    `gorm:"size:100;not null;" json:"fnameeng"`
	SyStfFnamethai    string    `gorm:"size:100;not null;" json:"fnamethai"`
	SyStfLnameeng     string    `gorm:"size:100;not null;" json:"lnameeng"`
	SyStfLnamethai    string    `gorm:"size:100;not null;" json:"lnamethai"`
	SyStfUsername     string    `gorm:"size:100;not null;" json:"username"`
	SyStfPassword     string    `gorm:"size:100;not null;" json:"password"`
	SyStfGroupid      int64     `gorm:"size:100;" json:"groupid,omitempty"`
	SyStfAddress      string    `gorm:"size:100;" json:"address,omitempty"`
	SyStfTelephone    string    `gorm:"size:100;" json:"telephone,omitempty"`
	SyStfMobile       string    `gorm:"size:100;" json:"mobile,omitempty"`
	SyStfEmail        string    `gorm:"size:100;not null;" json:"email"`
	SyStfOther        string    `gorm:"size:100;" json:"other,omitempty"`
	SyStfPicture      string    `gorm:"size:100;" json:"picture,omitempty"`
	SyStfCrebyid      int64     `gorm:"size:100;not null;" json:"crebyid"`
	SyStfCreby        string    `gorm:"size:100;not null;" json:"creby"`
	SyStfCredate      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"credate,omitempty"`
	SyStfLastdate     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"lastdate,omitempty"`
	SyStfStatus       string    `gorm:"size:100;not null;" json:"status"`
	SyStfOrder        int64     `gorm:"size:100;not null;" json:"order"`
	SyStfLogdate      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"logdate"`
	SyStfMasterkey    string    `gorm:"size:100;not null;" json:"masterkey"`
	SyStfUsercar      int64     `gorm:"size:100;not null;" json:"usercar"`
	SyStfUnitid       int64     `gorm:"size:100;not null;" json:"unitid"`
	SyStfTypeuser     int64     `gorm:"size:100;not null;" json:"typeuser"`
	SyStfTypeapprove  int64     `gorm:"size:100;not null;" json:"typeapprove"`
	SyStfPosition     string    `gorm:"size:100;not null;" json:"position"`
	SyStfPart         string    `gorm:"size:100;not null;" json:"part"`
	SyStfPositionuser string    `gorm:"size:100;not null;" json:"positionuser"`
	SyStfTypemini     int64     `gorm:"size:100;not null;" json:"typemini"`
	SyStfTypeusermini int64     `gorm:"size:100;not null;" json:"typeusermini"`
	SyStfUsertype     int64     `gorm:"size:100;not null;" json:"usertype"`
	SyStfStoreid      int64     `gorm:"size:100;not null;" json:"storeid"`
	SyStfUnitidSub    int64     `gorm:"size:100;not null;" json:"unitidsub"`
}

func (SysStf) TableName() string {
	return "sy_stf"
}
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

func (u *SysStf) BeforeSave() error {
	hashedPassword, err := MD5PasswordHash(u.SyStfPassword)
	if err != nil {
		return err
	}
	u.SyStfPassword = string(hashedPassword)
	return nil
}

func (u *SysStf) Prepare() {
	u.SyStfId = 0
	u.SyStfPrefix = html.EscapeString(strings.TrimSpace(u.SyStfPrefix))
	u.SyStfGender = html.EscapeString(strings.TrimSpace(u.SyStfGender))
	u.SyStfFnameeng = html.EscapeString(strings.TrimSpace(u.SyStfFnameeng))
	u.SyStfFnamethai = html.EscapeString(strings.TrimSpace(u.SyStfFnamethai))
	u.SyStfLnameeng = html.EscapeString(strings.TrimSpace(u.SyStfLnameeng))
	u.SyStfLnamethai = html.EscapeString(strings.TrimSpace(u.SyStfLnamethai))
	u.SyStfUsername = html.EscapeString(strings.TrimSpace(u.SyStfUsername))
	u.SyStfPassword = html.EscapeString(strings.TrimSpace(u.SyStfPassword))
	u.SyStfAddress = html.EscapeString(strings.TrimSpace(u.SyStfAddress))
	u.SyStfTelephone = html.EscapeString(strings.TrimSpace(u.SyStfTelephone))
	u.SyStfMobile = html.EscapeString(strings.TrimSpace(u.SyStfMobile))
	u.SyStfEmail = html.EscapeString(strings.TrimSpace(u.SyStfEmail))
	u.SyStfOther = html.EscapeString(strings.TrimSpace(u.SyStfOther))
	u.SyStfPicture = html.EscapeString(strings.TrimSpace(u.SyStfPicture))
	u.SyStfCreby = html.EscapeString(strings.TrimSpace(u.SyStfCreby))
	u.SyStfCredate = time.Now()
	u.SyStfLastdate = time.Now()
	u.SyStfStatus = html.EscapeString(strings.TrimSpace(u.SyStfStatus))
	u.SyStfLogdate = time.Now()
	u.SyStfMasterkey = html.EscapeString(strings.TrimSpace(u.SyStfMasterkey))
	u.SyStfPosition = html.EscapeString(strings.TrimSpace(u.SyStfPosition))
	u.SyStfPart = html.EscapeString(strings.TrimSpace(u.SyStfPart))
	u.SyStfPositionuser = html.EscapeString(strings.TrimSpace(u.SyStfPositionuser))
}

func (u *SysStf) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.SyStfFnameeng == "" {
			return errors.New("Required Nickname")
		}
		if u.SyStfPassword == "" {
			return errors.New("Required Password")
		}
		if u.SyStfEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.SyStfEmail); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.SyStfPassword == "" {
			return errors.New("Required Password")
		}
		if u.SyStfEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.SyStfEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.SyStfFnameeng == "" {
			return errors.New("Required Nickname")
		}
		if u.SyStfPassword == "" {
			return errors.New("Required Password")
		}
		if u.SyStfEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.SyStfEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *SysStf) SaveUser(db *gorm.DB) (*SysStf, error) {

	var err error
	err = db.Table("sy_stf").Create(&u).Error
	if err != nil {
		return &SysStf{}, err
	}
	return u, nil
}

func (u *SysStf) FindAllUsers(db *gorm.DB) (*[]SysStf, error) {
	var err error
	users := []SysStf{}
	err = db.Debug().Model(&SysStf{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]SysStf{}, err
	}
	return &users, err
}

func (u *SysStf) FindUserByID(db *gorm.DB, uid uint32) (*SysStf, error) {
	var err error
	err = db.Debug().Model(SysStf{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &SysStf{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &SysStf{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *SysStf) UpdateAUser(db *gorm.DB, uid uint32) (*SysStf, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&SysStf{}).Where("id = ?", uid).Take(&SysStf{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.SyStfPassword,
			"nickname":  u.SyStfFnameeng,
			"email":     u.SyStfEmail,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &SysStf{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&SysStf{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &SysStf{}, err
	}
	return u, nil
}

func (u *SysStf) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&SysStf{}).Where("id = ?", uid).Take(&SysStf{}).Delete(&SysStf{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
