package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Hello, %s.</h1></body></html>", r.URL.Path[1:])
}

func main() {
	token, err := getToken()
	if err != nil {
		log.Println(err)
	}

	log.Printf("Auth token: %s\n", token)
	err = http.ListenAndServe(":8080", http.HandlerFunc(handler))
	if err != nil {
		log.Println(err)
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func getToken() (string, error) {

	req, err := http.NewRequest("GET", "http://auth-service.sales-system.svc.cluster.local:6000/auth/token/54bb2165-71e1-41a6-af3e-7da4a0e1e2c1", nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Basic "+basicAuth("admin@example.com", "gophers"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	var tokenStruct struct {
		Token string `json:"token"`
	}

	var tokenBytes []byte

	tokenBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(tokenBytes, &tokenStruct)
	if err != nil {
		return "", err
	}

	log.Println(resp)

	return tokenStruct.Token, nil
}
