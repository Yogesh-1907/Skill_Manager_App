package model

//Entity Struct for Employee which contains GORM struct tags for ORM mapping with employee table in database
type Employee struct {
	EmpId      int    `gorm:"primary_key"`
	Name       string `gorm:"column:employee_name;notNull;"`
	Age        int    `gorm:"notNull"`
	SkillLevel string `gorm:"notNull"`
	SkillId    int    `gorm:"foreignKey:SkillId"`
}

func (s Employee) TableName() string {
	return "skillz_db.employee"
}
