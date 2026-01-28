package domain

type SubInterestResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type MainInterestResponse struct {
	ID           uint                  `json:"id"`
	Name         string                `json:"name"`
	SubInterests []SubInterestResponse `json:"sub_interests"`
}
