# metamask-login-sample

a sample application for testing login with [MetaMask](https://github.com/MetaMask), built with [Echo](https://github.com/labstack/echo) + [Vue.js](https://github.com/vuejs/vue)

## Demo

https://mls.m0t0k1ch1.com

## Overview

![sequence diagram](src/img/sequence-diagram.png)

## Install

Please install [dep](https://github.com/golang/dep) and [yarn](https://github.com/yarnpkg/yarn) in advance.

``` sh
$ mkdir -p $GOPATH/src/github.com/m0t0k1ch1
$ cd $GOPATH/src/github.com/m0t0k1ch1
$ git clone git@github.com:m0t0k1ch1/metamask-login-sample.git
$ cd metamask-login-sample
$ dep ensure
$ yarn install
```

## Build

``` sh
$ go build
$ yarn build
```

## Run

```
$ ./metamask-login-sample --conf _config/sample.json
```
