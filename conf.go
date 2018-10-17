/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @create date 2018-10-16 22:49:32
 * @modify date 2018-10-16 22:49:32
 * @desc [description]
 */

package main

import (
	"flag"
	"log"
)

type proxyConfig struct {
	hostName string
	port     int
	// request timeout in seconds (Not implemented)
	timeout int
	// if proxy server will use authorization for every request
	authorize bool
	userName  string
	password  string
}

// contains configuration details
var config proxyConfig

func setProxyConfig() {
	// Parse configuration
	var host = flag.String("host", "", "proxy hostname")
	var port = flag.Int("port", 9090, "proxy port")
	var timeout = flag.Int("timeout", 2, "proxy port")
	var authorize = flag.Bool("authorize", false, "use authorized proxy server")
	var userName = flag.String("user", "root", "proxy username")
	var password = flag.String("password", "root", "proxy password")
	// Parse all flags
	flag.Parse()
	log.Printf("Setting configuration [hostName: %s, port: %d, , authorize: %v, userName: %s, password: %s, timeout: %d]\n",
		*host, *port, *authorize, *userName, *password, *timeout)
	config = proxyConfig{
		hostName:  *host,
		port:      *port,
		timeout:   *timeout,
		authorize: *authorize,
		userName:  *userName,
		password:  *password,
	}
}
