package repository

import (
	"net/http"
	"strings"

	"github.com/xyz/model"
	"github.com/xyz/util"
)

type HrdeptRepository struct {
}

func (hrdeptrepo *HrdeptRepository) GetAllEmployee() model.ResponseDto {
	var output model.ResponseDto
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	
	var ab []model.Hrdept
	result := db.Order("code").Find(&ab)
	if result.RowsAffected == 0 {
		output.Message = "No Department info found"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusNotFound
		return output
	}
	type tempOutPut struct {
		Output      []model.Hrdept `json:"output"`
		OutputCount int            `json:"outputCount"`
	}
	var tOutput tempOutPut
	tOutput.Output = ab
	tOutput.OutputCount = len(ab)
	output.Message = "List of departments"
	output.IsSuccess = true
	output.Payload = tOutput
	output.StatusCode = http.StatusOK

	return output
}

func (hrdeptrepo *HrdeptRepository) GetEmployeeById(c model.Hrdept) model.ResponseDto {
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
	result := db.Raw("select * from public.hrdept where code = ?", c.Code).First(&c)
	if result.RowsAffected == 0 {
		output.Message = "No Dept Employee info found"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusNotFound
		return output
	}
	type tempOutput struct {
		Output model.Hrdept `json:"output"`
	}
	var tOutput tempOutput
	tOutput.Output = c
	output.Message = "Dept Employee info details found for given criteria"
	output.IsSuccess = true
	output.Payload = tOutput
	output.StatusCode = http.StatusOK

	return output

}

func (hrdeptrepo *HrdeptRepository) AddEmployee(c model.Hrdept) model.ResponseDto {
	var output model.ResponseDto
	if c.Code <= 0 {
		output.Message = "Invalid code"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output

	}
	if c.Dept == "" {
		output.Message = "Dept can't be null"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output
	}
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	result := db.Raw("Select * from public.hrdept where code =?", c.Code).First(&c)
	if result.RowsAffected != 0 {
		output.Message = "Employee Code is already exist"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output
	}
	res:= db.Raw("select * from public.hrdept where lower (dept)=? ",strings.ToLower(c.Dept)).First(&c)
	if res.RowsAffected !=0{
		output.Message ="Deptartment name alread exist"
		output.IsSuccess = false
		output.Payload=nil
		output.StatusCode = http.StatusBadRequest
		return output
	}
	result1 := db.Create(&c)
	if result1.RowsAffected == 0 {
		output.Message = "Department creation failed"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusNotFound
		return output
	}

	type abc struct {
		Output model.Hrdept `json:"output"`
	}
	var a abc
	a.Output = c
	output.Message = "Employee create succesfully"
	output.IsSuccess = true
	output.Payload = a
	output.StatusCode = http.StatusCreated
	return output
}

func (hrdeptrepo *HrdeptRepository) Update(input model.Hrdept) model.ResponseDto {
	var response model.ResponseDto
	if input.Code <= 0 {
		response.Message = " Code can't be null"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusBadRequest
		return response
	}
	if input.Dept == "" {
		response.Message = "Dept can't be null"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusBadRequest
		return response
	}
	db := util.CreateConnection()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var output model.Hrdept
	result := db.Where(&model.Hrdept{Code: input.Code}).First(&output)
	if result.RowsAffected == 0 {
		response.Message = "this code doesnot exists"
		response.IsSuccess = false
		response.Payload = nil
		response.StatusCode = http.StatusNotFound
		return response
	}
	output.Dept = input.Dept
	result1 := db.Where(&model.Hrdept{Code: input.Code}).Updates(&output)
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

func (hrdeptrepo *HrdeptRepository) Delete(c model.Hrdept) model.ResponseDto {
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
