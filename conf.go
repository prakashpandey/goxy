/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @create date 2018-10-16 22:49:32
 * @modify date 2018-10-16 22:49:32
 * @desc [description]
 */

package main

import (
	"os"
	"strconv"
)

type proxyConfig struct {
	hostName string
	port     int
	// in seconds
	timeout int
}

func getProxyConfig() *proxyConfig {
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		// default port
		port = 9090
	}
	timeout, err := strconv.Atoi(os.Getenv("timeout"))
	if err != nil {
		// default timeout in seconds
		timeout = 2
	}

	return &proxyConfig{
		hostName: os.Getenv("hostname"),
		port:     port,
		timeout:  timeout,
	}
}
