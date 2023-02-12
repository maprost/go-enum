package generator

import (
	"regexp"
	"strings"
)

var (
	specialNameRegExVal = `\_[a-zA-Z]+\_`
	specialNameRegEx    = regexp.MustCompile(specialNameRegExVal)
)

func commentSpecialNames(comment string, enum *Enum) {
	// find special names
	names := specialNameRegEx.FindAllString(comment, -1)

	if len(names) > 0 {
		if enum.SpecialNames == nil {
			enum.SpecialNames = make(map[string]string)
		}
		for _, n := range names {
			n = strings.TrimSpace(n)
			key := n
			key = strings.Trim(key, "_")
			if len(key) == 0 {
				continue
			}
			key = strings.ToUpper(key[0:1]) + key[1:]
			enum.SpecialNames[key] = n
		}
	}
}

var (
	groupNameRegExVal = `\^[a-zA-Z]+\^`
	groupNameRegEx    = regexp.MustCompile(groupNameRegExVal)
)

func commentGroupNames(comment string, enum *Enum) {
	// find special groups
	groups := groupNameRegEx.FindAllString(comment, -1)

	if len(groups) > 0 {
		if enum.GroupNames == nil {
			enum.GroupNames = make(map[string]string)
		}
		for _, n := range groups {
			n = strings.TrimSpace(n)
			key := n
			key = strings.Trim(key, "^")
			if len(key) == 0 {
				continue
			}
			key = strings.ToUpper(key[0:1]) + key[1:]
			enum.GroupNames[key] = n
		}
	}
}
