package validate

import "fmt"

type IntValidator struct {
	ValidatorFunc func(src int) bool
	FailMessage   string
}

func (v *IntValidator) Validate(value int) bool {
	valid := v.ValidatorFunc(value)
	if !valid && len(v.FailMessage) != 0 {
		fmt.Print(v.FailMessage)
	}
	return valid
}
