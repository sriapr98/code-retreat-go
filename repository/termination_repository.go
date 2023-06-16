package repository

import (
	"code-retreat-go/models"
	"encoding/json"
	"os"
)

type TerminationData struct {
	Termination []struct {
		Id                 int    `json:"id"`
		EmployeeName       string `json:"employee_name"`
		LastWorkingDate    string `json:"last_working_date"`
		Reason             string `json:"reason"`
		InitiationDate     string `json:"initiation_date"`
		Status             string `json:"status"`
		ExitCheckListItems []struct {
			ExitChecklistItemId int    `json:"exit_checklist_item_id"`
			IsCompleted         bool   `json:"is_completed"`
			CompletionDate      string `json:"completion_date"`
		} `json:"exit_check_list_items"`
	} `json:"termination"`
}

type TerminationRepository interface {
	LoadTerminationSeedData(repository EmployeeRepository) error
	GetAllTerminations() []models.Termination
}

type terminationRepository struct {
	terminationData []models.Termination
}

func (t *terminationRepository) GetAllTerminations() []models.Termination {
	return t.terminationData
}

func (t *terminationRepository) LoadTerminationSeedData(repository EmployeeRepository) error {
	fileBytes, err := os.ReadFile("repository/seed_data/termination.json")
	if err != nil {
		return err
	}
	var terminationData TerminationData
	err = json.Unmarshal(fileBytes, &terminationData)
	if err != nil {
		return err
	}
	for _, termination := range terminationData.Termination {
		var exitCheckListItems []models.EmployeeExitChecklistItem
		for _, exitCheckListItem := range termination.ExitCheckListItems {
			exitCheckListItems = append(exitCheckListItems, models.EmployeeExitChecklistItem{
				ExitChecklistItemID: exitCheckListItem.ExitChecklistItemId,
				IsCompleted:         exitCheckListItem.IsCompleted,
				CompletionDate:      exitCheckListItem.CompletionDate,
			})
		}
		t.terminationData = append(t.terminationData, models.Termination{
			Id:                termination.Id,
			Employee:          repository.GetEmployeeByName(termination.EmployeeName),
			LastWorkingDate:   termination.LastWorkingDate,
			Reason:            termination.Reason,
			InitiationDate:    termination.InitiationDate,
			TerminationStatus: models.TerminationStatus(termination.Status),
			ExitChecklist:     exitCheckListItems,
		})
	}
	return nil
}

func NewTerminationRepository() TerminationRepository {
	return &terminationRepository{}
}
