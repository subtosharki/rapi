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
	Root.AddCommand(newMiddlewareCmd)
}

var newMiddlewareCmd = &cobra.Command{
	Use:   "new:middleware",
	Short: "Create a new middleware",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("rapi.json")
		if err != nil {
			lib.Error("rapi.json not found, please run rapi init")
			lib.ExitBad()
		}
		lib.LoadConfig()
		middlewaresPath := viper.GetString("middlewarespath")
		if middlewaresPath == "" {
			lib.Error("middlewarespath not found in rapi.json")
		}
		framework := viper.GetString("framework")
		if framework == "" {
			lib.Error("framework not found in rapi.json")
		}
		mainFile := viper.GetString("mainfile")
		if mainFile == "" {
			lib.Error("mainfile not found in rapi.json")
		}
		middlewareName := args[0]
		_, err = os.Stat(middlewaresPath)
		if err != nil {
			lib.Error("middlewarespath not found")
			lib.ExitBad()
		}
		if strings.Contains(middlewareName, "/") {
			lib.Error("Middleware name cannot contain /")
			lib.ExitBad()
		}
		_, err = os.Stat(middlewaresPath + "/" + middlewareName + ".go")
		if err == nil {
			lib.Error("Middleware already exists")
			lib.ExitBad()
		}
		var middlewareType string
		for middlewareType != "1" && middlewareType != "2" && middlewareType != "3" {
			lib.Info("Select middleware type:\n")
			println("1. Global")
			println("2. Group")
			println("3. Route")
			_, err := fmt.Scanln(&middlewareType)
			lib.ErrorCheck(err)
		}
		var groupOrRouteName string
		if middlewareType == "2" {
			lib.Info("Enter group name:")
			for groupOrRouteName == "" {
				_, err := fmt.Scanln(&groupOrRouteName)
				lib.ErrorCheck(err)
			}
		} else if middlewareType == "3" {
			lib.Info("Enter route name:")
			for groupOrRouteName == "" {
				_, err := fmt.Scanln(&groupOrRouteName)
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
			if middlewareType == "1" {
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
				newLine := splitFile[line] + "\n" + wordsOfLine[0] + ".Use(middlewares." + lib.UpFirstLetter(middlewareName) + ")"
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
					if strings.Contains(v, "middlewares") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + middlewaresPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			} else if middlewareType == "2" {
				var appName string
				for _, v := range splitFile {
					if strings.Contains(v, "fiber.New") {
						appName = strings.Split(v, " ")[0]
						break
					}
				}
				var line int
				for i, v := range splitFile {
					if strings.Contains(v, appName+".Group("+groupOrRouteName+")") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("No group found")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitFile[line], " ")
				newLine := splitFile[line] + "\n" + wordsOfLine[0] + ".Use(middlewares." + lib.UpFirstLetter(middlewareName) + ")"
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
					if strings.Contains(v, "middlewares") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + middlewaresPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			} else if middlewareType == "3" {
				var appName string
				for _, v := range splitFile {
					if strings.Contains(v, "fiber.New") {
						appName = strings.Split(v, " ")[0]
						break
					}
				}
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
				newLine := splitFile[line] + "\n" + appName + ".Use(\"" + groupOrRouteName + "\", middlewares." + lib.UpFirstLetter(middlewareName) + ")"
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
					if strings.Contains(v, "middlewares") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + middlewaresPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		case "gin":
			if middlewareType == "1" {
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
				newLine := splitFile[line] + "\n" + wordsOfLine[0] + ".Use(middlewares." + lib.UpFirstLetter(middlewareName) + ")"
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
					if strings.Contains(v, "middlewares") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + middlewaresPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			} else if middlewareType == "2" {
				var appName string
				for _, v := range splitFile {
					if strings.Contains(v, "gin.Default") || strings.Contains(v, "gin.New") {
						appName = strings.Split(v, " ")[0]
						break
					}
				}
				var line int
				for i, v := range splitFile {
					if strings.Contains(v, appName+".Group("+groupOrRouteName+")") {
						line = i
						break
					}
				}
				if line == 0 {
					lib.Error("No group found")
					lib.ExitBad()
				}
				wordsOfLine := strings.Split(splitFile[line], " ")
				newLine := splitFile[line] + "\n" + wordsOfLine[0] + ".Use(middlewares." + lib.UpFirstLetter(middlewareName) + ")"
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
					if strings.Contains(v, "middlewares") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + middlewaresPath + "\"\n"

				}
				finalString := strings.Join(splitFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			} else if middlewareType == "3" {
				var appName string
				for _, v := range splitFile {
					if strings.Contains(v, "fiber.New") {
						appName = strings.Split(v, " ")[0]
						break
					}
				}
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
				newLine := splitFile[line] + "\n" + appName + ".Use(\"" + groupOrRouteName + "\", middlewares." + lib.UpFirstLetter(middlewareName) + ")"
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
					if strings.Contains(v, "middlewares") {
						found = true
						break
					}
				}
				if !found {
					goModFile := lib.LoadGoModuleFile()
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + middlewaresPath + "\"\n"

				}
				finalString := strings.Join(splitFile, "\n")
				_, err := file.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		default:
			lib.Error("Invalid framework")
		}
		file, err = os.Create(middlewaresPath + "/" + middlewareName + ".go")
		lib.ErrorCheck(err)
		pathName := strings.Split(middlewaresPath, "/")
		middlewareName = lib.UpFirstLetter(middlewareName)
		switch framework {
		case "fiber":
			_, err := file.WriteString(fiber.BasicRoute(middlewareName, pathName[len(pathName)-1]))
			lib.ErrorCheck(err)
		case "gin":
			_, err := file.WriteString(gin.BasicRoute(middlewareName, pathName[len(pathName)-1]))
			lib.ErrorCheck(err)
		default:
			lib.Error("Invalid framework")
			lib.ExitBad()
		}
		err = file.Close()
		lib.ErrorCheck(err)
		lib.Info("Middleware created successfully")
		lib.ExitOk()
	},
}
