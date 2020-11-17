package model

// StudentsModel main table
type StudentsModel struct {
	// ID         uint32 `gorm:"column:id"`
	StuID      string `gorm:"column:stu_id"`
	Date       string `gorm:"column:date"`
	Time       string `gorm:"column:time"`
	Cost       string `gorm:"column:cost"`
	Restaurant string `gorm:"column:restaurant"`
	Place      string `gorm:"column:place"`
}

func (u *StudentsModel) TableName() string {
	return "students"
}

// Create ... create table
func (u *StudentsModel) Create() error {
	return DB.Self.Create(&u).Error
}
