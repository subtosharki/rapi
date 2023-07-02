package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/subtosharki/rapi/src/lib"
	"github.com/subtosharki/rapi/src/templates/fiber"
	"github.com/subtosharki/rapi/src/templates/gin"
	"os"
	"strings"
)

func init() {
	Root.AddCommand(newRouteCmd)
}

var newRouteCmd = &cobra.Command{
	Use:   "new:route",
	Short: "Create a new route",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("rapi.json")
		if err != nil {
			lib.Error("rapi.json not found, please run rapi init")
			lib.ExitBad()
		}
		lib.LoadConfig()
		routesPath := viper.GetString("routespath")
		if routesPath == "" {
			lib.Error("routespath not found in rapi.json")
		}
		framework := viper.GetString("framework")
		if framework == "" {
			lib.Error("framework not found in rapi.json")
		}
		mainFile := viper.GetString("mainfile")
		if mainFile == "" {
			lib.Error("mainfile not found in rapi.json")
		}
		routeName := args[0]
		_, err = os.Stat(routesPath)
		if err != nil {
			lib.Error("routespath not found")
			lib.ExitBad()
		}
		if strings.Contains(routeName, "/") {
			lib.Error("Route name cannot contain /")
			lib.ExitBad()
		}
		_, err = os.Stat(routesPath + "/" + routeName + ".go")
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
		file, err := os.OpenFile(mainFile, os.O_RDWR, 0644)
		fileBytes, err := os.ReadFile(mainFile)
		lib.ErrorCheck(err)
		defer func(file *os.File) {
			err := file.Close()
			lib.ErrorCheck(err)
		}(file)
		splitFile := strings.Split(string(fileBytes), "\n")
		switch framework {
		case "fiber":
			if routeType == "1" {
				var line int
				for i, v := range splitFile {
					if strings.Contains(v, "fiber.New") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("Could not find fiber.New")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitFile[line], " ")
				newLine := splitFile[line] + "\n" + wordsOfLine[0] + ".Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitFile[line] = newLine
				var importStart int
				for i, v := range splitFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + routesPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			} else if routeType == "2" {
				var appName string
				for _, v := range splitFile {
					if strings.Contains(v, "fiber.New") {
						appName = strings.Split(v, " ")[0]
						break
					}
				}
				var line int
				for i, v := range splitFile {
					if strings.Contains(v, appName+".Group("+groupName+")") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("No group found")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitFile[line], " ")
				newLine := splitFile[line] + "\n" + wordsOfLine[0] + ".Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitFile[line] = newLine
				var importStart int
				for i, v := range splitFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + routesPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		case "gin":
			if routeType == "1" {
				var line int
				for i, v := range splitFile {
					if strings.Contains(v, "gin.Default") || strings.Contains(v, "gin.New") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("Could not find gin.Default or gin.New")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitFile[line], " ")
				newLine := splitFile[line] + "\n" + wordsOfLine[0] + ".Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitFile[line] = newLine
				var importStart int
				for i, v := range splitFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + routesPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			} else if routeType == "2" {
				var appName string
				for _, v := range splitFile {
					if strings.Contains(v, "gin.Default") || strings.Contains(v, "gin.New") {
						appName = strings.Split(v, " ")[0]
						break
					}
				}
				var line int
				for i, v := range splitFile {
					if strings.Contains(v, appName+".Group("+groupName+")") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("No group found")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitFile[line], " ")
				newLine := splitFile[line] + "\n" + wordsOfLine[0] + ".Use(routes." + lib.UpFirstLetter(routeName) + ")"
				splitFile[line] = newLine
				var importStart int
				for i, v := range splitFile {
					if strings.Contains(v, "(") {
						importStart = i
						break
					}
				}
				var importEnd int
				for i, v := range splitFile {
					if strings.Contains(v, ")") {
						importEnd = i
						break
					}
				}
				imports := splitFile[importStart:importEnd]
				var found bool
				for _, v := range imports {
					if strings.Contains(v, "routes") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + routesPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		default:
			lib.Error("Invalid framework")
		}
		file, err = os.Create(routesPath + "/" + routeName + ".go")
		if err != nil {
			lib.Error("Error creating route file")
			lib.ExitBad()
		}
		pathName := strings.Split(routesPath, "/")
		routeName = lib.UpFirstLetter(routeName)
		switch framework {
		case "fiber":
			_, err := file.WriteString(fiber.BasicRoute(routeName, pathName[len(pathName)-1]))
			if err != nil {
				lib.Error("Error writing to file")
				lib.ExitBad()
			}
		case "gin":
			_, err := file.WriteString(gin.BasicRoute(routeName, pathName[len(pathName)-1]))
			if err != nil {
				lib.Error("Error writing to file")
				lib.ExitBad()
			}
		default:
			lib.Error("Invalid framework")
			lib.ExitBad()
		}
		err = file.Close()
		if err != nil {
			lib.Error("Error closing file")
			lib.ExitBad()
		}
		lib.Info("Route created successfully")
		lib.ExitOk()
	},
}
