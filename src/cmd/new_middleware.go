package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
	Use:   "new:middleware [name]",
	Short: "Create a new middleware",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetConfig()
		middlewareName := args[0]
		if strings.Contains(middlewareName, "/") {
			lib.Error("Middleware name cannot contain /")
			lib.ExitBad()
		}
		_, err := os.Stat(config.MiddlewaresPath + "/" + middlewareName + ".go")
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
		mainFile, err := os.OpenFile(config.MainFilePath, os.O_RDWR, 0644)
		lib.ErrorCheck(err)
		defer func(file *os.File) {
			err := file.Close()
			lib.ErrorCheck(err)
		}(mainFile)
		fileBytes, err := os.ReadFile(config.MainFilePath)
		lib.ErrorCheck(err)
		splitFile := strings.Split(string(fileBytes), "\n")
		switch config.Framework {
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
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.MiddlewaresPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := mainFile.WriteString(finalString)
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
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.MiddlewaresPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := mainFile.WriteString(finalString)
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
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.MiddlewaresPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := mainFile.WriteString(finalString)
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
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.MiddlewaresPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := mainFile.WriteString(finalString)
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
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.MiddlewaresPath + "\"\n"
				}
				finalString := strings.Join(splitFile, "\n")
				_, err := mainFile.WriteString(finalString)
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
					splitFile[importStart+1] = splitFile[importStart+1] + "\n\"" + lib.GetGoModuleName(goModFile) + "/" + config.MiddlewaresPath + "\"\n"

				}
				finalString := strings.Join(splitFile, "\n")
				_, err := mainFile.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		default:
			lib.Error("Invalid framework")
		}
		mainFile, err = os.Create(config.MiddlewaresPath + "/" + middlewareName + ".go")
		lib.ErrorCheck(err)
		pathName := strings.Split(config.MiddlewaresPath, "/")
		middlewareName = lib.UpFirstLetter(middlewareName)
		switch config.Framework {
		case "fiber":
			_, err := mainFile.WriteString(fiber.BasicRoute(middlewareName, pathName[len(pathName)-1]))
			lib.ErrorCheck(err)
		case "gin":
			_, err := mainFile.WriteString(gin.BasicRoute(middlewareName, pathName[len(pathName)-1]))
			lib.ErrorCheck(err)
		default:
			lib.Error("Invalid framework")
			lib.ExitBad()
		}
		err = mainFile.Close()
		lib.ErrorCheck(err)
		lib.Info("New middleware " + middlewareName + " created successfully")
		lib.ExitOk()
	},
}
