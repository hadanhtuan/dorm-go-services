package model

import (
	"time"
	"github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type LoginSession struct {
	ID        string     `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//foreign key
	UserId string `json:"userId" gorm:"column:user_id"`

	RefreshToken string `json:"refreshToken,omitempty" gorm:"column:refresh_token"`
	SecretKey    string `json:"secretKey,omitempty" gorm:"column:secret_key"`
	DeviceID     string `json:"deviceId,omitempty" gorm:"column:device_id"`
	UserAgent    string `json:"userAgent,omitempty" gorm:"column:user_agent"`
	IpAddress    string `json:"ipAddress,omitempty" gorm:"column:ip_address"`
	ExpiresAt    int64  `json:"expiresAt,omitempty" gorm:"column:expires_at"`
}

var LoginSessionDB = &orm.Instance{
	TableName: "login_session",
	Model:     &LoginSession{},
}

func InitTableLoginSession(db *gorm.DB) {
	db.Table(LoginSessionDB.TableName).AutoMigrate(&LoginSession{})
	LoginSessionDB.ApplyDatabase(db)
}
