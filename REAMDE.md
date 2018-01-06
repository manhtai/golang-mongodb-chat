A toy chat app
==============

## Overview

The app contains 2 parts:

- Persistent structs: User, Channel, Message

These models holds information about User, Channel & Message from channels
when user send messages, data is saved to Mongodb.

- Websocket structs: Client, Room

These are responsible for opening Websocket connection to receive message from
user & broadcast it to all client in a specific channel.

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
