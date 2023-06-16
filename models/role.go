package models

type Role struct {
	Id   int
	Name string
}

var (
	Developer          = Role{1, "Developer"}
	QualityAnalyst     = Role{2, "Quality Analyst"}
	BusinessAnalyst    = Role{3, "Business Analyst"}
	ProductManager     = Role{4, "Product Manager"}
	ProjectManager     = Role{5, "Project Manager"}
	UiDeveloper        = Role{6, "UI Developer"}
	ExperienceDesigner = Role{7, "Experience Designer"}
	MobileDeveloper    = Role{8, "Mobile Developer"}
	DataEngineer       = Role{9, "Data Engineer"}
	DataScientist      = Role{10, "Data Scientist"}
	InfraConsultant    = Role{11, "Infra Consultant"}
)
