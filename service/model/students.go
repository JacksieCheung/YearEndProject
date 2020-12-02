package model

// StudentsModel main table
type StudentsModel struct {
	ID         uint32  `gorm:"column:id"`
	StuID      string  `gorm:"column:stu_id"`
	Date       string  `gorm:"column:date"`
	Time       string  `gorm:"column:time"`
	Cost       float64 `gorm:"column:cost"`
	Restaurant string  `gorm:"column:restaurant"`
	Place      string  `gorm:"column:place"`
}

// TableName return table name
func (u *StudentsModel) TableName() string {
	return "students"
}

// GetStudentList 获得信息列表
func GetStudentList(id uint32) ([]*StudentsModel, error) {
	students := make([]*StudentsModel, 0)

	if err := DB.Self.Table("students").Where("stu_id=?", id).Order("date asc").Scan(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
