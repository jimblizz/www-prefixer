package main

import (
	"fmt"
	"log"
	"net/http"
)

func simpleRedirect(w http.ResponseWriter, r *http.Request) {
	targetUrl := fmt.Sprintf("https://www.%s%s", r.Host, r.URL.RequestURI())
	http.Redirect(w, r, targetUrl, 301)
}

func main() {
	http.HandleFunc("/", simpleRedirect)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
