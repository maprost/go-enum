package generator

import (
	"regexp"
	"strings"
)

var (
	regExVal = `( |\t|^)\_[a-zA-Z]+\_( |$|\t)`
	regEx    = regexp.MustCompile(regExVal)
)

func commentSpecialNames(comment string, enum *Enum) {
	/*
		DSL:
		- _key_ is auto separated
	*/

	// find special names
	names := regEx.FindAllString(comment, -1)

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
