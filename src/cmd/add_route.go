package cmd

import (
	"errors"
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
		var config lib.Config
		config.Get()
		routeName := args[0]
		if strings.Contains(routeName, "/") {
			lib.Error("Route name cannot contain '/'")
		}
		_, err := os.Stat(config.RoutesPath + "/" + routeName + ".go")
		if err == nil {
			lib.Error("Route already exists")
		}
		var routeMethod string
		lib.Info("Select route method:\n")
		println("1. GET")
		println("2. POST")
		println("3. PUT")
		println("4. PATCH")
		println("5. DELETE")
		for routeMethod != "1" && routeMethod != "2" && routeMethod != "3" && routeMethod != "4" && routeMethod != "5" {
			_, err := fmt.Scanln(&routeMethod)
			lib.ErrorCheck(err)
		}
		var routeType string
		lib.Info("Select route type:\n")
		println("1. Global")
		println("2. Group")
		for routeType != "1" && routeType != "2" {
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
		mainFile, err := os.OpenFile(config.MainFilePath, os.O_RDWR, 0644)
		lib.ErrorCheck(err)
		defer func(file *os.File) {
			err := file.Close()
			lib.ErrorCheck(err)
		}(mainFile)
		mainFileBytes, err := os.ReadFile(config.MainFilePath)
		lib.ErrorCheck(err)
		splitMainFile := strings.Split(string(mainFileBytes), "\n")
		pathName := strings.Split(config.RoutesPath, "/")
		switch config.Framework {
		case "fiber":
			err := addRouteToMainFile(splitMainFile, routeName, routeType, mainFile, config, groupName, "fiber.New", "", false, routeMethod)
			CreateFileIfNil(err, config, routeName, pathName, fiber.BasicRoute)
		case "gin":
			err := addRouteToMainFile(splitMainFile, routeName, routeType, mainFile, config, groupName, "gin.Default", "gin.New", true, routeMethod)
			CreateFileIfNil(err, config, routeName, pathName, gin.BasicRoute)
		case "echo":
			err := addRouteToMainFile(splitMainFile, routeName, routeType, mainFile, config, groupName, "echo.New", "", true, routeMethod)
			CreateFileIfNil(err, config, routeName, pathName, echo.BasicRoute)
		case "chi":
			err := addRouteToMainFile(splitMainFile, routeName, routeType, mainFile, config, groupName, "chi.NewRouter", "", false, routeMethod)
			CreateFileIfNil(err, config, routeName, pathName, chi.BasicRoute)
		}
		lib.Info("New route " + routeName + " created successfully")
		lib.ExitOk()
	},
}

func addRouteToMainFile(splitMainFile []string, routeName string, routeType string, file *os.File, config lib.Config, groupName string, newRouterFuncName string, secondNewRouterFuncName string, methodNameCaps bool, routeMethod string) error {
	var methodVarName string
	switch routeMethod {
	case "1":
		methodVarName = "Get"
	case "2":
		methodVarName = "Post"
	case "3":
		methodVarName = "Put"
	case "4":
		methodVarName = "Patch"
	case "5":
		methodVarName = "Delete"
	}
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
	if routeType == "1" {
		var line int
		for i, v := range splitMainFile {
			if strings.Contains(v, newRouterFuncName) {
				line = i
				break
			}
		}
		if line == 0 && secondNewRouterFuncName == "" {
			return errors.New("Could not find " + newRouterFuncName + " in main.go")
		} else if line == 0 && secondNewRouterFuncName != "" {
			for i, v := range splitMainFile {
				if strings.Contains(v, secondNewRouterFuncName) {
					line = i
					break
				}
			}
			if line == 0 {
				return errors.New("Could not find " + newRouterFuncName + " or " + secondNewRouterFuncName + " in main.go")
			}
		}
		appVarName := strings.Split(splitMainFile[line], " ")[0]
		if methodNameCaps {
			methodVarName = strings.ToUpper(methodVarName)
		}
		newLine := splitMainFile[line] + "\n" + appVarName + "." + methodVarName + "(\"" + routeName + "\" ,routes." + lib.UpFirstLetter(routeName) + ")"
		splitMainFile[line] = newLine
		finalString := strings.Join(splitMainFile, "\n")
		_, err := file.WriteString(finalString)
		if err != nil {
			return err
		}
	} else if routeType == "2" {
		var groupVarName string
		var line int
		for i, v := range splitMainFile {
			if strings.Contains(v, groupName) {
				line = i
				groupVarName = strings.Split(splitMainFile[i], " ")[0]
				break
			}
		}
		if line == 0 {
			return errors.New("no group found")
		}
		newLine := splitMainFile[line] + "\n" + groupVarName + "." + methodVarName + "(\"" + routeName + "\" ,routes." + lib.UpFirstLetter(routeName) + ")"
		splitMainFile[line] = newLine
		finalString := strings.Join(splitMainFile, "\n")
		_, err := file.WriteString(finalString)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFileIfNil(err error, config lib.Config, routeName string, pathName []string, templateFunc func(string, string) string) {
	if err == nil {
		newRouteFile, err := os.Create(config.RoutesPath + "/" + routeName + ".go")
		lib.ErrorCheck(err)
		defer func(newRouteFile *os.File) {
			err := newRouteFile.Close()
			lib.ErrorCheck(err)
		}(newRouteFile)
		_, err = newRouteFile.WriteString(templateFunc(lib.UpFirstLetter(routeName), pathName[len(pathName)-1]))
		lib.ErrorCheck(err)
	} else {
		lib.ErrorCheck(err)
	}
}
