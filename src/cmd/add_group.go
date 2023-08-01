package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/subtosharki/rapi/src/lib"
	"os"
	"strings"
)

func init() {
	Root.AddCommand(newGroupCommand)
}

var newGroupCommand = &cobra.Command{
	Use:   "add:group [name]",
	Short: "Add a new group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.GetConfig()
		groupName := args[0]
		var nestedGroup string
		for nestedGroup != "y" && nestedGroup != "n" {
			lib.Info("Is this a nested group? (y/n)")
			_, err := fmt.Scanln(&nestedGroup)
			lib.ErrorCheck(err)
		}
		var parentGroup string
		if nestedGroup == "y" {
			lib.Info("Enter parent group name:")
			for parentGroup == "" {
				_, err := fmt.Scanln(&parentGroup)
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
		splitFile := strings.Split(string(mainFileBytes), "\n")
		switch config.Framework {
		case "fiber":
			if nestedGroup == "n" {
				var line int
				for i, v := range splitFile {
					if strings.Contains(v, "app := fiber.New()") {
						line = i
					}
				}
				if line == 0 {
					lib.Error("Could not find app := fiber.New() in main.go")
					lib.ExitBad()
				}
				splitFile[line] = splitFile[line] + "\n\t" + groupName + "Group := app.Group(\"/" + groupName + "\") \n{\n\n\t}"
				finalString := strings.Join(splitFile, "\n")
				_, err = mainFile.WriteString(finalString)
				lib.ErrorCheck(err)
			} else {
				var parentLine int
				for i, v := range splitFile {
					if strings.Contains(v, "app.Group(\"/"+parentGroup+"\")") {
						parentLine = i + 1
					}
				}
				if parentLine == 0 {
					lib.Error("Could not find app.Group(\"/" + parentGroup + "\") in main.go")
					lib.ExitBad()
				}
				if splitFile[parentLine+1] == "{" {
					parentLine += 1
				}
				splitFile[parentLine] = splitFile[parentLine] + "\n\t" + groupName + "Group := app.Group(\"/" + groupName + "\") \n{\n\n\t}"
				finalString := strings.Join(splitFile, "\n")
				_, err = mainFile.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		case "gin":
			if nestedGroup == "n" {
				var line int
				for i, v := range splitFile {
					if strings.Contains(v, "r := gin.Default()") {
						line = i
					}
				}
				if line == 0 {
					lib.Error("Could not find r := gin.Default() in main.go")
					lib.ExitBad()
				}
				splitFile[line] = splitFile[line] + "\n\t" + groupName + "Group := r.Group(\"/" + groupName + "\") \n{\n\n\t}"
				finalString := strings.Join(splitFile, "\n")
				_, err = mainFile.WriteString(finalString)
				lib.ErrorCheck(err)
			} else {
				var parentLine int
				for i, v := range splitFile {
					if strings.Contains(v, "r.Group(\"/"+parentGroup+"\")") {
						parentLine = i + 1
					}
				}
				if parentLine == 0 {
					lib.Error("Could not find r.Group(\"/" + parentGroup + "\") in main.go")
					lib.ExitBad()
				}
				if splitFile[parentLine+1] == "{" {
					parentLine += 1
				}
				splitFile[parentLine] = splitFile[parentLine] + "\n\t" + groupName + "Group := r.Group(\"/" + groupName + "\") \n{\n\n\t}"
				finalString := strings.Join(splitFile, "\n")
				_, err = mainFile.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		case "echo":
			if nestedGroup == "n" {
				var line int
				for i, v := range splitFile {
					if strings.Contains(v, "e := echo.New()") {
						line = i
					}
				}
				if line == 0 {
					lib.Error("Could not find e := echo.New() in main.go")
					lib.ExitBad()
				}
				splitFile[line] = splitFile[line] + "\n\t" + groupName + "Group := e.Group(\"/" + groupName + "\") \n{\n\n\t}"
				finalString := strings.Join(splitFile, "\n")
				_, err = mainFile.WriteString(finalString)
				lib.ErrorCheck(err)
			} else {
				var parentLine int
				for i, v := range splitFile {
					if strings.Contains(v, "e.Group(\"/"+parentGroup+"\")") {
						parentLine = i + 1
					}
				}
				if parentLine == 0 {
					lib.Error("Could not find e.Group(\"/" + parentGroup + "\") in main.go")
					lib.ExitBad()
				}
				if splitFile[parentLine+1] == "{" {
					parentLine += 1
				}
				splitFile[parentLine] = splitFile[parentLine] + "\n\t" + groupName + "Group := e.Group(\"/" + groupName + "\") \n{\n\n\t}"
				finalString := strings.Join(splitFile, "\n")
				_, err = mainFile.WriteString(finalString)
				lib.ErrorCheck(err)
			}
		default:
			lib.Error("Invalid framework")
		}
		lib.Info("New group created successfully")
		lib.ExitOk()
	},
}
