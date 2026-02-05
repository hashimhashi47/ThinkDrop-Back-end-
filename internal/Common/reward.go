package domain

import "time"

type Wallet struct {
	ID       uint   `gorm:"primaryKey"`
	WalletID string `gorm:"uniqueIndex;not null"`
	UserID   uint   `gorm:"uniqueIndex;not null"`

	IsWalletActive string `gorm:"default:inactive"`

	PointsAvailable int `gorm:"default:0"`
	TotalLikes      int `gorm:"default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type BankAccount struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"index"`

	AccountHolderName string
	AccountNumber     string
	IFSCCode          string
	BankName          string

	RazorpayContactID     string
	RazorpayFundAccountID string

	IsVerified bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Withdrawal struct {
	ID uint `gorm:"primaryKey"`

	UserID        uint `gorm:"index"`
	BankAccountID uint

	PointsUsed int
	AmountINR  int // store ₹ or paise (decide once)

	Status string // pending, processing, success, failed

	RazorpayPayoutID string
	FailureReason    string

	CreatedAt time.Time
	UpdatedAt time.Time
}
