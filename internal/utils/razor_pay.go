package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	domain "thinkdrop-backend/internal/Common"
)

func razorpayAuth(req *http.Request) {
	req.SetBasicAuth(
		os.Getenv("RAZORPAY_KEY_ID"),
		os.Getenv("RAZORPAY_KEY_SECRET"),
	)
}

// -> This is the request is used to create a Contact id
func RazorpayContact(User domain.User) (domain.RazorpayContact, error) {

	Url := "https://api.razorpay.com/v1/contacts"

	Payload := domain.RazorpayContactRequest{
		Name:  User.FullName,
		Email: User.Email,
		Type:  "employee",
		Notes: map[string]interface{}{
			"user_id": User.ID,
		},
	}

	PayloadBytes, _ := json.Marshal(Payload)

	req, _ := http.NewRequest("POST", Url, bytes.NewBuffer(PayloadBytes))
	razorpayAuth(req)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return domain.RazorpayContact{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return domain.RazorpayContact{}, fmt.Errorf("Razorpay API error: %s", string(bodyBytes))
	}

	var result domain.RazorpayContact
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

// ->This request is used used to razorpay fund account and the id
func RazorpayFundAccount(Account domain.BankAccount, Username string) (domain.RazorpayFundAccount, error) {
	Url := "https://api.razorpay.com/v1/fund_accounts"

	payload := domain.RazorpayFundAccountRequest{
		ContactID:   Account.RazorpayContactID,
		AccountType: "bank_account",
		BankAccount: &domain.BankAccountDetailsRequest{
			Name:          Username,
			AccountNumber: Account.AccountNumber,
			IFSC:          Account.IFSCCode,
		},
	}

	payloadBytes, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", Url, bytes.NewBuffer(payloadBytes))
	razorpayAuth(req)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return domain.RazorpayFundAccount{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return domain.RazorpayFundAccount{}, fmt.Errorf("Razorpay API error: %s", string(bodyBytes))
	}

	var Result domain.RazorpayFundAccount
	json.NewDecoder(resp.Body).Decode(&Result)

	return Result, nil
}

// -> razorpay payout setting
func RazorpayPayout(Amount int64, FundAccoutID string) (domain.RazorpayPayout, error) {
	Url := "https://api.razorpay.com/v1/payouts"

	Payload := domain.RazorpayPayoutRequest{
		AccountNumber: os.Getenv("RAZORPAY_SOURCE_ACCOUNT"),
		FundAccountID: FundAccoutID,
		Amount:        Amount,
		Currency:      "INR",
		Mode:          "IMPS",
		Purpose:       "payout",
		Narration:     "ThinkDrop Reward",
	}

	PayloadBytes, _ := json.Marshal(Payload)

	req, _ := http.NewRequest("POST", Url, bytes.NewBuffer(PayloadBytes))
	razorpayAuth(req)
	req.Header.Set("Content-Type", "application/json")

	Client := &http.Client{}
	resp, err := Client.Do(req)

	if err != nil {
		return domain.RazorpayPayout{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return domain.RazorpayPayout{}, fmt.Errorf("Razorpay API error: %s", string(bodyBytes))
	}

	var result domain.RazorpayPayout
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

// -> get the payout details
func GetRazorpayPayout(payoutID string) (domain.RazorpayPayout, error) {
	url := fmt.Sprintf("https://api.razorpay.com/v1/payouts/%s", payoutID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return domain.RazorpayPayout{}, err
	}

	razorpayAuth(req)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return domain.RazorpayPayout{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return domain.RazorpayPayout{}, fmt.Errorf("razorpay error: %s", string(body))
	}

	var payout domain.RazorpayPayout
	if err := json.NewDecoder(resp.Body).Decode(&payout); err != nil {
		return domain.RazorpayPayout{}, err
	}

	return payout, nil
}
