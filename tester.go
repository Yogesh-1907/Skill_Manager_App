package main

import (
	cont "SkillManager_API/controller"
	"SkillManager_API/middleware"
	"SkillManager_API/repository"
	"SkillManager_API/service"
	"SkillManager_API/utility"
	"SkillManager_API/validator"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	logger = utility.Log
)

func main() {

	utility.DbConnect()

	serverPort, _ := utility.ResourceManager{}.GetProperty("server.port")
	fmt.Println(serverPort)
	run(":" + serverPort)
}

func run(serverPort string) {

	ginEngine := gin.Default()
	ginEngine.Use(middleware.ErrorHandler)
	controller := cont.NewSkillController(service.NewSkillService(repository.SkillRepository{}, repository.EmployeeRepository{}, validator.Validator{}))

	ginEngine.GET("/skill/:skillname", controller.FindBySkill)
	ginEngine.POST("/skill", controller.CreateSkill)
	ginEngine.PUT("/employeeskill/:skill", controller.AddEmployeeToSkill)
	ginEngine.DELETE("/employeeskill/:skill/:employeename", controller.DeleteEmployeeFromSkill)
	ginEngine.DELETE("/:skill", controller.DeleteSkill)
	logger.Fatal(ginEngine.Run(serverPort))
}
