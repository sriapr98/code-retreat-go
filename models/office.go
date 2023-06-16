package models

type Office struct {
	Id      int
	Name    string
	Country Country
	Status  Status
}

var (
	London       = Office{1, "London", UnitedKingdom, Active}
	SanFrancisco = Office{2, "San Francisco", UnitedStates, Active}
	Dallas       = Office{3, "Dallas", UnitedStates, Active}
	Melbourne    = Office{4, "Melbourne", Australia, Active}
	Sydney       = Office{5, "Sydney", Australia, Active}
	Bangalore    = Office{6, "Bangalore", India, Active}
	Chennai      = Office{7, "Chennai", India, Active}
	Coimbatore   = Office{8, "Coimbatore", India, Active}
	Noida        = Office{9, "Noida", India, Closed}
	Hamburg      = Office{8, "Hamburg", Germany, Active}
	Beijing      = Office{9, "Beijing", China, Active}
	Wuhan        = Office{10, "Wuhan", China, Active}
	SaoPaulo     = Office{11, "Sao Paulo", Brazil, Active}
	Toronto      = Office{12, "Toronto", Canada, Closed}
)
