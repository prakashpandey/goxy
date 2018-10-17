# goxy

A proxy server build using golang

# Generate self-signed certificate and private key(for HTTPS)

- Install `openssl` if not already install, then run the below commands to generate `certtificate` and `private   key` for your server.

'''
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
'''


