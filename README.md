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

Add test for http connections

Add external ftp connection support

Add external emailing support

Add redis support

## Requirement
Go: 1.12+

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

### Database
After install PostgreSQL, import database/database.sql file in your database.

### Set configuration file
```
$cp -r configuration-default configuration
```
Change any options that you need

### Installation
```
$make install
```
If you have some problem with golang.org/x/ dependency, please update your dependencies with:
```
$go get -u golang.org/x/name_of_package
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
