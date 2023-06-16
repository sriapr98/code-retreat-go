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
