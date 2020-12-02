package model

type FirstItem struct {
	CostCount int
	CostSum   float64
}

type SecondItem struct {
	Restaurant string
	Place      string
}

type ThirdItem struct {
	Date       string
	Time       string
	Restaurant string
	Place      string
	Percent    float64
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

type Page struct {
	FirstCostCount       string
	FirstCostSum         string
	SecondRestaurant     string
	SecondPlace          string
	ThirdDate            string
	ThirdTime            string
	ThirdRestaurant      string
	ThirdPlace           string
	ThirdPercent         string
	FourthRestaurant     string
	FourthPlace          string
	FourthCostSum        string
	FourthCostCount      string
	FifthCostMax         string
	FifthEarlyDate       string
	FifthEarlyRestaurant string
	FifthEarlyPlace      string
	FifthPercent         string
	FifthWellRestaurant  string
}
