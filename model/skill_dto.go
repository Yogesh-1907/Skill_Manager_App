package model

type SkillDTO struct {
	Id        int           `json:"id" binding:"required,gt=0,lte=100"`
	SkillName string        `json:"skillname" binding:"required"`
	Category  string        `json:"category" binding:"required"`
	Employees []EmployeeDTO `json:"Employee" binding:"-"`
}

func FromSkillEntity(skill Skill) SkillDTO {

	var employeeDTOList []EmployeeDTO

	for _, employee := range skill.Employees {
		employeeDTOList = append(employeeDTOList, FromEmployeeEntity(employee))
	}
	return SkillDTO{
		Id:        skill.Id,
		SkillName: skill.SkillName,
		Category:  skill.Category,
		Employees: employeeDTOList,
	}
}

func ToSkillEntity(skillDTO SkillDTO) Skill {

	var employeeList []Employee

	for _, employeeDTO := range skillDTO.Employees {

		employeeList = append(employeeList, ToEmployeeEntity(employeeDTO))

	}
	return Skill{
		Id:        skillDTO.Id,
		SkillName: skillDTO.SkillName,
		Category:  skillDTO.Category,
		Employees: employeeList,
	}
}
