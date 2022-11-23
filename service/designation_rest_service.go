package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xyz/model"
	"github.com/xyz/repository"
)

type DesignationRestService struct {
}

func (designationRestService *DesignationRestService) GetAllDesignation(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	rep := new(repository.DesignationRepository)
	res := rep.GetAllDesignation()
	c.JSON(res.StatusCode, res)
}

func (designationRestService *DesignationRestService) GetById(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var id model.Desig
	c.ShouldBind(&id)
	rep1 := new(repository.DesignationRepository)
	res1 := rep1.GetById(id)
	c.JSON(res1.StatusCode, res1)
}

func (designationRestService *DesignationRestService) AddDesignation(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var obj model.Desig
	c.ShouldBind(&obj)
	rep := new(repository.DesignationRepository)
	res := rep.AddDesignation(obj)
	c.JSON(res.StatusCode, res)
}

func (designationRestService *DesignationRestService) UpdateDesignation(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var obj1 model.Desig
	c.ShouldBind(&obj1)
	repo := new(repository.DesignationRepository)
	result2 := repo.UpdateDesignation(obj1)
	c.JSON(result2.StatusCode, result2)
}

func (designationRestService *DesignationRestService) DeleteDesignation(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var id model.Desig
	c.ShouldBind(&id)
	repo := new(repository.DesignationRepository)
	result3 := repo.Delete(id)
	c.JSON(result3.StatusCode, result3)


}

func (designationRestService *DesignationRestService) AddRouters(router *gin.Engine) {
	router.GET("/getalldesignation", designationRestService.GetAllDesignation)
	router.POST("/getdesigbyid", designationRestService.GetById )
	router.POST("/adddasignation", designationRestService.AddDesignation)
	router.PATCH("/updatedesigntion", designationRestService.UpdateDesignation)
	router.DELETE("/deletedesigntion", designationRestService.DeleteDesignation)
}
