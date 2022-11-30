package repository

import (
	"net/http"
	"strings"

	"github.com/xyz/model"
	"github.com/xyz/util"
)

type DesignationRepository struct {
}

func (designationrepo *DesignationRepository) GetAllDesignation() model.ResponseDto {
	var output model.ResponseDto
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var obj []model.Desig
	result := db.Order("code").Find(&obj)
	if result.RowsAffected == 0 {
		output.Message = "No country info found"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusNotFound
		return output
	}
	type tempOutPut struct {
		T_output    []model.Desig `json:"output"`
		OutputCount int           `json:"outputCount"`
	}
	var tOutput tempOutPut
	tOutput.T_output = obj
	tOutput.OutputCount = len(obj)
	output.Message = "List of designations"
	output.IsSuccess = true
	output.Payload = tOutput
	output.StatusCode = http.StatusOK

	return output
}
func (designationrepo *DesignationRepository) GetById(id model.Desig) model.ResponseDto {
	var output model.ResponseDto
	if id.Code <= 0 {
		output.Message = "Designation can't be null"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output
	}

	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	result := db.Raw("select * from public.Desig where code = ?", id.Code).First(&id)
	if result.RowsAffected == 0 {
		output.Message = "No country info found"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusNotFound
		return output
	}
	type tempOutput struct {
		Output model.Desig `json:"output"`
	}
	var tOutput tempOutput
	tOutput.Output = id
	output.Message = "Designation info details found for given criteria"
	output.IsSuccess = true
	output.Payload = tOutput
	output.StatusCode = http.StatusOK
	return output
}

func (designationrepo *DesignationRepository) AddDesignation(c model.Desig) model.ResponseDto {
	var output model.ResponseDto
	if c.Code <= 0 {
		output.Message = "Invalid code"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output

	}
	if c.Designation == "" {
		output.Message = "Name can't be null"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output

	}
	if c.Sdesignation == "" {
		output.Message = "Dept can't be null"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output
	}
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//result := db.Raw("Select * from public.desig where code =?", c.Code).First(&c)
	result := db.Where(&model.Desig{Code: c.Code}).First(&c)
	if result.RowsAffected != 0 {
		output.Message = "Department Code is already exist"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusConflict
		return output
	}
	result1 := db.Raw("Select * from public.desig where lower(designation) =? and lower(sdesignation) = ?", strings.ToLower(c.Designation), strings.ToLower(c.Sdesignation)).First(&c)
	if result1.RowsAffected != 0 {
		output.Message = "Designation Or Short_Designation is alread exist"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusConflict
		return output
	}
	// result2 := db.Raw("Select * from public.desig where sdesignation =?", c.Sdesignation).First(&c)
	// if result2.RowsAffected !=0{
	// 	output.Message ="Designation is alread exist"
	// 	output.IsSuccess = false
	// 	output.Payload=nil
	// 	output.StatusCode = http.StatusBadRequest
	// 	return output
	// }
	result3 := db.Create(&c)
	if result3.RowsAffected == 0 {
		output.Message = "Designation creation failed"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusInternalServerError
		return output
	}

	type abc struct {
		Output model.Desig `json:"output"`
	}
	var a abc
	a.Output = c
	output.Message = "Designation create succesfully"
	output.IsSuccess = true
	output.Payload = a
	output.StatusCode = http.StatusOK
	return output
}

func (designationrepo *DesignationRepository) UpdateDesignation(input model.Desig) model.ResponseDto {
	var response model.ResponseDto
	if input.Code <= 0 {
		response.Message = " Code can't be null"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusBadRequest
		return response
	}
	if input.Designation == "" {
		response.Message = "Designation can't be null"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusBadRequest
		return response

	}
	if input.Sdesignation == "" {
		response.Message = "ShortDesig can't be null"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusBadRequest
		return response
	}
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var output model.Desig
	result := db.Where(&model.Desig{Code: input.Code}).First(&output)
	if result.RowsAffected == 0 {
		response.Message = "this code doesnot exists"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusNotFound
		return response
	}
	output.Designation = input.Designation
	output.Sdesignation = input.Sdesignation
	result1 := db.Where(&model.Desig{Code: input.Code}).Updates(&output)
	if result1.RowsAffected == 0 {
		response.Message = "No Employee info found for given criteria"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusInternalServerError
		return response
	}
	response.Message = "Employee info updated successfully"
	response.IsSuccess = true
	response.Payload = output
	response.StatusCode = http.StatusOK

	return response
}

func (designationrepo *DesignationRepository) Delete(c model.Desig) model.ResponseDto {
	var output model.ResponseDto
	if c.Code <= 0 {
		output.Message = "Invalid code"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output
	}
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	result := db.Where("code = ?", c.Code).Delete(&c)
	if result.RowsAffected == 0 {
		output.Message = "No info found for given criteria"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusNotFound
		return output
	}
	output.Message = "Deleted successfully"
	output.IsSuccess = true
	output.Payload = nil
	output.StatusCode = http.StatusOK
	return output
}
