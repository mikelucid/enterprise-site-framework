package models

import "gorm.io/datatypes"

type Payment struct {
	BaseModel
	ExternalPaymentID string         `json:"external_payment_id" gorm:"uniqueIndex;not null"`
	SiteID    string         `json:"site_id" gorm:"index;not null"`
	Amount    float64        `json:"amount" gorm:"not null"`
	Currency  string         `json:"currency" gorm:"size:8;not null"`
	GatewayID string         `json:"gateway_id"`
	Status    string         `json:"status"`
	Metadata  datatypes.JSON `json:"metadata"`
}
