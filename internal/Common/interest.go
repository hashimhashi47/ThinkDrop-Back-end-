package domain

type MainInterest struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`

	SubInterests []SubInterest `gorm:"constraint:OnDelete:CASCADE;"`
}

type SubInterest struct {
	ID uint `gorm:"primaryKey"`

	MainInterestID uint         `gorm:"not null;index"`
	MainInterest   MainInterest `gorm:"foreignKey:MainInterestID;references:ID"`

	Posts []Post `gorm:"many2many:post_sub_interests;"`
	Name  string `gorm:"not null"`
}

type Req struct {
	SubInterestIDs []uint `json:"sub_interest_ids"`
}

type MainInterestResponseDTO struct {
	ID           uint                     `json:"id"`
	Name         string                   `json:"name"`
	SubInterests []SubInterestResponseDTO `json:"sub_interests"`
}

type SubInterestResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
