package domain

type Site struct {
	ID         string
	Name       string
	Domain     string
	TemplateID string
	Status     string
	Config     map[string]string
}
