package model

// Data Transfer Struct for Employee which contains JSON and BINDING struct tags for mapping with JSON Request/Response
// Bodies for various HTTP Requests and their corresponding responses
type EmployeeDTO struct {
	EmpId      int    `json:"id"`
	Name       string `json:"employeeName" binding:"required"`
	Age        int    `json:"age" binding:"required"`
	SkillLevel string `json:"skillLevel" binding:"required"`
	SkillId    int    `json:"skill_id" binding:"required,gte=1,lte=100"`
}

func FromEmployeeEntity(employee Employee) EmployeeDTO {
	return EmployeeDTO{
		EmpId:      employee.EmpId,
		Name:       employee.Name,
		Age:        employee.Age,
		SkillLevel: employee.SkillLevel,
		SkillId:    employee.SkillId,
	}
}

func ToEmployeeEntity(employeeDTO EmployeeDTO) Employee {
	return Employee{
		EmpId:      employeeDTO.EmpId,
		Name:       employeeDTO.Name,
		Age:        employeeDTO.Age,
		SkillLevel: employeeDTO.SkillLevel,
		SkillId:    employeeDTO.SkillId,
	}
}
