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

// TableName return table name
func (u *StudentsModel) TableName() string {
	return "students"
}

// Create ... create table
func (u *StudentsModel) Create() error {
	return DB.Self.Create(&u).Error
}

// GetStudentsRecord find record before insert
func GetStudentsRecord(stuID, date, time string) error {
	d := DB.Self.Table("students").Where("stu_id=? AND date=? AND time=?", stuID, date, time).First(&StudentsModel{})
	return d.Error
}
