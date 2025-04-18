package config

type Config struct {
	Domain     string
	Host       string
	Port       string
	UseGit     bool   // wether to use git (and thus the repository)
	Repository string // the reposityory to get the pages from
	DataPath   string // where to store the pages for caching (should be a docker volume)
}
