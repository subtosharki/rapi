package lib

import "strings"

func UpFirstLetter(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}
