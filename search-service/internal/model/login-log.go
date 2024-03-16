package model

import (
	"time"
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
