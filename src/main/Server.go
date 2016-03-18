package main

import (
	"net/http"
	"debugutils"
)

func main() {
	setupServer()
}

func setupServer() {
	fs := http.FileServer(http.Dir("html"))
	http.Handle("/", fs)
	http.Handle("/html/", http.StripPrefix("/html/", fs))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	
	debugutils.ShutdownRunningServer()
	debugutils.EnableShutdownUtility()
	
	http.ListenAndServe(":8080", nil)
}