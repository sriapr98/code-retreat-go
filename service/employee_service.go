package service

import (
	"code-retreat-go/models"
	"code-retreat-go/repository"
)

type EmployeeService interface {
	GetAllEmployees() []models.Employee
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
}

func (e employeeService) GetAllEmployees() []models.Employee {
	return e.employeeRepository.GetAllEmployees()
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository) EmployeeService {
	return employeeService{employeeRepository}
}
