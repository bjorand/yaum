# YAUM Yet Another URL Minifier

[![Build Status](https://travis-ci.org/bjorand/yaum.svg?branch=master)](https://travis-ci.org/bjorand/yaum)
[![Coverage Status](https://coveralls.io/repos/github/bjorand/yaum/badge.svg)](https://coveralls.io/github/bjorand/yaum)


YAUM is a Web service used to minify URLs.

# Requirements

- redis

# Installation

## Binary installation

You can download a pre-built release from Github.

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
| LISTEN_ADDR | `localhost:8080` | Yaum listen address |
| REDIS_ADDR | `localhost:6379` | Redis address | 
| REDIS_PASSWORD | empty | Redis address | 
| LOG_LEVEL | `DEBUG` | Log level. One of `DEBUG`, `INFO`, `WARNING`, `ERROR`|


# Deployment

1. `deploy/terraform`: Terraform handles creation of AWS resources (ec2 and networking)
2. `deploy/ansible`: Ansible manages the local configuration of the ec2 VM and deploy `yaum` from a release hosted on Github.
