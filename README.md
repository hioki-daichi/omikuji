# omikuji-server

[![Build Status](https://travis-ci.com/hioki-daichi/omikuji-server.svg?branch=master)](https://travis-ci.com/hioki-daichi/omikuji-server)
[![codecov](https://codecov.io/gh/hioki-daichi/omikuji-server/branch/master/graph/badge.svg)](https://codecov.io/gh/hioki-daichi/omikuji-server)
[![Go Report Card](https://goreportcard.com/badge/hioki-daichi/omikuji-server)](https://goreportcard.com/report/github.com/hioki-daichi/omikuji-server)
[![GoDoc](https://godoc.org/github.com/hioki-daichi/omikuji-server?status.svg)](https://godoc.org/github.com/hioki-daichi/omikuji-server)

omikuji-server is a JSON API server that randomly returns fortune.

## How to try

The server side starts as follows.

```shell
$ make build
$ ./omikuji-server
```

The client side sends a request as follows.

```shell
$ curl -s localhost:8080 | jq .
{
  "name": "Gopher",
  "fortune": "吉"
}
```

You can change the name returned from the default "Gopher" by specifying the name parameter.

```shell
$ curl -s 'localhost:8080/?name=hioki-daichi' | jq .
{
  "name": "hioki-daichi",
  "fortune": "大凶"
}
```

The name can be up to 32 characters.

```shell
$ curl -s 'localhost:8080/?name=A%20name%20longer%20than%20thirty%20two%20characters' | jq .
{
  "errors": [
    "Name is too long (maximum is 32 characters)"
  ]
}
```
