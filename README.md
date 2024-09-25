# Kemendagri SIPD Service Boilerplate GO
Kemendagri SIPD Service Boilerplate GO

## Prerequisites
Prequisites package:
* [Docker](https://www.docker.com/get-started) - for developing, shipping, and running applications (Application Containerization).
* [Go](https://golang.org/) - Go Programming Language.
* [Make](https://golang.org/) - Automated Execution using Makefile.
* [swag](https://github.com/swaggo/swag) Converts Go annotations to Swagger Documentation 2.0. We've created a variety of plugins for popular Go web frameworks.
* [golang-migrate/migrate](https://github.com/golang-migrate/migrate#cli-usage) Database migrations written in Go. Use as CLI or import as library for apply migrations.

Optional package:
* [gocritic](https://github.com/go-critic/go-critic) Highly extensible Go source code linter providing checks currently missing from other linters.
* [gosec](https://github.com/securego/gosec) Golang Security Checker. Inspects source code for security problems by scanning the Go AST.
* [golangci-lint](https://github.com/golangci/golangci-lint) Go linters runner. It runs linters in parallel, uses caching, supports yaml config, has integrations with all major IDE and has dozens of linters included.

## ⚡️ Quick start
These instructions will get you a copy of the project up and running on docker container and on your local machine.
1. Install Prequisites and optional package to your system:
2. Rename `Makefile.example` to `Makefile` then fill it with your make setting.
3. Instant run by this command
```shell
make go
```

## Testing Notes

### go-critic failed test
- commentedOutCode: //some comment
- example failed code
```
#utils/validator.go 37
if regexp.MustCompile(`^[a-zA-Z0-9/*-]*$`).MatchString(fieldstr) {
    return true
}else{
    return false
}
```

### gosec failed test
```
h := md5.New()
log.Println(h)
```

### lint failed test
```
#controllers/menu.go 143
if _, found := data[name]; found {
    data[name] = append(data[name], role)
} else {
    data[name] = []string{role}
}
```