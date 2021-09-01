package environment

import "mcmodlib/pkg/validate"

type PropertiesList map[string]string

func (pl *PropertiesList) Validate(value string, assignKey string, validateFunc func(string) bool, failMsg string) {
	// Init StringValidator to validate passed string data
	validator := validate.StringValidator{
		ValidatorFunc: validateFunc,
		IgnoreEmpty:   true,
		FailMessage:   failMsg,
	}
	// Validate value using validator
	valid := validator.Validate(value)
	// If valid, write out to PropertiesList
	if valid {
		(*pl)[assignKey] = value
	}
}

func (pl *PropertiesList) ValidateFlag(flags map[string]string, name string, validateFunc func(string) bool, failMsg string) {
	pl.Validate(flags[name], name, validateFunc, failMsg)
}
