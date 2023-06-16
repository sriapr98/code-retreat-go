package models

type Country struct {
	Id       int
	Name     string
	Iso3Code string
	Status   Status
}

var (
	UnitedKingdom = Country{1, "United Kingdom", "GBR", Active}
	UnitedStates  = Country{2, "United States", "USA", Active}
	Australia     = Country{3, "Australia", "AUS", Active}
	India         = Country{4, "India", "IND", Active}
	Germany       = Country{5, "Germany", "DEU", Active}
	China         = Country{6, "China", "CHN", Active}
	Brazil        = Country{7, "Brazil", "BRA", Active}
	Canada        = Country{8, "Canada", "CAN", Closed}
)

var countries = []Country{
	UnitedKingdom,
	UnitedStates,
	Australia,
	India,
	Germany,
	China,
	Brazil,
	Canada,
}

func GetCountryByName(name string) Country {
	for _, country := range countries {
		if country.Name == name {
			return country
		}
	}
	return Country{}
}
