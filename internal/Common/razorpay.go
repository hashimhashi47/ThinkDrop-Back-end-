package domain

// -> 1.razorpay response struct(contact)
type RazorpayContact struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Type      string                 `json:"type"`
	Notes     map[string]interface{} `json:"notes,omitempty"`
	CreatedAt int64                  `json:"created_at"`
}

// -> 1.razorpay response struct
type RazorpayContactRequest struct {
	Name  string                 `json:"name"`
	Email string                 `json:"email"`
	Type  string                 `json:"type"`
	Notes map[string]interface{} `json:"notes,omitempty"`
}

// -> 2.razorpay response struct(fund account)
type RazorpayFundAccount struct {
	ID          string              `json:"id"`
	ContactID   string              `json:"contact_id"`
	AccountType string              `json:"account_type"`
	BankAccount *BankAccountDetails `json:"bank_account"`
	CreatedAt   int64               `json:"created_at"`
}
type BankAccountDetails struct {
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
	IFSC          string `json:"ifsc"`
	BankName      string `json:"bank_name,omitempty"`
}

// -> 2.razorpay response struct
type RazorpayFundAccountRequest struct {
	ContactID   string                     `json:"contact_id"`
	AccountType string                     `json:"account_type"`
	BankAccount *BankAccountDetailsRequest `json:"bank_account"`
}

type BankAccountDetailsRequest struct {
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
	IFSC          string `json:"ifsc"`
}

// -> 3.razorpay response struct(payout)
type RazorpayPayout struct {
	ID            string `json:"id"`
	Entity        string `json:"entity"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	Status        string `json:"status"`
	FundAccountID string `json:"fund_account_id"`
	Mode          string `json:"mode"`
	Purpose       string `json:"purpose"`
	Narration     string `json:"narration,omitempty"`

	UTR string `json:"utr,omitempty"`

	Tax           int64  `json:"tax"`
	FailureReason string `json:"failure_reason,omitempty"`
	ReferenceID   string `json:"reference_id"`
	CreatedAt     int64  `json:"created_at"`
}

// -> 3.razorpay request struct
type RazorpayPayoutRequest struct {
	AccountNumber string `json:"account_number"`
	FundAccountID string `json:"fund_account_id"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	Mode          string `json:"mode"`
	Purpose       string `json:"purpose"`
	Narration     string `json:"narration,omitempty"`
	ReferenceID   string `json:"reference_id,omitempty"`
}
