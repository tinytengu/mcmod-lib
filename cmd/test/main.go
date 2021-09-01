package main

import (
	"fmt"
	"mcmodlib/pkg/utils"
	"mcmodlib/pkg/validate"
)

func main() {
	storage := map[string]string{
		"test": "testval",
	}

	input := map[string]string{
		"modtype": "release1",
		"mcver":   "",
	}

	mcVer := input["mcver"]
	if len(mcVer) != 0 && !utils.IsValidMcVersion(mcVer) {
		fmt.Printf("Invalid Minecraft version: %v\n", mcVer)
	} else if len(mcVer) != 0 {
		storage["mcver"] = mcVer
	}

	validator := validate.StringValidator{
		ValidatorFunc: utils.IsValidModType,
		IgnoreEmpty:   true,
		FailMessage:   fmt.Sprintf("Invalid mod type: %v\n", input["modtype"]),
	}
	if validator.Validate(input["modtype"]) {
		storage["modtype"] = input["modtype"]
	}

	validator = validate.StringValidator{
		ValidatorFunc: utils.IsValidModType,
		IgnoreEmpty:   true,
		FailMessage:   fmt.Sprintf("Invalid mod type: %v\n", input["modtype"]),
	}
	if validator.Validate(input["modtype"]) {
		storage["modtype"] = input["modtype"]
	}

	fmt.Println(storage)
}
