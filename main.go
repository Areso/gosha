package main

import (
	"fmt"
	"log"       //required for log
	"net/http"  //required for http
)

func game_heartbeat(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "OK")
}

func main() {
	http.HandleFunc("/game_heartbeat", game_heartbeat)
	log.Println("Starting server on port 6399")
	log.Fatal(http.ListenAndServe(":6399", nil))
	//fmt.Println(account.login)
}
