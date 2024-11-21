package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Hello, %s.</h1></body></html>", r.URL.Path[1:])
}

type User struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Roles        []string `json:"roles"`
	PasswordHash []byte   `json:"-"`
	Department   string   `json:"department"`
	Enabled      bool     `json:"enabled"`
	DateCreated  string   `json:"dateCreated"`
	DateUpdated  string   `json:"dateUpdated"`
}

type Items struct {
	Items []User `json:"items"`
}

func Users(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "http://sales-service.sales-system.svc.cluster.local:3000/users?page=1&rows=2", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	token, err := getToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not retrieve Auth token: %s", err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	var bytes []byte

	var items Items

	bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	err = json.Unmarshal(bytes, &items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	fmt.Fprint(w, "<html><body><table>")
	for _, user := range items.Items {
		fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td></tr>", user.Name, user.Email)
	}
	fmt.Fprint(w, "</table></body></html>")

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
