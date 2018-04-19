# metamask-login-sample

a sample application for testing login with [MetaMask](https://github.com/MetaMask), built with [Echo](https://github.com/labstack/echo) + [Vue.js](https://github.com/vuejs/vue)

## Demo

https://mls.m0t0k1ch1.com

## Overview

![sequence diagram](src/img/sequence-diagram.png)

## Installation

Please install [dep](https://github.com/golang/dep) and [yarn](https://github.com/yarnpkg/yarn) in advance.

``` sh
$ git clone git@github.com:m0t0k1ch1/metamask-login-sample.git
$ git clone -b v1.8.4 git@github.com:ethereum/go-ethereum.git
$ cd metamask-login-sample
$ dep ensure
$ cp -r ../go-ethereum/crypto/secp256k1/libsecp256k1 vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/
$ yarn install
```

## Build & Run

``` sh
$ go build
$ yarn build
$ ./metamask-login-sample --conf _config/sample.json
```
