package environment

type Storage struct {
	Properties   PropertiesList `yaml:"properties"`
	Repositories RepositoryList
	Mods         ModsList `yaml:"mods"`
}
