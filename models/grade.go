package models

type Grade struct {
	Id   int
	Name string
}

var (
	ConsultantGraduate  = Grade{1, "Consultant Graduate"}
	Consultant          = Grade{2, "Consultant"}
	SeniorConsultant    = Grade{3, "Senior Consultant"}
	LeadConsultant      = Grade{4, "Lead Consultant"}
	PrincipalConsultant = Grade{5, "Principal Consultant"}
	Director            = Grade{6, "Director"}
)

var grades = []Grade{
	ConsultantGraduate,
	Consultant,
	SeniorConsultant,
	LeadConsultant,
	PrincipalConsultant,
	Director,
}

func GetGradeByName(name string) Grade {
	for _, grade := range grades {
		if grade.Name == name {
			return grade
		}
	}
	return Grade{}
}
