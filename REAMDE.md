A toy chat bot
==============

## Get started

```sh
go get github.com/manhtai/cusbot
dep ensure
go run main.go
```

## SSL & Live reload support

```sh
go get github.com/codegangsta/gin
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
gin --certFile cert.pem --keyFile key.pem --all main.go
```
