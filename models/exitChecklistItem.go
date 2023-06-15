package models

type ExitChecklistItem struct {
	Id          int
	Description string
	Category    ExitChecklistItemCategory
	RegionType  ExitChecklistItemType
	RegionId    int
}
