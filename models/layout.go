package models

type Layout struct {
	BasePanels map[string]Panel `yaml:"base_panels"`
	Pages      []Page           `yaml:"pages"`
}
