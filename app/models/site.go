package models

import "gorm.io/datatypes"

type SiteStatus string

const (
	SitePending       SiteStatus = "pending"
	SiteRunning       SiteStatus = "running"
	SiteDecommissioned SiteStatus = "decommissioned"
)

type Site struct {
	BaseModel
	Name       string         `json:"name" gorm:"not null"`
	Domain     string         `json:"domain" gorm:"uniqueIndex;not null"`
	TemplateID *string        `json:"template_id"`
	Status     SiteStatus     `json:"status" gorm:"type:varchar(30);default:'pending'"`
	Config     datatypes.JSON `json:"config"`
	Budgets    []Budget       `json:"budgets"`
	Campaigns  []Campaign     `json:"campaigns"`
	Payments   []Payment      `json:"payments"`
}

type Budget struct{ BaseModel; SiteID string }
type Campaign struct{ BaseModel; SiteID string }
