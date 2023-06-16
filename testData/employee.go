package testdata

import "code-retreat-go/models"

var employee1 = models.Employee{
	Id:       1,
	Name:     "John Doe",
	Role:     models.Developer,
	Grade:    models.Consultant,
	Office:   models.Chennai,
	HireDate: "10-10-2015",
}

var employee2 = models.Employee{
	Id:       2,
	Name:     "Jane Doe",
	Role:     models.Developer,
	Grade:    models.LeadConsultant,
	Office:   models.Bangalore,
	HireDate: "10-11-2011",
}
