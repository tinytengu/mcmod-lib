package utils

import "regexp"

var McVerRegexp = regexp.MustCompile(`^(([1-9])\.([\d]{1,2})(?:\.([\d]{1,2}))?)$`)

func IsValidMcVersion(src string) bool {
	return McVerRegexp.MatchString(src)
}
