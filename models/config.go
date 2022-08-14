package models

type Config struct {
	Git struct {
		Repository string `yaml:"repository"`
	} `yaml:"git"`
	Files []string `yaml:"files"`
}
