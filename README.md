# go-api

## Desciption
RESTful API framework for Gin

## Feature
MVC Pattern

Custom CORS Parameter

TLS support

PostgreSQL support

Default token system

## Futur update
Add expire date support in token system

Add Basic Auth support

Add more unit test

Add external ftp connection support

Add external emailing support

Add redis support

## Requirement
Go: 1.12+

Makefile: any

Git: any

Verify GOPATH variable:
```
$go env | grep PATH
GOPATH="/home/{user}/go"
```

If the variable is empty, set it with:
```
$export GOPATH="{The path}"
```

## Project setup

### Set configuration file
```
$cp -r configuration-default configuration
```
Change any option that you need

### Installation
```
$make install
```
If you have some problem with golang.org/x/ dependency, please update your dependencies with:
```
go get -u golang.org/x/name_of_package
```

### Build
```
$make build
```

### Clean all and rebuild
```
$make re
```

### Run project
```
$make run
```

### Run tests
```
$make test
```

## Customize configuration
FLAGS in Makefile contains an option to boost compile process.
