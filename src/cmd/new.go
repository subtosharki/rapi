package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
	Use:   "new",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Open(args[0])
		if err == nil {
			lib.Error("Directory " + args[0] + " already exists")
			lib.ExitBad()
		}
		println("Choose a Framework to use:\n")
		println("1. Fiber")
		println("2. Gin")
		var framework string
		println("Enter a number: ")
		for framework != "1" && framework != "2" {
			_, err := fmt.Scanln(&framework)
			lib.ErrorCheck(err)
		}
		lib.Info("Creating " + args[0] + " directory...")
		err = os.Mkdir(args[0], 0755)
		lib.ErrorCheck(err)
		err = os.Chdir(args[0])
		lib.ErrorCheck(err)
		err = os.Mkdir("src", 0755)
		lib.ErrorCheck(err)
		err = os.Mkdir("src/routes", 0755)
		lib.ErrorCheck(err)
		err = os.Mkdir("src/middlewares", 0755)
		lib.ErrorCheck(err)

		lib.Info("Initializing go mod...")
		err = exec.Command("go", "mod", "init", args[0]).Run()
		lib.ErrorCheck(err)

		lib.Info("Creating main.go...")
		mainFile, err := os.Create("src/main.go")
		lib.ErrorCheck(err)
		switch framework {
		case "1":
			_, err = mainFile.WriteString(fiber.BasicProject(args[0]))
			lib.ErrorCheck(err)
		case "2":
			_, err = mainFile.WriteString(gin.BasicProject(args[0]))
			lib.ErrorCheck(err)
		}

		lib.Info("Creating routes...")
		routesFile, err := os.Create("src/routes/basic_route.go")
		lib.ErrorCheck(err)
		switch framework {
		case "1":
			_, err = routesFile.WriteString(fiber.BasicRoute("BasicRoute", "routes"))
			lib.ErrorCheck(err)
		case "2":
			_, err = routesFile.WriteString(gin.BasicRoute("BasicRoute", "routes"))
			lib.ErrorCheck(err)
		}
		lib.Info("Creating middlewares...")
		middlewareFile, err := os.Create("src/middlewares/basic_middleware.go")
		lib.ErrorCheck(err)
		switch framework {
		case "1":
			_, err = middlewareFile.WriteString(fiber.BasicMiddleware("BasicMiddleware", "middlewares"))
			lib.ErrorCheck(err)
		case "2":
			_, err = middlewareFile.WriteString(gin.BasicMiddleware("BasicMiddleware", "middlewares"))
			lib.ErrorCheck(err)
		}

		switch framework {
		case "1":
			lib.Info("Installing Fiber...")
			err = exec.Command("go", "get", "-u", "github.com/gofiber/fiber/v2").Run()
			lib.ErrorCheck(err)
		case "2":
			lib.Info("Installing Gin...")
			err = exec.Command("go", "get", "-u", "github.com/gin-gonic/gin").Run()
			lib.ErrorCheck(err)
		}
		lib.Info("Creating rapi.json file")
		var frameworkName string
		switch framework {
		case "1":
			frameworkName = "fiber"
		case "2":
			frameworkName = "gin"
		}
		lib.SetupConfig(args[0], frameworkName, "src/routes", "src/middlewares", "src/main.go")
		lib.Info("Done! Run `cd " + args[0] + "` and `go run src/main.go` to start your server.")
		lib.ExitOk()
	},
}
