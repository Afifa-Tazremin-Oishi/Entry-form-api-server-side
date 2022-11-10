package service

import (
	//"fmt"

	// "strconv"

	"github.com/xyz/model"
	"github.com/xyz/repository"

	//"github.com/xyz/util"

	"github.com/gin-gonic/gin"
)

type CountryRestService struct {
}

// GetAll returns all
func (countryRestService *CountryRestService) GetAll(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	// db := util.CreateConnection()
	// var ab []model.Country
	// result := db.Find(&ab)
	// if result.RowsAffected == 0 {
	// 	c.AbortWithStatus(404)
	// } else {
	// 	c.JSON(200, ab)
	// }
	rep := new(repository.CountryRepository)
	res := rep.Getall()
	c.JSON(res.StatusCode, res)

}

// GetOne returns all info of one
func (countryRestService *CountryRestService) Getbyid(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	// type amarInput struct {
	// 	AmarId int `json:"id"`
	// }
	var id model.Country
	c.ShouldBind(&id)
	// if id <= 0 {
	// 	c.AbortWithStatusJSON(404, `Id can't be zero or negetive\n`)
	// } else {
	// db := util.CreateConnection()
	rep := new(repository.CountryRepository)
	res := rep.GetById(id)
	//var cs model.Country
	//result := db.Where(&model.Country{Id: i.AmarId}).First(&cs)

	c.JSON(res.StatusCode, res)
}

// }

func (countryRestService *CountryRestService) addCountry(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//c.Writer.Header().Set("Content-Type", "application/json")
	var crs model.Country
	c.ShouldBind(&crs)

	// db := util.CreateConnection()

	// re := db.Create(&crs)

	// if re.RowsAffected > 0 {
	// 	c.AbortWithStatusJSON(200, `Data successfully added`)
	// }
	rep := new(repository.CountryRepository)
	res := rep.AddCountry(crs)
	c.JSON(res.StatusCode, res)
}

func (countryRestService *CountryRestService) modifycountry(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	var input model.Country
	//var getFromDb model.Country

	c.ShouldBind(&input)

	//db := util.CreateConnection()
	// res := db.Raw("select * from public.country where id = ?", input.Id).First(&getFromDb)
	// if res.RowsAffected > 0 {
	// 	getFromDb.Names = input.Names
	// 	if input.Names == "" {
	// 		c.AbortWithStatusJSON(404, `Invalid Input`)
	// 	} else {
	// 		db := util.CreateConnection()
	// 		result := db.Where(&model.Country{Id: input.Id}).Save(getFromDb)
	// 		if result.RowsAffected == 0 {
	// 			c.AbortWithStatus(404)
	// 		} else {
	// 			c.JSON(200, input)
	// 		}
	// 	}
	// } else {
	// 	c.AbortWithStatusJSON(409, `Id number do not exist`)
	// }
	repo := new(repository.CountryRepository)
	res := repo.Update(input)
	c.JSON(res.StatusCode, res)

}

// DeleteOne deletes one
func (countryRestService *CountryRestService) deletecountry(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var id model.Country
	c.ShouldBind(&id)
	rep := new(repository.CountryRepository)
	res := rep.Delete(id)
	c.JSON(res.StatusCode, res)

}

//AddRouters add api end points specific to this service
func (countryRestService *CountryRestService) AddRouters(router *gin.Engine) {
	router.GET("/getall", countryRestService.GetAll)
	router.POST("/getbyid", countryRestService.Getbyid)
	router.POST("/addcountry", countryRestService.addCountry)
	router.PUT("/update", countryRestService.modifycountry)
	router.DELETE("/delete", countryRestService.deletecountry)
}
