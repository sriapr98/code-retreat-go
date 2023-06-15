package models

type Office struct {
	Id      int
	Name    string
	Country Country
	Status  Status
}
