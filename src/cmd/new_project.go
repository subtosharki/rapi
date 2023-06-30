package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/subtosharki/rapi/src/lib"
	"github.com/subtosharki/rapi/src/templates/fiber"
	"github.com/subtosharki/rapi/src/templates/gin"
	"os"
	"os/exec"
)

func init() {
	Root.AddCommand(createProjectCmd)
}

var createProjectCmd = &cobra.Command{
	Use:   "new:project",
	Short: "Create a new project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a project name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Open(args[0])
		if err == nil {
			lib.RapiError("Directory " + args[0] + " already exists")
			lib.RapiExitBad()
		}
		println("Choose a Framework to use:\n")
		println("1. Fiber")
		println("2. Gin")
		var framework string
		println("Enter a number: ")
		for framework != "1" && framework != "2" {
			_, err := fmt.Scanln(&framework)
			lib.RapiErrorCheck(err)
		}
		setup(args[0], framework)
	},
}

func setup(projectName string, framework string) {
	lib.RapiInfo("Creating " + projectName + " directory...")
	err := os.Mkdir(projectName, 0755)
	lib.RapiErrorCheck(err)
	err = os.Chdir(projectName)
	lib.RapiErrorCheck(err)
	err = os.Mkdir("src", 0755)
	lib.RapiErrorCheck(err)
	err = os.Mkdir("src/routes", 0755)
	lib.RapiErrorCheck(err)
	err = os.Mkdir("src/middlewares", 0755)
	lib.RapiErrorCheck(err)

	lib.RapiInfo("Initializing go mod...")
	err = exec.Command("go", "mod", "init", projectName).Run()
	lib.RapiErrorCheck(err)

	lib.RapiInfo("Creating main.go...")
	mainFile, err := os.Create("src/main.go")
	lib.RapiErrorCheck(err)
	switch framework {
	case "1":
		_, err = mainFile.WriteString(fiber.BasicProject(projectName))
		lib.RapiErrorCheck(err)
	case "2":
		_, err = mainFile.WriteString(gin.BasicProject(projectName))
		lib.RapiErrorCheck(err)
	}

	lib.RapiInfo("Creating routes...")
	routesFile, err := os.Create("src/routes/basic_route.go")
	lib.RapiErrorCheck(err)
	switch framework {
	case "1":
		_, err = routesFile.WriteString(fiber.BasicRoute)
		lib.RapiErrorCheck(err)
	case "2":
		_, err = routesFile.WriteString(gin.BasicRoute)
		lib.RapiErrorCheck(err)
	}
	lib.RapiInfo("Creating middleware...")
	middlewareFile, err := os.Create("src/middleware/basic_middleware.go")
	lib.RapiErrorCheck(err)
	switch framework {
	case "1":
		_, err = middlewareFile.WriteString(fiber.BasicMiddleware)
		lib.RapiErrorCheck(err)
	case "2":
		_, err = middlewareFile.WriteString(gin.BasicMiddleware)
		lib.RapiErrorCheck(err)
	}

	switch framework {
	case "1":
		lib.RapiInfo("Installing Fiber...")
		err = exec.Command("go", "get", "-u", "github.com/gofiber/fiber/v2").Run()
		lib.RapiErrorCheck(err)
	case "2":
		lib.RapiInfo("Installing Gin...")
		err = exec.Command("go", "get", "-u", "github.com/gin-gonic/gin").Run()
		lib.RapiErrorCheck(err)
	}

	lib.RapiInfo("Creating rapi.json file")
	viper.AddConfigPath(".")
	viper.SetConfigName("rapi")
	viper.SetConfigType("json")
	viper.Set("projectName", projectName)
	viper.Set("framework", "fiber")
	viper.Set("routesPath", "src/routes")
	viper.Set("middlewarePath", "src/middleware")
	err = viper.SafeWriteConfig()
	lib.RapiErrorCheck(err)

	lib.RapiInfo("Done! Run `cd " + projectName + "` and `go run src/main.go` to start your server.")
	lib.RapiExitOk()
}
