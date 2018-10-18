# Goxy

A proxy server written on golang

[![Build Status](https://travis-ci.org/prakashpandey/goxy.svg?branch=master)](https://travis-ci.org/prakashpandey/goxy)

## Clone program

- `git clone https://github.com/prakashpandey/goxy`

## Configuration

- `host`: Host of proxy server

- `port`: The port you want your proxy server to run at

- `authorize`: `true/false` based on if you want to use an authenticated proxy server or not

- `user`: Username of proxy server. Valid only if `authorize` is set `true`

- `password`: Password of proxy server. Valid only if `authorize` is set `true`

- `proto`: define protocol("http/https")

- `cert`: proxy server certificate file name. Valid only if `proto` is set to `https`

- `key`: proxy server private-key file name. Valid only if `proto` is set to `https`

## Build & Run

- `cd github.com/prakashpandey/goxy`

- `go build`

- Run proxy server in http mode(with authorization true): 

```
 ./goxy -host="localhost" -port=9090 -authorize=true -user="root" -password="pass" -proto="http"
```

- Run proxy server in https mode(with authorization true): 

```
 ./goxy -host="localhost" -port=9090 -authorize=true -user="root" -password="pass" -proto="https" -cert="cert.pem" -key="key.pem"
```

## Generate self-signed certificate(Optional for HTTPS based proxy server)
 
 - Install `openssl` if not already install, then run the below commands to generate `certtificate` and `private   key` for your server.

 ```
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
```

## Curl commands for testing 

- Curl command for get http url. Remove `proxyUser:proxyPassword` if `authorize` is set to `false`

```
 curl  -U proxyUser:proxyPassword --proxy http://localhost:9090 http://google.com
``` 



- Curl command for https request. Remove `proxyUser:proxyPassword` if `authorize` is set to `false`
 
```
 curl -Lv -U proxyUser:proxyPassword --proxy https://localhost:9090 --proxy-cacert cert.pem https://google.com
```

