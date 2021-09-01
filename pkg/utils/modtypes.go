package utils

var ModTypes []string = []string{
	"alpha",
	"beta",
	"release",
}

func IsValidModType(src string) bool {
	for _, v := range ModTypes {
		if src == v {
			return true
		}
	}
	return false
}
