package model

type Skill struct {
	Id        int        `gorm:"primarykey"`
	SkillName string     `gorm:"column:skill_name;unique;not null;"`
	Category  string     `gorm:"not null;"`
	Employees []Employee `gorm:"constraint:OnDelete:CASCADE;"`
}

func (s Skill) TableName() string {
	return "skillz_db.skill"
}
