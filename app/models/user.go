package models

import "github.com/lib/pq"

type UserStatus string

const (
	UserActive    UserStatus = "active"
	UserInactive  UserStatus = "inactive"
	UserSuspended UserStatus = "suspended"
)

type User struct {
	BaseModel
	Email        string         `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string         `json:"-" gorm:"not null"`
	Roles        pq.StringArray `json:"roles" gorm:"type:text[]"`
	Status       UserStatus     `json:"status" gorm:"type:varchar(30);default:'active'"`
}
