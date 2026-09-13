package controller

import (
	"SkillManager_API/model"
	"SkillManager_API/service"
	"SkillManager_API/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	log             = utility.Log
	resourceManager = utility.ResourceManager{}
)

// Context is the most important part of gin. It allows us to pass variables between middleware, manage the flow,
// validate the JSON of a request and render a JSON response for example.
type ISkillController interface {
	FindBySkill(context *gin.Context)
	CreateSkill(context *gin.Context)
	DeleteSkill(context *gin.Context)
	AddEmployeeToSkill(context *gin.Context)
	DeleteEmployeeFromSkill(context *gin.Context)
}

type SkillController struct {
	isService service.ISkillService
}

func NewSkillController(isService service.ISkillService) ISkillController {
	return SkillController{
		isService: isService,
	}
}

func (isController SkillController) FindBySkill(context *gin.Context) {
	var skill = context.Param("skillname")
	skillDTO, err := isController.isService.FindBySkill(skill)

	if err != nil {

		log.Error(err.Error())

	} else {
		msg, _ := resourceManager.GetProperty("Controller.FIND_BY_SKILL")
		msg += skill
		log.Info(msg)
		context.JSON(http.StatusOK, skillDTO)
	}
}

func (isController SkillController) CreateSkill(context *gin.Context) {

	var skillDTO model.SkillDTO
	if err := context.BindJSON(&skillDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := isController.isService.CreateSkill(&skillDTO)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		msg, _ := resourceManager.GetProperty("Controller.CREATE_NEW_SKILL")
		msg += skillDTO.SkillName
		log.Info(msg)
		context.JSON(http.StatusCreated, gin.H{"message": msg})
	}

}

func (isController SkillController) DeleteSkill(context *gin.Context) {
	var skill = context.Param("skill")
	err := isController.isService.DeleteSkill(skill)

	if err != nil {
		msg, _ := resourceManager.GetProperty("Controller.DELETE_SKILL_ERROR")
		msg += skill
		log.Error(msg)
		context.JSON(http.StatusBadRequest, gin.H{"error": msg})

	} else {
		msg, _ := resourceManager.GetProperty("Controller.DELETE_SKILL")
		msg += skill
		log.Info(msg)
		context.JSON(http.StatusOK, gin.H{"message": msg})
	}

}

func (isController SkillController) AddEmployeeToSkill(context *gin.Context) {

	var skill = context.Param("skill")

	var employeeDTO model.EmployeeDTO

	if err := context.BindJSON(&employeeDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err := isController.isService.AddEmployeeToSkill(skill, employeeDTO)

	if err != nil {
		msg, _ := resourceManager.GetProperty("Controller.ADD_EMPLOYEE_TO_SKILL_ERROR")
		msg += skill
		log.Info(msg)
		context.JSON(http.StatusBadRequest, gin.H{"error": msg})
	} else {
		msg, _ := resourceManager.GetProperty("Controller.ADD_EMPLOYEE_TO_SKILL")
		msg += skill
		log.Info(msg)
		context.JSON(http.StatusCreated, gin.H{"message": msg})
	}

}

func (isController SkillController) DeleteEmployeeFromSkill(context *gin.Context) {

	skill := context.Param("skill")
	employeeName := context.Param("employeeName")

	_, err := isController.isService.DeleteEmployeeFromSkill(skill, employeeName)

	if err != nil {

		log.Error(err.Error())

		msg, _ := resourceManager.GetProperty("Controller.DELETE_EMPLOYEE_FROM_SKILL_ERROR")
		msg += skill
		context.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	msg, _ := resourceManager.GetProperty("Controller.DELETE_EMPLOYEE_FROM_SKILL")
	msg += skill
	context.JSON(http.StatusOK, gin.H{"message": msg})
}
