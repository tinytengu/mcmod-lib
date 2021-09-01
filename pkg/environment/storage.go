package environment

type Storage struct {
	Properties   PropertiesList `yaml:"properties"`
	Repositories RepositoryList
	Mods         []Mod `yaml:"mods"`
}
