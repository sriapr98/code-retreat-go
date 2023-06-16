package models

type ExitCheckListItem struct {
	Id          int
	Description string
	Category    ExitChecklistItemCategory
	RegionType  ExitChecklistItemType
	RegionId    int
}

var (
	DeskKeys            = ExitCheckListItem{1, "Desk keys", Admin, COUNTRY, 4}
	FileCabinetKeys     = ExitCheckListItem{2, "File cabinet keys", Admin, COUNTRY, 4}
	AccessCard          = ExitCheckListItem{3, "Access card", Admin, COUNTRY, 4}
	CellPhone           = ExitCheckListItem{4, "Cell phone", Admin, COUNTRY, 4}
	SimCard             = ExitCheckListItem{5, "SIM card", Admin, COUNTRY, 4}
	LibraryBooks        = ExitCheckListItem{6, "Library books", Admin, COUNTRY, 4}
	SezCard             = ExitCheckListItem{7, "SEZ card", Admin, OFFICE, 8}
	Laptop              = ExitCheckListItem{8, "Laptop", Techops, COUNTRY, 4}
	Monitors            = ExitCheckListItem{9, "Monitors", Techops, COUNTRY, 4}
	Charger             = ExitCheckListItem{10, "Charger", Techops, COUNTRY, 4}
	RsaFob              = ExitCheckListItem{11, "RSA FOB", Techops, COUNTRY, 4}
	OutstandingExpenses = ExitCheckListItem{12, "Outstanding expenses", Finance, COUNTRY, 4}
	SodexoReimbursement = ExitCheckListItem{13, "Sodexo reimbursement", Finance, COUNTRY, 4}
	ExitSurvey          = ExitCheckListItem{13, "Exit survey", PeopleTeam, COUNTRY, 4}
	LeaveEncashment     = ExitCheckListItem{14, "Leave encashment", PeopleTeam, COUNTRY, 4}
	Gratuity            = ExitCheckListItem{15, "Gratuity", PeopleTeam, COUNTRY, 4}
)
