package lib

import "strings"

// MakeValidItemName returns the first word in the string titled and the rest toLower
func MakeValidItemName(name string) string {
	words := strings.Split(name, " ")

	if len(words) > 0 {
		words[0] = strings.Title(words[0])
		if len(words) > 1 {
			for i := 1; i < len(words); i++ {
				words[i] = strings.ToLower(words[i])
			}
		}
	}

	return strings.Join(words, " ")
}
