/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @create date 2018-10-18 03:34:02
 * @modify date 2018-10-18 03:34:02
 * @desc [description]
 */

package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func auth(r *http.Request) bool {
	// Check for authorization only if 'authorize' flag is enabled
	if config.authorize {
		authHeader := r.Header.Get("Proxy-Authorization")
		if len(strings.TrimSpace(authHeader)) > 0 {
			log.Printf("Proxy-Authorization=%s", authHeader)
			// get base64 encoded user credentials string which contains userName, pass
			headerStr := strings.Split(authHeader, "Basic ")
			if len(headerStr) == 2 {
				fmt.Println("headerStr: ", headerStr)
				encodedCredentials, err := base64.StdEncoding.DecodeString(headerStr[1])
				fmt.Println("encodedCredentials: ", encodedCredentials)
				if err == nil {
					decodedCredentials := string(encodedCredentials)
					fmt.Println("decodedCredentials: ", decodedCredentials)
					// userName, userPassword are seperated by ':'
					userCredentials := strings.Split(decodedCredentials, ":")
					if len(userCredentials) == 2 {
						userName := userCredentials[0]
						userPassword := userCredentials[1]
						if userName == config.userName && userPassword == config.password {
							return true
						}
					}
				}
			}
			// no need to  return false here as outer `return flase` will do the same.
		}
		return false
	}
	return true
}
