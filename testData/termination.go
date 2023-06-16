package testdata

import "code-retreat-go/models"

var resignation1 = models.Termination{
	Id:              1,
	Employee:        employee1,
	LastWorkingDate: "20-11-2023",
	Reason:          "NA",
	InitiationDate:  "15-06-2023",
	Termination:     models.SUBMITTED,
	ExitChecklist: []models.EmployeeExitChecklistItem{
		{
			ExitChecklistItemID: 3,
			IsCompleted:         false,
			CompletionDate:      "",
		},
		{
			ExitChecklistItemID: 8,
			IsCompleted:         false,
			CompletionDate:      "",
		},
	},
}

var resignation2 = models.Termination{
	Id:              2,
	Employee:        employee2,
	LastWorkingDate: "17-11-2023",
	Reason:          "NA",
	InitiationDate:  "17-04-2023",
	Termination:     models.ACCEPTED,
	ExitChecklist: []models.EmployeeExitChecklistItem{
		{
			ExitChecklistItemID: 3,
			IsCompleted:         true,
			CompletionDate:      "16-06-2023",
		},
		{
			ExitChecklistItemID: 8,
			IsCompleted:         true,
			CompletionDate:      "16-06-2023",
		},
	},
}
