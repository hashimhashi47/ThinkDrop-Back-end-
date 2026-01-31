package domain

type Userdata struct {
	Name    string `json:"name"`
	Intrest []Intrests
}

type Intrests struct {
	MainInterestID   uint   `json:"maininterest_id"`
	MainInterestName string `json:"maininterest_name"`
	SubIntretsID     uint   `json:"subinterest_id"`
	SubIntrestName   string `json:"subinterest_name"`
}
