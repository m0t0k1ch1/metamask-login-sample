# metamask-login-sample

a sample application with login function using [MetaMask](https://github.com/MetaMask), built with [Echo](https://github.com/labstack/echo) + [Vue.js](https://github.com/vuejs/vue)

## Demo

https://mls.m0t0k1ch1.com

## Overview

![sequence diagram](src/img/sequence-diagram.png)

## Download

Please install [yarn](https://github.com/yarnpkg/yarn) in advance.

``` sh
$ mkdir -p $GOPATH/src/github.com/m0t0k1ch1
$ cd $GOPATH/src/github.com/m0t0k1ch1
$ git clone git@github.com:m0t0k1ch1/metamask-login-sample.git
$ cd metamask-login-sample
$ yarn install
```

## Build

__NOTICE: This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.__

``` sh
$ go build
$ yarn build
```

## Run

```
$ ./metamask-login-sample --conf _config/sample.json
```
