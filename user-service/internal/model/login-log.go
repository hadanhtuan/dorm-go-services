package model

import (
	"time"

	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type LoginLog struct {
	ID        string     `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//foreign key
	UserId string `json:"userId" gorm:"column:user_id"`

	ExpiresAt time.Time `json:"expiresAt,omitempty" gorm:"column:expires_at"`
	IsLogout  bool      `json:"isLogout,omitempty" gorm:"column:is_logout"`
	UserAgent string    `json:"userAgent,omitempty" gorm:"column:user_agent"`
	IpAddress string    `json:"ipAddress,omitempty" gorm:"column:ip_address"`
	DeviceID  string    `json:"deviceId,omitempty" gorm:"column:device_id"`
}

var LoginLogDB = &orm.Instance{
	TableName: "login_log",
	Model:     &LoginLog{},
}

func InitTableLoginLog(db *gorm.DB) {
	db.Table(LoginLogDB.TableName).AutoMigrate(&LoginLog{})
	LoginLogDB.ApplyDatabase(db)
}
