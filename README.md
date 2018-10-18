# Goxy

A proxy server written on golang

## Clone program

- `git clone https://github.com/prakashpandey/goxy`

## Configuration

- `host`: Host of proxy server

- `port`: The port you want your proxy server to run at

- `authorize`: `true/false` based on if you want to use an authenticated proxy server or not

- `user`: Username of proxy server. Valid only if `authorize` is set `true`

- `password`: Password of proxy server. Valid only if `authorize` is set `true`

## Build & Run

- `cd github.com/prakashpandey/goxy`

- `go build`

- `./goxy -host="" -port=9090 -authorize=true -user="root" -password="pass"`

## Curls commands for testing 

- `curl  -U proxyUser:proxyPassword --proxy http://localhost:9090 http://google.com`

