package validate

import "fmt"

type StringValidator struct {
	ValidatorFunc func(src string) bool
	IgnoreEmpty   bool
	FailMessage   string
}

func (v *StringValidator) Validate(value string) bool {
	if len(value) == 0 && v.IgnoreEmpty {
		return true
	}

	valid := v.ValidatorFunc(value)
	if !valid && len(v.FailMessage) != 0 {
		fmt.Print(v.FailMessage)
	}
	return valid
}
