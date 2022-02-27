package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var VERSION string

func main() {
	//	VERSION初始化
	if VERSION = os.Getenv("VERSION"); VERSION == "" {
		fmt.Println("SET VERSION")
		VERSION = "unknown"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("start http err:", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	//	需求1
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}
	//	需求2
	w.Header().Set("VERSION", VERSION)
}

//	需求4
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK.")
}
