package repository

import (
	"SkillManager_API/model"
	"SkillManager_API/utility"
)

type ISkillRepository interface {
	CreateSkill(skill *model.Skill) error
	FindBySkill(skillName string) (model.Skill, error)
	AddEmployeeToSkill(skill model.Skill, employee model.Employee) (int, error)
	DeleteSkill(skill model.Skill) error
	DeleteEmployeeAssociation(skill model.Skill) error
}

type SkillRepository struct {
	//its purpose is not to hold any data but rather to serve as the receiver for the methods
}

func (skillRepository SkillRepository) CreateSkill(skill *model.Skill) error {

	err := utility.Db.Create(&skill).Error

	if err != nil {
		return err
	}

	return nil
}

func (skillRepository SkillRepository) FindBySkill(skillName string) (model.Skill, error) {
	var skill model.Skill
	var err error
	if err = utility.Db.Where("skill_name = ?", skillName).First(&skill).Error; err != nil {
		return skill, err
	}

	if err = utility.Db.Model(&skill).Association("Employees").Find(&skill.Employees); err != nil {
		return skill, err
	}
	return skill, nil
}

func (skillRepository SkillRepository) AddEmployeeToSkill(skill model.Skill, employee model.Employee) (int, error) {

	err := utility.Db.Model(&skill).Association("Employees").Append(&employee)

	if err != nil {
		return 0, err
	}

	return employee.EmpId, nil
}

func (skillRepository SkillRepository) DeleteSkill(skill model.Skill) error {

	err := utility.Db.Delete(&skill).Error

	if err != nil {
		return err
	}

	return nil
}

func (skillRepository SkillRepository) DeleteEmployeeAssociation(skill model.Skill) error {

	err := utility.Db.Model(&skill).Association("Employees").Delete(&skill.Employees)

	if err != nil {
		return err
	}

	return nil
}
