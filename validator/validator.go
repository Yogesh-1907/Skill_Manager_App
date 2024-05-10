package validator

import (
	"SkillManager_API/model"
	"SkillManager_API/utility"
	"errors"
	"regexp"
)

var (
	log             = utility.Log
	resourceManager = utility.ResourceManager{}
)

type IValidator interface {
	ValidateSkill(skill model.SkillDTO) error
	ValidateEmployee(employee model.EmployeeDTO) error
	ValidateSkillName(skillName string) error
	ValidateSkillCategory(category string) error
	ValidateEmployeeLevel(level string) error
	ValidateEmployeeName(name string) error
	ValidateEmployeeAge(age int) error
}

type Validator struct {
}

// Validate Skill model
func (validator Validator) ValidateSkill(skill model.SkillDTO) error {

	err := validator.ValidateSkillName(skill.SkillName)

	if err != nil {
		return err
	}

	for _, employee := range skill.Employees {

		err = validator.ValidateEmployee(employee)
		if err != nil {
			return err
		}
	}
	return nil
}

// Validate Employee model
func (validator Validator) ValidateEmployee(employee model.EmployeeDTO) error {

	err := validator.ValidateEmployeeName(employee.Name)

	if err != nil {
		return err
	}

	err = validator.ValidateEmployeeLevel(employee.SkillLevel)

	if err != nil {
		return err
	}

	return nil
}

func (validator Validator) ValidateSkillName(skillName string) error {
	success, _ := regexp.MatchString("^[A-Z][a-zA-Z]*$", skillName)

	if !success {
		msg, err := resourceManager.GetProperty("Validator.INVALID_SKILL_NAME")
		if err != nil {
			return err
		}
		log.Error(msg)
		return errors.New(msg)
	}

	return nil
}

func (validator Validator) ValidateSkillCategory(category string) error {
	success, _ := regexp.MatchString("^[A-Z][a-z]+$", category)

	if !success {
		msg, err := resourceManager.GetProperty("Validator.INVALID_SKILL_CATEGORY")
		if err != nil {
			return err
		}
		log.Error(msg)
		return errors.New(msg)
	}

	return nil
}

func (validator Validator) ValidateEmployeeLevel(level string) error {
	success, _ := regexp.MatchString("^[A-Z]+$", level)

	if !success {
		msg, err := resourceManager.GetProperty("Validator.INVALID_SKILL_LEVEL")
		if err != nil {
			return err
		}
		log.Error(msg)
		return errors.New(msg)
	}

	return nil
}

func (validator Validator) ValidateEmployeeName(name string) error {
	success, _ := regexp.MatchString("^[A-Z][a-z]+ [A-Z][a-z]+$", name)
	if !success {
		msg, err := resourceManager.GetProperty("Validator.INVALID_EMPLOYEE_NAME")
		if err != nil {
			log.Error(err.Error())
		}
		log.Error(msg)
		return errors.New(msg)
	}
	return nil
}

func (validator Validator) ValidateEmployeeAge(age int) error {
	if age >= 20 && age <= 60 {
		return nil
	}

	msg, err := resourceManager.GetProperty("Validator.INVALID_EMPLOYEE_AGE")

	if err != nil {
		log.Error(err.Error())
	}
	log.Error(msg)
	return errors.New(msg)
}
