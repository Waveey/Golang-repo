package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Goserver)
	http.ListenAndServe(":8000", nil)
}

func Goserver(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "name:Precious Okwu | credit-card:xxxx-xxxx-xxxx-0897 %s", r.URL.Path[1:])
}
