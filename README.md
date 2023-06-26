# go-api

## Description
RESTful API template project for Go using Gin

## Feature
CMD layout

Use golang.org/x/exp/slog as Logger

autoTLS support

Postgres support

Existing use case for user management

## Side packages
- JWT
- Query builder
- FTP
- Recaptcha

## Future update
Add external emailing support

More use cases

## Requirement
Go: 1.20+

Makefile: any

Git: any

PostgreSQL: 15+

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

### Update dependencies
```
$make update
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
