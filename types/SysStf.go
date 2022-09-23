package types

import "time"

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
