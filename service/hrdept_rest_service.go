package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xyz/model"
	"github.com/xyz/repository"
)

type HrService struct{}

func(hrrestservice *HrService)GetAllEmployee(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")	

	rep := new(repository.HrdeptRepository)
	res := rep.GetAllEmployee()
	c.JSON(res.StatusCode, res)

}

// GetOne returns all info of one
func (hrrestservice *HrService) GetEmployeeById(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	var code model.Hrdept
	c.ShouldBind(&code)
	rep := new(repository.HrdeptRepository)
	res := rep.GetEmployeeById(code)
	c.JSON(res.StatusCode, res)
}

func (hrrestservice *HrService) InsertEmployeeService(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var a model.Hrdept
	c.ShouldBind(&a)

	repo := new(repository.HrdeptRepository)
	response := repo.AddEmployee(a)
	c.JSON(response.StatusCode,response)
}

func (hrrestservice *HrService) UpdateEmployeeService(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var a model.Hrdept
	c.ShouldBind(&a)

	repo := new(repository.HrdeptRepository)
	response := repo.Update(a)
	c.JSON(response.StatusCode,response)
}

func (hrrestservice *HrService) DeleteEmployeeService(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var a model.Hrdept
	c.ShouldBind(&a)

	repo := new(repository.HrdeptRepository)
	response := repo.Delete(a)
	c.JSON(response.StatusCode,response)
}

func (hrrestservice *HrService) AddRouters(router *gin.Engine){
	router.POST("/add",hrrestservice.InsertEmployeeService)
	router.GET("/getallemployee",hrrestservice.GetAllEmployee)
	router.POST("/getemployeebyid",hrrestservice.GetEmployeeById)
	router.PATCH("/updateemployee",hrrestservice.UpdateEmployeeService)
	router.DELETE("/deleteemployee",hrrestservice.DeleteEmployeeService)
}