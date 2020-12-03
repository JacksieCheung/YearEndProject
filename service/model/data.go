package model

type Request struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type FirstItem struct {
	CostCount int
	CostSum   float64
	Status    string
}

type SecondItem struct {
	Restaurant string
	Place      string
}

type ThirdItem struct {
	Date          string
	Time          string
	Restaurant    string
	Place         string
	Percent       int
	TimeStatus    string
	PercentStatus string
	FifthStatus   string
}

type FourthInfo struct {
	Restaurant string
	Place      string
}

type FourthExtra struct {
	CostSum   float64
	CostCount int
}

type FourthItem struct {
	Restaurant string
	Place      string
	CostSum    float64
	CostCount  int
}

type Response struct {
	FirstCostCount       string `json:"first_cost_count"`
	FirstCostSum         string `json:"first_cost_sum"`
	FirstStatus          string `json:"first_status"`
	SecondRestaurant     string `json:"second_restaurant"`
	SecondPlace          string `json:"second_place"`
	ThirdDate            string `json:"third_date"`
	ThirdTime            string `json:"third_time"`
	ThirdRestaurant      string `json:"third_restaurant"`
	ThirdPlace           string `json:"third_place"`
	ThirdPercent         string `json:"third_percent"`
	ThirdTimeStatus      string `json:"third_time_status"`
	ThirdPercentStatus   string `json:"third_percent_status"`
	FourthRestaurant     string `json:"fourth_restaurant"`
	FourthPlace          string `json:"fourth_place"`
	FourthCostSum        string `json:"fourth_cost_sum"`
	FourthCostCount      string `json:"fourth_cost_count"`
	FifthCostMax         string `json:"fifth_cost_max"`
	FifthEarlyDate       string `json:"fifth_early_date"`
	FifthEarlyRestaurant string `json:"fifth_early_restaurant"`
	FifthEarlyPlace      string `json:"fifth_early_place"`
	FifthPercent         string `json:"fifth_percent"`
	FifthWellRestaurant  string `json:"fifth_well_retaurant"`
	FifthStatus          string `json:"fifth_status"`
}
