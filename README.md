# YAUM Yet Another URL Minifier

YAUM is a Web service used to minify URLs.

# Requirements

- redis

# Installation

## Binary installation

## Building source code

Install `dep` to fetch `yaum` Go dependencies then build source code:

```
go get github.com/golang/dep/cmd/dep
dep ensure -v
go build
```

# Runtime configuration

Configuration is read from environment variables. Below is the list of available environment variables:

| Name | Default | Role |
|---|---|---|
| REDIS_ADDR | `localhost:6379` | Redis address | 
| REDIS_PASSWORD | empty | Redis address | 
| LOG_LEVEL | `DEBUG` | Log level. One of `DEBUG`, `INFO`, `WARNING`, `ERROR`|
|||

# Deployment

Deploying to aws with tf:

1. Terraform handles creation of AWS resources (ec2, dns record...)
2. Ansible manages the local configuration of the ec2 vm and deploy `yaum`
