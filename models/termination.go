package models

type Termination struct {
	Id                int
	Employee          Employee
	LastWorkingDate   string
	Reason            string
	InitiationDate    string
	TerminationStatus TerminationStatus
	ExitChecklist     []EmployeeExitChecklistItem
}
