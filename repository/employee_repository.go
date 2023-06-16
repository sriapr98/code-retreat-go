package repository

import (
	"code-retreat-go/models"
	"encoding/json"
	"os"
)

type EmployeeRepository interface {
	LoadSeedData() error
	GetAllEmployees() []models.Employee
	GetEmployeeByName(name string) models.Employee
}

type EmployeesData struct {
	Employees []struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Role     string `json:"role"`
		Grade    string `json:"grade"`
		Office   string `json:"office"`
		HireDate string `json:"hireDate"`
	} `json:"employees"`
}

type employeeRepository struct {
	employeesData []models.Employee
}

func (e *employeeRepository) GetEmployeeByName(name string) models.Employee {
	for _, employee := range e.employeesData {
		if employee.Name == name {
			return employee
		}
	}
	return models.Employee{}
}

func (e *employeeRepository) GetAllEmployees() []models.Employee {
	return e.employeesData
}

func (e *employeeRepository) LoadSeedData() error {
	fileBytes, err := os.ReadFile("repository/seed_data/employees.json")
	if err != nil {
		return err
	}
	var employeesData EmployeesData
	err = json.Unmarshal(fileBytes, &employeesData)
	if err != nil {
		return err
	}
	for _, employee := range employeesData.Employees {
		e.employeesData = append(e.employeesData, models.Employee{
			Id:       employee.Id,
			Name:     employee.Name,
			Role:     models.GetRoleByName(employee.Role),
			Grade:    models.GetGradeByName(employee.Grade),
			Office:   models.GetOfficeByName(employee.Office),
			HireDate: employee.HireDate,
		})
	}
	return nil
}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}
