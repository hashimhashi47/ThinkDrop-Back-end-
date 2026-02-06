package domain

type RazorpayContact struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Type      string                 `json:"type"` // employee / customer
	Notes     map[string]interface{} `json:"notes,omitempty"`
	CreatedAt int64                  `json:"created_at"`
}

type RazorpayFundAccount struct {
	ID          string              `json:"id"`
	ContactID   string              `json:"contact_id"`
	AccountType string              `json:"account_type"` // "bank_account"
	BankAccount *BankAccountDetails `json:"bank_account"`
	CreatedAt   int64               `json:"created_at"`
}

type BankAccountDetails struct {
	Name          string `json:"name"`                // Account Holder Name
	AccountNumber string `json:"account_number"`      // Bank Account Number
	IFSC          string `json:"ifsc"`                // Bank IFSC Code
	BankName      string `json:"bank_name,omitempty"` // Optional
}

type RazorpayPayout struct {
	ID            string `json:"id"`
	Entity        string `json:"entity"` // "payout"
	Amount        int64  `json:"amount"` // in paise
	Currency      string `json:"currency"`
	Status        string `json:"status"` // processing, processed, failed
	FundAccountID string `json:"fund_account_id"`
	Mode          string `json:"mode"`    // IMPS, NEFT, RTGS
	Purpose       string `json:"purpose"` // payout, salary, vendor_payment
	Narration     string `json:"narration,omitempty"`
	CreatedAt     int64  `json:"created_at"`
}
