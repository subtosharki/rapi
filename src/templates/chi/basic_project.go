package chi

func BasicProject(packageName string) string {
	return `package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"` + packageName + `/src/middlewares"
  	"` + packageName + `/src/routes"
)


func main() {
	r := chi.NewRouter()
	r.Use(middlewares.BasicMiddleware)
	r.Get("/", routes.BasicRoute)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}`
}
