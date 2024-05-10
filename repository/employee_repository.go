package repository

import (
	"SkillManager_API/model"
	"SkillManager_API/utility"
)

type IEmployeeRepository interface {
	FindByName(name string) (model.Employee, error)
	DeleteEmployee(employee model.Employee) (int, error)
}

type EmployeeRepository struct {
	//its purpose is not to hold any data but rather to serve as the receiver for the methods
}

func (employeeRepository EmployeeRepository) FindByName(name string) (model.Employee, error) {

	var err error
	employee := model.Employee{Name: name}

	err = utility.Db.Where(&employee).First(&employee).Error

	if err != nil {
		return employee, err
	}

	return employee, err
}

func (employeeRepository EmployeeRepository) DeleteEmployee(employee model.Employee) (int, error) {

	err := utility.Db.Delete(&employee).Error

	if err != nil {
		return 0, err
	}

	return employee.EmpId, err

}
