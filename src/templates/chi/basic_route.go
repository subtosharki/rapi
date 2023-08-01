package chi

func BasicRoute(routeName string, packageName string) string {
	return `package ` + packageName + `

import "net/http"

func ` + routeName + `(w http.ResponseWriter, r *http.Request) {
  _, err := w.Write([]byte("Hello, World!"))
  if err != nil {
	panic(err)
  }
}`
}
