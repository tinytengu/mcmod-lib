package environment

import "mcmodlib/pkg/validate"

type Storage struct {
	Properties   PropertiesList `yaml:"properties"`
	Repositories RepositoryList
	Mods         []Mod `yaml:"mods"`
}

func (st *Storage) ValidateStringFlag(
	flags map[string]string,
	name string,
	validFunc func(string) bool,
	dest map[string]string,
	failMsg string) bool {
	// Init StringValidator to validate passed string data
	validator := validate.StringValidator{
		ValidatorFunc: validFunc,
		IgnoreEmpty:   true,
		FailMessage:   failMsg,
	}
	// Validate flags[name] value using validator
	valid := validator.Validate(flags[name])
	// If valid, write out to dest
	if valid {
		dest[name] = flags[name]
	}
	return valid
}
