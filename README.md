# go-api

## Description
RESTful API template project for Go using Gin

## Feature
Cloud friendly

CMD layout

MVC Pattern

JWT support

TLS support

PostgreSQL support

Redis support without use case

FTP support without use case

reCaptcha support (can be remove by finding the keyword "WARN" or "captcha")

Existing use case for user management

## Futur update
Add external emailing support

Redis example use case

FTP example use case

## Requirement
Go: 1.14+

Makefile: any

Git: any

PostgreSQL

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

### Set configuration
```
$EXPORT ABC=DEF
```
Add any configs that you need

### Installation
```
$make install
```

### Vendoring
```
$make vendor
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
