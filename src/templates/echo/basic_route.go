package echo

func BasicRoute(packageName string) string {
	return `package ` + packageName + `
	middleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}`
}
