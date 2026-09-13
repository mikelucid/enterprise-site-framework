package models

import "gorm.io/datatypes"

type AuditLog struct {
	BaseModel
	ActionType   string         `json:"action_type" gorm:"not null"`
	UserID       string         `json:"user_id" gorm:"index"`
	ResourceType string         `json:"resource_type" gorm:"index"`
	ResourceID   string         `json:"resource_id" gorm:"index"`
	Details      datatypes.JSON `json:"details"`
}
