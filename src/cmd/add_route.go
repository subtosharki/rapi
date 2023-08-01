package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/subtosharki/rapi/src/lib"
	"github.com/subtosharki/rapi/src/templates/chi"
	"github.com/subtosharki/rapi/src/templates/echo"
	"github.com/subtosharki/rapi/src/templates/fiber"
	"github.com/subtosharki/rapi/src/templates/gin"
	"os"
	"strings"
)

func init() {
	Root.AddCommand(newRouteCmd)
}

var newRouteCmd = &cobra.Command{
	Use:   "add:route [name]",
	Short: "Add a new route",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetConfig()
		routeName := args[0]
		if strings.Contains(routeName, "/") {
			lib.Error("Route name cannot contain '/'")
			lib.ExitBad()
		}
		_, err := os.Stat(config.RoutesPath + "/" + routeName + ".go")
		if err == nil {
			lib.Error("Route already exists")
			lib.ExitBad()
		}
		var routeType string
		for routeType != "1" && routeType != "2" {
			lib.Info("Select route type:\n")
			println("1. Global")
			println("2. Group")
			_, err := fmt.Scanln(&routeType)
			lib.ErrorCheck(err)
		}
		var groupName string
		if routeType == "2" {
			lib.Info("Enter group name:")
			for groupName == "" {
				_, err := fmt.Scanln(&groupName)
				lib.ErrorCheck(err)
			}
		}
		file, err := os.OpenFile(config.MainFilePath, os.O_RDWR, 0644)
		lib.ErrorCheck(err)
		defer func(file *os.File) {
			err := file.Close()
			lib.ErrorCheck(err)
		}(file)
		mainFileBytes, err := os.ReadFile(config.MainFilePath)
		lib.ErrorCheck(err)
		splitMainFile := strings.Split(string(mainFileBytes), "\n")
		switch config.Framework {
		case "fiber":
			if routeType == "1" {
				var line int
				for i, v := range splitMainFile {
					if strings.Contains(v, "fiber.New") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("Could not find fiber.New")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitMainFile[line], " ")
				newLine := splitMainFile[line] + "\n" + wordsOfLine[0] + ".Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitMainFile[line] = newLine
				var importStart int
				for i, v := range splitMainFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitMainFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitMainFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitMainFile[importStart+1] = splitMainFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.RoutesPath + "\"\n"
				}
				finalString := strings.Join(splitMainFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			} else if routeType == "2" {
				var appName string
				for _, v := range splitMainFile {
					if strings.Contains(v, "fiber.New") {
						appName = strings.Split(v, " ")[0]
						break
					}
				}
				var line int
				for i, v := range splitMainFile {
					if strings.Contains(v, appName+".Group("+groupName+")") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("No group found")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitMainFile[line], " ")
				newLine := splitMainFile[line] + "\n" + wordsOfLine[0] + ".Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitMainFile[line] = newLine
				var importStart int
				for i, v := range splitMainFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitMainFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitMainFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitMainFile[importStart+1] = splitMainFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.RoutesPath + "\"\n"
				}
				finalString := strings.Join(splitMainFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		case "gin":
			if routeType == "1" {
				var line int
				for i, v := range splitMainFile {
					if strings.Contains(v, "gin.Default") || strings.Contains(v, "gin.New") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("Could not find gin.Default or gin.New")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitMainFile[line], " ")
				newLine := splitMainFile[line] + "\n" + wordsOfLine[0] + ".Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitMainFile[line] = newLine
				var importStart int
				for i, v := range splitMainFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitMainFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitMainFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitMainFile[importStart+1] = splitMainFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.RoutesPath + "\"\n"
				}
				finalString := strings.Join(splitMainFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			} else if routeType == "2" {
				var appName string
				for _, v := range splitMainFile {
					if strings.Contains(v, "gin.Default") || strings.Contains(v, "gin.New") {
						appName = strings.Split(v, " ")[0]
						break
					}
				}
				var line int
				for i, v := range splitMainFile {
					if strings.Contains(v, appName+".Group("+groupName+")") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("No group found")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitMainFile[line], " ")
				newLine := splitMainFile[line] + "\n" + wordsOfLine[0] + ".Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitMainFile[line] = newLine
				var importStart int
				for i, v := range splitMainFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitMainFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitMainFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitMainFile[importStart+1] = splitMainFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.RoutesPath + "\"\n"
				}
				finalString := strings.Join(splitMainFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		case "echo":
			if routeType == "1" {
				var line int
				for i, v := range splitMainFile {
					if strings.Contains(v, "e := echo.New()") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("Could not find e := echo.New()")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitMainFile[line], " ")
				newLine := splitMainFile[line] + "\n" + wordsOfLine[0] + ".Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitMainFile[line] = newLine
				var importStart int
				for i, v := range splitMainFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitMainFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitMainFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitMainFile[importStart+1] = splitMainFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.RoutesPath + "\"\n"
				}
				finalString := strings.Join(splitMainFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			} else if routeType == "2" {
				var line int
				for i, v := range splitMainFile {
					if strings.Contains(v, "e := echo.New()") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("Could not find e := echo.New()")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitMainFile[line], " ")
				newLine := splitMainFile[line] + "\n" + wordsOfLine[0] + ".Group(" + groupName + ").Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitMainFile[line] = newLine
				var importStart int
				for i, v := range splitMainFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitMainFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitMainFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitMainFile[importStart+1] = splitMainFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.RoutesPath + "\"\n"
				}
				finalString := strings.Join(splitMainFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		case "chi":
			if routeType == "1" {
				var line int
				for i, v := range splitMainFile {
					if strings.Contains(v, "r := chi.NewRouter()") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("Could not find r := chi.NewRouter()")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitMainFile[line], " ")
				newLine := splitMainFile[line] + "\n" + wordsOfLine[0] + ".Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitMainFile[line] = newLine
				var importStart int
				for i, v := range splitMainFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitMainFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitMainFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitMainFile[importStart+1] = splitMainFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.RoutesPath + "\"\n"
				}
				finalString := strings.Join(splitMainFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			} else if routeType == "2" {
				var line int
				for i, v := range splitMainFile {
					if strings.Contains(v, "r := chi.NewRouter()") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("Could not find r := chi.NewRouter()")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitMainFile[line], " ")
				newLine := splitMainFile[line] + "\n" + wordsOfLine[0] + ".Group(" + groupName + ").Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitMainFile[line] = newLine
				var importStart int
				for i, v := range splitMainFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitMainFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitMainFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitMainFile[importStart+1] = splitMainFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.RoutesPath + "\"\n"
				}
				finalString := strings.Join(splitMainFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		}
		newRouteFile, err := os.Create(config.RoutesPath + "/" + routeName + ".go")
		lib.ErrorCheck(err)
		defer func(newRouteFile *os.File) {
			err := newRouteFile.Close()
			lib.ErrorCheck(err)
		}(newRouteFile)
		pathName := strings.Split(config.RoutesPath, "/")
		routeName = lib.UpFirstLetter(routeName)
		switch config.Framework {
		case "fiber":
			_, err := newRouteFile.WriteString(fiber.BasicRoute(routeName, pathName[len(pathName)-1]))
			lib.ErrorCheck(err)
		case "gin":
			_, err := newRouteFile.WriteString(gin.BasicRoute(routeName, pathName[len(pathName)-1]))
			lib.ErrorCheck(err)
		case "echo":
			_, err := newRouteFile.WriteString(echo.BasicRoute(routeName, pathName[len(pathName)-1]))
			lib.ErrorCheck(err)
		case "chi":
			_, err := newRouteFile.WriteString(chi.BasicRoute(routeName, pathName[len(pathName)-1]))
			lib.ErrorCheck(err)
		default:
			lib.Error("Invalid framework")
			lib.ExitBad()
		}
		lib.Info("New route " + routeName + " created successfully")
		lib.ExitOk()
	},
}
