package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
)

func getIP(r *http.Request) string {

	ip := r.Header.Get("X-FORWARDED-FOR")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	return ip
}

func index(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(map[string]string{
		"ip": getIP(r),
	})
	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}

func main() {
	log.Println("Starting app...")
	log.Println("Getting port...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}
	port = ":" + port
	log.Println("Listening on port: " + port)

	http.HandleFunc("/", index)

	log.Println("Starting http server...")
	log.Fatal(http.ListenAndServe(port, nil))
}
