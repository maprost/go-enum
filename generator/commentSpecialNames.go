package generator

import (
	"regexp"
	"strings"
)

var (
	regExVal = `\_[a-zA-Z]+\_`
	regEx    = regexp.MustCompile(regExVal)
)

func commentSpecialNames(comment string, enum *Enum) {
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
