package utils

import "regexp"

var ModFileRegexp = regexp.MustCompile(`^[\w._-]+\.jar$`)

func IsValidModFile(src string) bool {
	return ModFileRegexp.MatchString(src)
}
