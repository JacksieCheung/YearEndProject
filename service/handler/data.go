package handler

import (
	"YearEndProject/service/log"
	"YearEndProject/service/model"
	"YearEndProject/service/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Data return result
func Data(c *gin.Context) {
	log.Info("api server function called")

	var err error

	// get id
	id := 2019210001

	// 获取信息
	students, err := model.GetStudentList(uint32(id))
	if err != nil {
		SendBadRequest(c, errno.ErrDatabase, nil, err.Error(), GetLine())
	}

	// 并发信息
	// firstItem := getFirstPage(students)
	// secondItem := getSecondPage(students)
	// thirdItem := getThirdPage(strconv.Itoa(id), students)
	// fourthItem := getFourthPage(students)
	// fmt.Println(firstItem, secondItem, thirdItem, fourthItem)

	firstChannel := make(chan *model.FirstItem, 1)
	secondChannel := make(chan *model.SecondItem, 1)
	thirdChannel := make(chan *model.ThirdItem, 1)
	fourthChannel := make(chan *model.FourthItem, 1)

	go goGetFirstPage(firstChannel, students)
	go goGetSecondPage(secondChannel, students)
	go goGetThirdPage(thirdChannel, students, strconv.Itoa(id))
	go goGetFourthPage(fourthChannel, students)

	mission := 4

	page := &model.Page{
		FifthCostMax:         "",
		FifthEarlyDate:       "",
		FifthEarlyRestaurant: "",
		FifthEarlyPlace:      "",
		FifthWellRestaurant:  "",
	}

	for mission > 0 {
		select {

		case item := <-firstChannel:
			page.FirstCostCount = strconv.Itoa(item.CostCount)
			page.FirstCostSum = strconv.FormatFloat(item.CostSum, 'f', 2, 64)
			mission--

		case item := <-secondChannel:
			page.SecondPlace = item.Place
			page.SecondRestaurant = item.Restaurant
			mission--

		case item := <-thirdChannel:
			page.ThirdDate = item.Date
			page.ThirdTime = item.Time
			page.ThirdRestaurant = item.Restaurant
			page.ThirdPlace = item.Place
			page.ThirdPercent = strconv.FormatFloat(item.Percent, 'f', 2, 64)
			mission--

		case item := <-fourthChannel:
			page.FourthCostCount = strconv.Itoa(item.CostCount)
			page.FourthCostSum = strconv.FormatFloat(item.CostSum, 'f', 2, 64)
			page.FourthPlace = item.Place
			page.FourthRestaurant = item.Restaurant
			mission--

		}
	}

	SendResponse(c, nil, page)
}

func goGetFirstPage(firstChannel chan<- *model.FirstItem,
	students []*model.StudentsModel) {
	firstItem := getFirstPage(students)
	firstChannel <- firstItem
}

func goGetSecondPage(secondChannel chan<- *model.SecondItem,
	students []*model.StudentsModel) {
	secondItem := getSecondPage(students)
	secondChannel <- secondItem
}

func goGetThirdPage(thirdChannel chan<- *model.ThirdItem,
	students []*model.StudentsModel,
	id string) {
	thirdItem := getThirdPage(id, students)
	thirdChannel <- thirdItem
}

func goGetFourthPage(fourthChannel chan<- *model.FourthItem,
	students []*model.StudentsModel) {
	fourthItem := getFourthPage(students)
	fourthChannel <- fourthItem
}

func getFirstPage(students []*model.StudentsModel) *model.FirstItem {
	// first
	firstItem := &model.FirstItem{}

	for _, v := range students {
		// first
		firstItem.CostCount++
		firstItem.CostSum += v.Cost
	}

	return firstItem
}

func getSecondPage(students []*model.StudentsModel) *model.SecondItem {
	// second
	date := "2020-08-01"
	secondItem := &model.SecondItem{}
	var timeMin string
	var dateNow string

	for i := 0; i < len(students); i++ {
		if students[i].Date > date {
			dateNow = students[i].Date
			timeMin = students[i].Time
			for ; students[i].Date == dateNow && i < len(students); i++ {
				if timeMin > students[i].Time {
					timeMin = students[i].Time
				}
			}
			secondItem.Restaurant = students[i-1].Restaurant
			secondItem.Place = students[i-1].Place
			break
		}
	}

	return secondItem
}

// 遍历一遍，记录大于九点的记录，比较维护最早时间
func getThirdPage(id string, students []*model.StudentsModel) *model.ThirdItem {
	timeIndex := "09:00:00"
	dateNow := ""
	var count float64

	thirdItem := &model.ThirdItem{Time: students[0].Time}

	var days float64
	if id > "2020210000" {
		days = 122.0
	} else {
		days = 137.0
	}

	for _, v := range students {
		if v.Time < timeIndex {
			if v.Date != dateNow {
				count++
				dateNow = v.Date
			}
			if thirdItem.Time > v.Time {
				thirdItem.Time = v.Time
				thirdItem.Date = v.Date
				thirdItem.Place = v.Place
				thirdItem.Restaurant = v.Restaurant
			}
		}
	}

	thirdItem.Percent = count / days

	return thirdItem
}

// 遍历一遍，维护一个map/切片，地方、次数、总额，最多消费地方
func getFourthPage(students []*model.StudentsModel) *model.FourthItem {
	item := make(map[model.FourthInfo]*model.FourthExtra)
	for _, v := range students {
		key := model.FourthInfo{
			Restaurant: v.Restaurant,
			Place:      v.Place,
		}

		_, ok := item[key]

		if !ok {
			item[key] = &model.FourthExtra{}
		}

		item[key].CostSum = item[key].CostSum + v.Cost
		item[key].CostCount = item[key].CostCount + 1
	}

	fourthItem := &model.FourthItem{}
	for i, v := range item {
		if fourthItem.CostCount < v.CostCount {
			fourthItem.CostCount = v.CostCount
			fourthItem.Restaurant = i.Restaurant
			fourthItem.Place = i.Place
			fourthItem.CostSum = v.CostSum
		}
	}

	return fourthItem
}

// 年度巅峰，直接写死，上数据
