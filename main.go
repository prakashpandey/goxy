/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @create date 2018-10-16 20:04:02
 * @modify date 2018-10-16 20:04:02
 * @desc [description]
 */
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

}

func handleHttp(w http.ResponseWriter, r *http.Request) {

}

func handleHttps(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// get configuration such as hostname, port, timeout etc
	config := getProxyConfig()
	fmt.Printf("Running goxy at [hostName: %s, port: %d, timeout: %d]\n", config.hostName, config.port, config.timeout)

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(config.hostName+":"+strconv.Itoa(config.port), nil)
}
