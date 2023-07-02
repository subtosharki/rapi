package lib

import "strings"

func GetGoModuleName(goModFile []byte) string {
	return strings.Split(strings.Split(string(goModFile), " ")[1], "\n")[0]
}
