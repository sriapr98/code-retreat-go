package models

type Employee struct {
	Id       int
	Name     string
	Role     Role
	Grade    Grade
	Office   Office
	HireDate string
}
