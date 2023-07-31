# rapi
rapi, (rap-ee) is a CLI tool to make and manage Go APIs, supporting popular frameworks like Gin, Fiber, and many more!

## Features
- Create a new project
- Initialize rapi in a current project
- Create a new route and middleware
- Add routes to your project, supporting Global and Grouping routes
- Add middlewares to your project, with options for Global, Group, and Route middleware
- Add groups to your project

## Installation
```
go get github.com/subtsoharki/rapi
```

## Usage
```
rapi [command] [arguments]
```

## Commands
```
  help             [command]    Help about any command
  init                          Initialize rapi in a current project
  new              [name]       Create a new project
  new:middleware   [name]       Create a new middleware
  new:route        [name]       Create a new route
```

Currently, rapi supports the following frameworks (more to come):
- [x] [Gin](https://github.com/gin-gonic/gin)
- [x] [Fiber](https://github.com/gofiber/fiber)
- [ ] [Echo](https://github.com/labstack/echo)
- [ ] [Chi](https://github.com/go-chi/chi)
- [ ] [Beego](https://github.com/beego/beego)
- [ ] [Buffalo](https://github.com/gobuffalo/buffalo)
- [ ] [Revel](https://github.com/revel/revel)
- [ ] [Iris](https://github.com/kataras/iris)
- [ ] [Martini](https://github.com/go-martini/martini)
- [ ] [Kit](https://github.com/go-kit/kit)
- [ ] [Go-zero](https://github.com/zeromicro/go-zero)
- [ ] [Kratos](https://github.com/go-kratos/kratos)
- [ ] [Fast HTTP](https://github.com/valyala/fasthttp)
- [ ] [Gocraft](https://github.com/gocraft/web)
- [ ] [Mocha](https://github.com/cloudretic/matcha/tree/main)

## Future Plans
- [ ] Add support for other languages frameworks (Express, Flask, etc.)Laravel
- [ ] Add `rapi run` command to run frameworks start command, with a port flag, and any other flags if needed