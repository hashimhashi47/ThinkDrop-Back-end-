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

func RazorpayContact(User domain.User) (domain.RazorpayContact, error) {

	KeyID := os.Getenv("RAZORPAY_KEY_ID")
	keySecret := os.Getenv("RAZORPAY_KEY_SECRET")
	Url := "https://api.razorpay.com/v1/contacts"

	Payload := domain.RazorpayContact{
		Name:  User.FullName,
		Email: User.Email,
		Type:  "employee",
		Notes: map[string]interface{}{
			"user_id": User.ID,
		},
	}

	PayloadBytes, _ := json.Marshal(Payload)

	req, _ := http.NewRequest("POST", Url, bytes.NewBuffer(PayloadBytes))
	req.SetBasicAuth(KeyID, keySecret)
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
