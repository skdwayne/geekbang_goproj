package main

import (
	"fmt"
	"github.com/Am2901/httpserver/src/metrics"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
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
	mux.HandleFunc("/images", images)

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
	clientip := getClientIP(r)
	log.Println("Success! Response code:", 200, "ClientIp:", clientip)
}

func images(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("work done！", randInt)))
}

//	需求3
func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		xForwardedFor := r.Header.Get("X-Forwarded-For")
		ip = strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	}
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	if ip == "" {
		res, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
		if err == nil {
			ip = res
		}
	}
	return ip
}

//	需求4
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK.")
}
