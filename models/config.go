package models

type Config struct {
	BasePanels     map[string]Panel `yaml:"base_panels"`
	Navigation     []string         `yaml:"navigation"`
	Packages       []Section        `yaml:"packages"`
	Configurations []Section        `yaml:"configurations"`
	Actions        []Action         `yaml:"actions"`
	Pages          []Page           `yaml:"pages"`
}
