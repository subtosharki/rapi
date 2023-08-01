package chi

func BasicMiddleware(middlewareName string, middlewarePathName string) string {
	return `package ` + middlewarePathName + `

import "net/http"

func ` + middlewareName + `(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    next.ServeHTTP(w, r)
  })
}`
}
