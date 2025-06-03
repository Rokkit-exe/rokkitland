package models

type Panel struct {
	Id       int    `yaml:"id"`
	Title    string `yaml:"title"`
	Format   string `yaml:"format"`
	X        int    `yaml:"x"`
	Y        int    `yaml:"y"`
	Width    int    `yaml:"width"`
	Height   int    `yaml:"height"`
	PaddingX int    `yaml:"padding-x"`
	PaddingY int    `yaml:"padding-y"`
}
