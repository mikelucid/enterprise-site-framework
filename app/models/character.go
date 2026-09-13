package models

type Character struct {
	BaseModel
	CharacterID  string  `json:"character_id" gorm:"uniqueIndex;not null"`
	CreatorID    string  `json:"creator_id" gorm:"index;not null"`
	Name         string  `json:"name" gorm:"not null"`
	Description  string  `json:"description"`
	Status       string  `json:"status" gorm:"default:'active'"`
	Type         string  `json:"type"`
	LicensingFee float64 `json:"licensing_fee"`
}
