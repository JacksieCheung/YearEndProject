package model

import (
	"github.com/spf13/viper"
)

// Stu save scope of students
var Stu *StuInfo

// StuInfo the structure of Stu
type StuInfo struct {
	Year     string
	StuIDMax string
	StuIDMin string
}

// Init the Stu
func (stu *StuInfo) Init() {
	Stu = &StuInfo{
		Year:     viper.GetString("stu.year"),
		StuIDMax: viper.GetString("stu.stu_id_max"),
		StuIDMin: viper.GetString("stu.stu_id_min"),
	}
}
