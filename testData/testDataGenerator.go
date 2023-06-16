package testdata

import "code-retreat-go/models"

func GetTerminations() []models.Termination {
	resignationList := []models.Termination{
		resignation1,
		resignation2,
	}

	return resignationList
}
