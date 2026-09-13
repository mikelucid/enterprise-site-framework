package models

type Creator struct {
	BaseModel
	Name         string  `json:"name"`
	Email        string  `json:"email" gorm:"uniqueIndex"`
	Bio          string  `json:"bio"`
	Earnings     float64 `json:"earnings"`
	IsVerified   bool    `json:"is_verified"`
	Characters   []Character
}
