package main

import "net/http"

func main() {
	addr := ":8888"
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.Handle("/www/", http.StripPrefix("/www/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(addr, nil)

}
