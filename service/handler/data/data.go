package data

import (
	. "YearEndProject/service/handler"
	"YearEndProject/service/log"
	"YearEndProject/service/model"
	"YearEndProject/service/pkg/errno"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Data return result
func Data(c *gin.Context) {
	log.Info("api server function called")

	var err error

	// get id
	id := c.MustGet("id").(int)

	// 获取信息
	students, err := model.GetStudentList(uint32(id))
	if err != nil {
		SendBadRequest(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	page := &model.Response{
		FifthCostMax:         model.FifthInfo.CostMax,
		FifthEarlyDate:       model.FifthInfo.EarlyDate,
		FifthEarlyRestaurant: model.FifthInfo.EarlyRestaurant,
		FifthEarlyPlace:      model.FifthInfo.EarlyPlace,
		FifthWellRestaurant:  model.FifthInfo.WellRestaruant,
	}

	// 并发信息
	firstChannel := make(chan *model.FirstItem, 1)
	secondChannel := make(chan *model.SecondItem, 1)
	thirdChannel := make(chan *model.ThirdItem, 1)
	fourthChannel := make(chan *model.FourthItem, 1)

	defer close(firstChannel)
	defer close(secondChannel)
	defer close(thirdChannel)
	defer close(fourthChannel)

	go goGetFirstPage(firstChannel, students)
	go goGetSecondPage(secondChannel, students)
	go goGetThirdPage(thirdChannel, students, strconv.Itoa(id))
	go goGetFourthPage(fourthChannel, students)

	mission := 4

	for mission > 0 {
		select {

		case item := <-firstChannel:
			page.FirstCostCount = strconv.Itoa(item.CostCount)
			page.FirstCostSum = strconv.FormatFloat(item.CostSum, 'f', 2, 64)
			page.FirstStatus = item.Status
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
			page.ThirdPercent = strconv.Itoa(item.Percent)
			page.ThirdTimeStatus = item.TimeStatus
			page.ThirdPercentStatus = item.PercentStatus
			page.FifthStatus = item.FifthStatus
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

	var buff strings.Builder
	if firstItem.CostSum >= 5000.0 {
		buff.WriteString("可以购入")
		buff.WriteString(strconv.Itoa(int(firstItem.CostSum) / 300))
		buff.WriteString("双匡威帆布鞋")
	} else if firstItem.CostSum >= 4000 {
		buff.WriteString("可以买")
		buff.WriteString(strconv.Itoa(int(firstItem.CostSum) / 300))
		buff.WriteString("支YSL水光唇釉613")
	} else if firstItem.CostSum >= 3000 {
		buff.WriteString("相当于吃")
		buff.WriteString(strconv.Itoa(int(firstItem.CostSum) / 100))
		buff.WriteString("次海底捞")
	} else if firstItem.CostSum >= 2000 {
		buff.WriteString("可以纵享")
		buff.WriteString(strconv.Itoa(int(firstItem.CostSum) / 100))
		buff.WriteString("年视频网站VIP")
	} else if firstItem.CostSum >= 1000 {
		buff.WriteString("相当于重刷")
		buff.WriteString(strconv.Itoa(int(firstItem.CostSum) / 50))
		buff.WriteString("次四级英语")
	} else if firstItem.CostSum >= 500 {
		buff.WriteString("足够喝")
		buff.WriteString(strconv.Itoa(int(firstItem.CostSum) / 12))
		buff.WriteString("次烧仙草奶茶")
	} else {
		buff.WriteString("可以吃")
		buff.WriteString(strconv.Itoa(int(firstItem.CostSum) / 4))
		buff.WriteString("次食堂热干面")
	}

	firstItem.Status = buff.String()

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

	thirdItem.Percent = int(count / days * 100)

	if thirdItem.Time < "07:00:00" {
		thirdItem.TimeStatus = "清晨的朝霞美丽吗？"
	} else if thirdItem.Time < "10:00:00" {
		thirdItem.TimeStatus = "还记得你那天早餐吃的是什么吗？"
	} else {
		thirdItem.TimeStatus = "同志仍需努力，争取早起！"
	}

	if thirdItem.Percent >= 80 {
		thirdItem.PercentStatus = "早起冠军就是你，做最勤奋的学习人！"
		thirdItem.FifthStatus = "早起模范标兵"
	} else if thirdItem.Percent >= 60 {
		thirdItem.PercentStatus = "早起优等生，请继续保持！"
		thirdItem.FifthStatus = "追逐太阳的学习人"
	} else if thirdItem.Percent >= 40 {
		thirdItem.PercentStatus = "努努力，早起及格线就在眼前！"
		thirdItem.FifthStatus = "挣扎温饱线的学习人"
	} else {
		thirdItem.PercentStatus = "努力早起吧！努力的学习人需要好身体！"
		thirdItem.FifthStatus = "饿肚子的白日梦想家"
	}

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
