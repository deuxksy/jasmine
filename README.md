# Templat Go

[![Go Report Card](https://goreportcard.com/badge/github.com/deuxksy/template-go-application)](https://goreportcard.com/report/github.com/deuxksy/template-go-application) Go Application 을 만들기 위한 기본 template

## Tree

```bash
.
├── README.md
├── assets
├── build
│   └── ci
│       └── build.jenkinsfiles
├── cmd
│   └── template
│       └── main.go
├── configs
│   ├── dev.yml
│   └── local.yml
├── go.mod
├── go.sum
├── internal
│   ├── configuration
│   │   └── config_model.go
│   └── logger
│       └── logger.go
├── logs
│   └── 220707
│       ├── error.log
│       └── out.log
├── pkg
└── test
```

## required module

- zap
- viper

## build

### windows

```bash
# GOOS=windows GOARCH=amd64 go build main.go
```

### linux

```bash
# GOOS=linux GOARCH=386 go build main.go
```

### mac

```bash
# GOOS=darwin GOARCH=arm64 go build main.go
```
