package domain

import (
	"time"
)

type DashboardResponse struct {
	TotalUsers   int64       `json:"totalUsers"`
	TotalPosts   int64       `json:"totalPosts"`
	ReportsCount int64       `json:"reportsCount"`
	TopRedeemer  TopRedeemer `json:"topRedeemer"`
}

type TopRedeemer struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type AdminEvent struct {
	Module    string      `json:"module"`
	Type      string      `json:"type"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

type UpdateProfile struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"Password"`
	ImageURL string `json:"imageurl"`
	Role     string `json:"role"`
}

type MainCategory struct {
	CategoryName string `json:"name"`
}

type SubInterest struct {
	SubInterestName string `json:"name"`
}

type CreateSubInterestRequest struct {
	ParentID uint   `json:"parent_id"`
	Name     string `json:"name"`
}

type UpdatePostInterestRequest struct {
	InterestIDs []string `json:"interest_ids"`
}


type UpdateComplaintStatusRequest struct {
	Status string `json:"status" binding:"required"` 
}
