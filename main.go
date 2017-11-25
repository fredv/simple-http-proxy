package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func backendHost() string {
	relay := os.Getenv("BACKEND_HOST")
	return relay
}

func officialDomain() string {
	relay := os.Getenv("OFFICIAL_DOMAIN")
	return relay
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", backendHost(), r.URL), nil)
	if err != nil {
		log.Printf("Error in Do Request: %q", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error %q", err)
	}
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error in reading proxied response %q", err)
	}
	if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		fmt.Fprintf(w, "%s", strings.Replace(string(s), backendHost(), officialDomain(), -1))
	} else {
		fmt.Fprintf(w, "%s", s)
	}
}

func main() {
	server := &http.Server{
		Addr: ":3000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handleHTTP(w, r)
		}),
	}
	log.Fatal(server.ListenAndServe())
}
