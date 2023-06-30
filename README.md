# rapi
rapi, (rap-ee) is a CLI tool to make and manage Go APIs, supporting popular frameworks like Gin, Fiber, and many more!

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
  help           Help about any command
  init           Initialize rapi in a current project
  new            Create a new project
  new:middleware Create a new middleware
  new:route      Create a new route
```

Currently, rapi supports the following frameworks:
- [x] Gin
- [x] Fiber
- [ ] Echo
- [ ] Chi
- [ ] Gorilla Mux