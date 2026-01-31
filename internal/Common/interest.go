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

	Posts []Post `gorm:"foreignKey:SubInterestID"`
	Name  string        `gorm:"not null"`
}

type Req struct {
	SubInterestIDs []uint `json:"sub_interest_ids"`
}
