# goxy

A proxy server build using golang

# Clone program

- `git clone https://github.com/prakashpandey/goxy`

# Confuguration

- `host`: Host of proxy server

- `port`: The port you want your proxy server to run at

- `authorize`: `true/false` based on if you want to use an authenticated proxy server or not

- `user`: Username of proxy server. Valid only if `authorize` is set `true`

- `password`: Password of proxy server. Valid only if `authorize` is set `true`

# Build & Run

- `cd github.com/prakashpandey/goxy`

- `go build`

- `./goxy -host="" -port=9090 -authorize=false -user="root" -password="pass"`

# Generate self-signed certificate and private key(Optional for HTTPS based proxy server)

- Install `openssl` if not already install, then run the below commands to generate `certtificate` and `private   key` for your server.

```
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
```
    


