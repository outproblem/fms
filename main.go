package main

import (
	"net/http"
	"fmt"
)

func main() {
	println("Net Demo server")

	http.HandleFunc("/", indexHandler)
	
	http.ListenAndServe("127.0.0.1:9010", nil)
}

	
func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("%#v\n", r)

	w.Write([]byte("Hello!"))
}
	
