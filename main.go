package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"net/http"
	"github.com/gorilla/mux"
)
func EventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
	hostname, err := os.Hostname()
	if err != nil {
		w.Write([]byte("Can not get hostname!"))	
	} else {
		message := []string{
			"<html>",
			"<head><title>Hello World</title></head>",
			"<body style=\"text-align:center;\">",
			"<p><img src='./image/qingyuanos.jpg' /></p>",
			"<h3>Hello XiaoYuan!</h3><br />",
			"<p>Hey! My hostname is <b>" + hostname + "</b></p>",
			"</body>",
			"<html>",
		}
		messages := strings.Join(message[:], "")
		w.Write([]byte(messages))	
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", EventHandler)
	r.PathPrefix("/image/").Handler(http.StripPrefix("/image/", http.FileServer(http.Dir("./image/"))))
	
	fmt.Println("Server is listening on 80")
	
	log.Fatal(http.ListenAndServe(":80", r))
}
