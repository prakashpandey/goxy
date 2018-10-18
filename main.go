/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @create date 2018-10-16 20:04:02
 * @modify date 2018-10-16 20:04:02
 * @desc [description]
 */
package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

// handles http requests
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

// handles https requests
func handleTunneling(w http.ResponseWriter, r *http.Request) {
	log.Println("__INIT__ handleTunneling")
	destConn, err := net.DialTimeout("tcp", r.Host, 2*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
	log.Println("__DONE__ handleTunneling")
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	log.Println("__INIT__ transfer")
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
	log.Println("__DONE__ transfer")
}

func main() {
	// set configuration such as hostname, port, timeout etc
	setProxyConfig()
	log.Printf("Running goxy at [hostName: %s, port: %d]\n",
		config.hostName, config.port)

	server := &http.Server{
		Addr: config.hostName + ":" + strconv.Itoa(config.port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if auth(r) {
				if r.Method == http.MethodConnect {
					handleTunneling(w, r)
				} else {
					handleHTTP(w, r)
				}
			} else {
				log.Println("Unauthorized request received")
				http.Error(w, "Unauthorized request", http.StatusUnauthorized)
			}

		}),
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	if config.proto == "http" {
		log.Fatal(server.ListenAndServe())
	} else {
		log.Fatal(server.ListenAndServeTLS(config.proxyServCert, config.proxyServKey))
	}
}
