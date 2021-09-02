package utils

import (
	"regexp"
	"sort"
	"strings"

	"github.com/tinytengu/go-cfapi"
)

type ModSelector struct {
	Id      string
	Type    string
	Version string
	File    string
}

func ParseSelector(sel string) ModSelector {
	result := ModSelector{}

	parts := strings.Split(sel, ":")
	result.Id = parts[0]

	for _, part := range parts[1:] {
		if McVerRegexpAll.MatchString(part) {
			if !IsValidMcVersion(part) {
				continue
			} else {
				result.Version = part
			}
		} else if IsValidModType(part) {
			result.Type = part
		} else if IsValidModFile(part) {
			result.File = part
		}
	}

	return result
}

func FilterModFiles(mod cfapi.ApiResponse, sel ModSelector) []cfapi.ApiFile {
	files := []cfapi.ApiFile{}

	for _, file := range mod.Files {
		if len(sel.Type) != 0 && sel.Type != file.Type {
			continue
		}
		if len(sel.Version) != 0 && sel.Version != file.Version {
			continue
		}
		expr := regexp.MustCompile(sel.File)
		if !expr.MatchString(file.Name) {
			continue
		}
		files = append(files, file)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].UploadedAt.After(files[j].UploadedAt)
	})

	return files
}
