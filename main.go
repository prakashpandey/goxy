/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @create date 2018-10-16 20:04:02
 * @modify date 2018-10-16 20:04:02
 * @desc [description]
 */
package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("__INIT__ rootHandler")
	if auth(r) {
		if r.Method == http.MethodConnect {
			handleHTTPS(w, r)
		} else {
			handleHTTP(w, r)
		}
	} else {
		http.Error(w, "Unauthorized request", http.StatusUnauthorized)
	}
	log.Println("__DONE__ rootHandler")
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("__INIT__ handleHTTP")
	req, err := http.NewRequest(r.Method, r.RequestURI, r.Body)
	if err != nil {
		log.Fatal(err)
	}
	// add headers to the request
	for k, v := range r.Header {
		req.Header.Set(k, v[0])
	}
	// server the target request
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// write the response back to the http connection
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
	log.Println("__DONE__ handleHTTP")
}

func handleHTTPS(w http.ResponseWriter, r *http.Request) {
	log.Println("__INIT__ handleHTTPS")
	log.Println("__DONE__ handleHTTPS")
}

func main() {
	// set configuration such as hostname, port, timeout etc
	setProxyConfig()
	log.Printf("Running goxy at [hostName: %s, port: %d]\n",
		config.hostName, config.port)

	// add root handler
	http.HandleFunc("/", rootHandler)

	// start https server in new goroutine
	// go http.ListenAndServeTLS(":9091", "cert.pem", "key.pem", nil)

	// start http server
	http.ListenAndServe(config.hostName+":"+strconv.Itoa(config.port), nil)
}
