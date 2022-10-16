package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main()  {
	http.HandleFunc("/", httpFunc)
	//要求4
	http.HandleFunc("/healthz", httpHealthzFunc)
}

func httpHealthzFunc(w http.ResponseWriter, r *http.Request)  {
	successCode := "200"
	w.Write([]byte(successCode))
}

func httpFunc(w http.ResponseWriter, r *http.Request)  {
	//要求1
	if len(r.Header) > 0 {
		for k,v := range r.Header{
			w.Header().Set(k,v[0])
		}
	}

	//要求2
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION",version)

	//要求3
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}

	if net.ParseIP(ip) != nil {
		log.Println(ip)
	}
	log.Println(http.StatusOK)
	
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Success"))
}