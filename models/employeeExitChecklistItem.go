package models

type EmployeeExitChecklistItem struct {
	TerminationId  int
	ExitChecklistItemID int
	IsCompleted    bool
	CompletionDate string
}
