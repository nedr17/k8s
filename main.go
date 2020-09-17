package main

import (
	"fmt"
	"net/http"
	"os"
)

var dbHost, dbPort, dbUser, dbPassword, logLevel, podIp string

func init() {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	logLevel = os.Getenv("LOG_LEVEL")
	podIp = os.Getenv("POD_IP")
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
	fmt.Fprintf(w, "%s:%s \nuser: %s, password: %s\nlog level: %s", dbHost, dbPort, dbUser, dbPassword, logLevel)
	fmt.Fprintf(w, "Pod ip: %s", podIp)
}

func main() {

	http.HandleFunc("/", handle)
	http.ListenAndServe(":4000", nil)
}
