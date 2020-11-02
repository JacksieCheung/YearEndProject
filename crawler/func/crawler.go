package Crawler

import (
	"YearEndProject/crawler/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func Get(db *gorm.DB, stuNum int, Month int) {
	strNum := strconv.Itoa(stuNum)
	strMonth := strconv.Itoa(Month)
	var err error
	var reg, reg2, reg3, reg4 *regexp.Regexp
	var body []byte
	var strBody string
	var res *http.Response
	var req *http.Request
	//爬虫处理
	fmt.Println(strNum + " " + strMonth)
	for {
		url := "http://ecardwx.ccnu.edu.cn/data_chart/ecard_deal/month_deal_list.php?m_xxh=" + strNum

		payload := strings.NewReader("year=2019&month=" + strMonth)

		req, _ = http.NewRequest("POST", url, payload)

		req.Header.Add("Cookie", "PHPSESSID=38c65398f17794628c4227252ff74d93")
		req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; MI 8 UD Build/PKQ1.180729.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/045008 Mobile Safari/537.36 MMWEBID/3648 MicroMessenger/7.0.8.1540(0x27000834) Process/tools NetType/WIFI Language/zh_CN ABI/arm64")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Accept", "*/*")
		req.Header.Add("Cache-Control", "no-cache")
		req.Header.Add("Postman-Token", "dbdb0d9c-833b-447e-8d50-e09335de613b,9c504494-8583-479b-a494-5846ec3eda72")
		req.Header.Add("Host", "ecardwx.ccnu.edu.cn")
		req.Header.Add("Accept-Encoding", "gzip, deflate")
		req.Header.Add("Content-Length", "17")
		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("cache-control", "no-cache")

		res, _ = http.DefaultClient.Do(req)

		if res == nil {
			fmt.Println(strNum + "&" + strMonth + "连接失败")
			continue
		} else {
			fmt.Println(strNum + "&" + strMonth + "连接成功")
			break
		}
	}
	body, _ = ioutil.ReadAll(res.Body)
	res.Body.Close()
	//fmt.Println(res)
	//fmt.Println(string(body))
	strBody = string(body)

	//爬虫结束，正则表达是处理
	reg = regexp.MustCompile("<h2>时间：(.*?)</h2><span>消费：.*?</span><span>地点：后勤集团/饮食中心/.*?</span>|<h2>时间：(.*?)</h2><span>消费：.*?</span><span>地点：后勤集团/商贸中心/.*?</span>")
	result := reg.FindAllStringSubmatch(strBody, -1)
	if len(result) == 0 { //判断是否为有效学号，不是就跳过
		fmt.Println("No Data")
		return
	}
	fmt.Println(strNum)
	//fmt.Printf("%T",result)
	reg2 = regexp.MustCompile("<span>消费：(.*?)元，</span><span>地点：后勤集团/饮食中心/.*?</span>|<span>消费：(.*?)元，</span><span>地点：后勤集团/商贸中心/.*?</span>")
	result2 := reg2.FindAllStringSubmatch(strBody, -1)

	reg3 = regexp.MustCompile("<span>地点：后勤集团/饮食中心/(.*?)/.*?</span>|<span>地点：后勤集团/商贸中心/.*?/(.*?)</span>")
	result3 := reg3.FindAllStringSubmatch(strBody, -1)

	reg4 = regexp.MustCompile("<span>地点：后勤集团/饮食中心/.*?/.*?/(.*?)</span>|<span>地点：后勤集团/商贸中心/(.*?)/.*?</span>|<span>地点：后勤集团/饮食中心/学子中西餐厅/(.*?)</span>|<span>地点：后勤集团/饮食中心/冷库中心/(.*?)</span>")

	result4 := reg4.FindAllStringSubmatch(strBody, -1)

	//检查数据是否正确，错误直接停掉，避免存入错误信息
	if len(result) == len(result2) && len(result) == len(result3) && len(result) == len(result4) {
		fmt.Println("Data Got Correctly")
	} else {
		fmt.Println("Data not Correct")
		panic(strNum + "and" + strMonth)
	}
	//连接数据库
	for i := 0; i < len(result); i = i + 2 {
		/*if err := db.Ping(); err != nil {
			fmt.Println("can not connect MySql,this stuNum is " + strNum + "and time is " + result[i][1])
			panic(err)
		}*/
		if result[i][1] == "" {
			result[i][1] = result[i][2]
		}
		if result2[i][1] == "" {
			result2[i][1] = result2[i][2]
		}
		if result3[i][1] == "" {
			result3[i][1] = result3[i][2]
		}
		if result4[i][1] == "" && result4[i][2] != "" {
			result4[i][1] = result4[i][2]
		}
		if result4[i][1] == "" && result4[i][2] == "" && result4[i][3] != "" {
			result4[i][1] = result4[i][3]
		}
		if result4[i][1] == "" && result4[i][2] == "" && result4[i][3] == "" {
			result4[i][1] = result4[i][4]
		}

		if !db.Where("stunum=? AND time=?", strNum, result[i][1]).Find(&model.Stuinfos{}).RecordNotFound() {
			fmt.Println("Already Existed")
			return
		}
		err = db.Create(&model.Stuinfos{Stunum: strNum, Time: result[i][1], Cost: result2[i][1], Restaurant: result3[i][1], Place: result4[i][1]}).Error
		if err != nil {
			fmt.Println("This stuNum is " + strNum + " time is " + result[i][1])
			fmt.Println(result[i][1], result2[i][1], result3[i][1], result4[i][1])
			fmt.Println(err)
			panic(err)
		}
	}
	fmt.Println("Working Status:Everything OK")
	fmt.Println(strNum + ":" + strMonth + " Finished")
}
