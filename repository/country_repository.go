package repository

import (
	"net/http"

	"github.com/xyz/model"
	"github.com/xyz/util"
)

type CountryRepository struct {
}

func (countryRepo *CountryRepository) Getall() model.ResponseDto {
	var op model.ResponseDto
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	var ab []model.Country
	result := db.Find(&ab)
	if result.RowsAffected == 0 {
		op.Message = "No country info found"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusNotFound
		return op
	}
	type tempOutPut struct {
		Output      []model.Country `json:"output"`
		OutputCount int             `json:"outputCount"`
	}
	var tOutput tempOutPut
	tOutput.Output = ab
	tOutput.OutputCount = len(ab)
	op.Message = "List of countries"
	op.IsSuccess = true
	op.Payload = tOutput
	op.StatusCode = http.StatusOK

	return op
}

func (countryRepo *CountryRepository) GetById(id model.Country) model.ResponseDto {
	var op model.ResponseDto
	if id.Id <= 0 {
		op.Message = "Country id can't be null"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusBadRequest
		return op
	}

	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	//var cs []model.Country
	result := db.Raw("select * from public.country where id = ?", id.Id).First(&id)
	if result.RowsAffected == 0 {
		op.Message = "No country info found"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusNotFound
		return op
	}
	type tempOutput struct {
		Output model.Country `json:"output"`
	}
	var tOutput tempOutput
	tOutput.Output = id
	op.Message = "Country info details found for given criteria"
	op.IsSuccess = true
	op.Payload = tOutput
	op.StatusCode = http.StatusOK
	return op
}

func (countryRepo *CountryRepository) AddCountry(crs model.Country) model.ResponseDto {
	var op model.ResponseDto
	if crs.Id <= 0 {
		op.Message = "Country id can't be null"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusBadRequest
		return op
	}
	if crs.Names == "" {
		op.Message = "Country name can't be null"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusBadRequest
		return op
	}

	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	//_ = db.Create(&crs.Id)
	_ = db.Raw("select coalesce ((max(id) + 1), 1) from public.country").First(&crs.Names)
	result := db.Create(&crs)
	if result.RowsAffected == 0 {
		op.Message = "Country creation failed"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusNotFound
		return op
	}
	type abc struct {
		Op model.Country `json:"op"`
	}
	var xyz abc
	xyz.Op = crs
	op.Message = "Country create successfully"
	op.IsSuccess = true
	op.Payload = xyz
	op.StatusCode = http.StatusCreated
	return op
}

func (countryRepo *CountryRepository) Update(crs model.Country) model.ResponseDto {
	var op model.ResponseDto
	if crs.Id <= 0 {
		op.Message = "Code is invalid"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusBadRequest
		return op
	}
	if crs.Names == "" {
		op.Message = "Name can't be null"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusBadRequest
		return op
	}

	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var input model.Country
	hhh := db.Where(&model.Country{Id: crs.Id}).First(&input)
	if hhh.RowsAffected == 0 {
		op.Message = "this code doesnot exists"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusNotFound
		return op
	}
	// _ = db.Create(&crs.Id)
	// _ = db.Where(&model.Country{Id: crs.Id})
	result := db.Save(&crs)
	if result.RowsAffected == 0 {
		op.Message = "No country info found for given criteria"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusNotFound
		return op
	}
	op.Message = "Country info updated successfully"
	op.IsSuccess = true
	op.Payload = nil
	op.StatusCode = http.StatusOK

	return op
}
func (countryRepository *CountryRepository) Delete(cs model.Country) model.ResponseDto {
	var op model.ResponseDto
	if cs.Id <= 0 {
		op.Message = "Country id can't be null"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusBadRequest
		return op
	}
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	result := db.Where("id = ?", cs.Id).Delete(&cs)
	if result.RowsAffected == 0 {
		op.Message = "No country info found for given criteria"
		op.IsSuccess = false
		op.Payload = nil
		op.StatusCode = http.StatusNotFound
		return op
	}
	op.Message = "Country info deleted successfully"
	op.IsSuccess = true
	op.Payload = nil
	op.StatusCode = http.StatusOK
	return op
}
