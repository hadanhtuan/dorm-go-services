package model

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
	orm "github.com/hadanhtuan/go-sdk/db/orm"
)

type LoginLog struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//foreign key
	UserId uuid.UUID `json:"userId" gorm:"column:user_id"`

	UserAgent string `json:"userAgent,omitempty" gorm:"column:user_agent"`
	IpAddress string `json:"ipAddress,omitempty" gorm:"column:ip_address"`
	DeviceID   string `json:"deviceId,omitempty" gorm:"column:device_id"`
}

var LoginLogDB = &orm.Instance{
	TableName: "login_log",
	Model:     &LoginLog{},
}

func InitTableLoginLog(db *gorm.DB) {
	LoginLogDB.ApplyDatabase(db)
}
