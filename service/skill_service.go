package service

import (
	"SkillManager_API/model"
	"SkillManager_API/repository"
	"SkillManager_API/utility"
	"SkillManager_API/validator"
	"errors"
)

var (
	log             = utility.Log
	resourceManager = utility.ResourceManager{}
)

type ISkillService interface {
	CreateSkill(skill *model.SkillDTO) error
	FindBySkill(skillName string) (model.SkillDTO, error)
	AddEmployeeToSkill(skillName string, employee model.EmployeeDTO) (int, error)
	DeleteEmployeeFromSkill(skillName string, name string) (int, error)
	DeleteSkill(skillName string) error
}

type SkillService struct {
	skillRepo    repository.ISkillRepository
	employeeRepo repository.IEmployeeRepository
	validator    validator.IValidator
}

func NewSkillService(skillRepo repository.ISkillRepository, employeeRepo repository.IEmployeeRepository, validator validator.IValidator) ISkillService {
	return SkillService{
		skillRepo:    skillRepo,
		employeeRepo: employeeRepo,
		validator:    validator,
	}
}

func (skillService SkillService) CreateSkill(skillDTO *model.SkillDTO) error {

	err := skillService.validator.ValidateSkill(*skillDTO)

	if err != nil {
		return err
	}

	skill := model.ToSkillEntity(*skillDTO)

	error := skillService.skillRepo.CreateSkill(&skill)

	if error != nil {
		msg, err := resourceManager.GetProperty("Service.CREATE_SKILL_ERROR")

		if err != nil {
			log.Error(err.Error())
		}

		log.Error(msg)
		return errors.New(msg)
	}

	return error
}

func (skillService SkillService) FindBySkill(skillName string) (model.SkillDTO, error) {
	var skill model.Skill

	if err := skillService.validator.ValidateSkillName(skillName); err != nil {
		return model.SkillDTO{}, err
	}

	// Fetch skill from repository
	skill, err := skillService.skillRepo.FindBySkill(skillName)
	if err != nil {
		return model.SkillDTO{}, err
	}

	return model.FromSkillEntity(skill), nil
}

func (skillService SkillService) AddEmployeeToSkill(skillName string, employeeDTO model.EmployeeDTO) (int, error) {

	var err error

	var empId int

	var skill model.Skill

	err = skillService.validator.ValidateSkillName(skillName)
	if err != nil {
		return empId, err
	}

	err = skillService.validator.ValidateEmployee(employeeDTO)
	if err != nil {
		return empId, err
	}

	skill, err = skillService.skillRepo.FindBySkill(skillName)

	if err != nil {
		msg, error := resourceManager.GetProperty("Service.FIND_SKILL_ERROR")

		if error != nil {
			log.Error(error.Error())
		}

		log.Error(msg)
		return empId, errors.New(msg)
	}

	employee := model.ToEmployeeEntity(employeeDTO)

	empId, err = skillService.skillRepo.AddEmployeeToSkill(skill, employee)

	if err != nil {
		msg, error := resourceManager.GetProperty("Service.REGISTER_SKILL_ERROR")

		if error != nil {
			log.Error(error.Error())
		}
		log.Error(msg)

		return empId, errors.New(msg)
	}

	return empId, err

}

func (skillService SkillService) DeleteEmployeeFromSkill(skillName string, name string) (int, error) {

	var err error
	var empId int
	var skill model.Skill

	err = skillService.validator.ValidateSkillName(skillName)
	if err != nil {
		return empId, err
	}
	err = skillService.validator.ValidateEmployeeName(name)
	if err != nil {
		return empId, err
	}

	skill, err = skillService.skillRepo.FindBySkill(skillName)
	if err != nil {
		msg, error := resourceManager.GetProperty("Service.FIND_EMPLOYEE_ERROR")

		if error != nil {

			log.Error(error.Error())

		}
		log.Error(msg)
		return empId, errors.New(msg)
	}

	for _, employee := range skill.Employees {

		if employee.Name == name {

			employee, _ := skillService.employeeRepo.FindByName(name)
			empId, err = skillService.employeeRepo.DeleteEmployee(employee)

		}
	}
	return empId, err
}

func (skillService SkillService) DeleteSkill(skillName string) error {

	var err error
	var skill model.Skill

	err = skillService.validator.ValidateSkillName(skillName)
	if err != nil {
		return err
	}

	skill, err = skillService.skillRepo.FindBySkill(skillName)
	if err != nil {
		msg, error := resourceManager.GetProperty("Service.FIND_EMPLOYEE_ERROR")

		if error != nil {
			log.Error(error.Error())
		}

		log.Error(msg)
		return errors.New(msg)
	}

	if err = skillService.skillRepo.DeleteEmployeeAssociation(skill); err != nil {
		msg, error := resourceManager.GetProperty("Service.DELETE_ASSOCIATION_ERROR")

		if error != nil {
			log.Error(error.Error())
		}

		log.Error(msg)
		return errors.New(msg)
	}

	return err
}
