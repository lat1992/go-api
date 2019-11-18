# Enigm-MVC

## Desciption
RESTful API framework for Gin

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
$cp config-default.json config.json
```
Change any option you need

### Installation
```
$make install
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

### Debug build without optimisations
```
$make debug_build
```

### Debug rebuild
```
$make dre
```

### Run tests
```
$make test
```

## Customize configuration
FLAGS in Makefile contains an option to boost compile process.
