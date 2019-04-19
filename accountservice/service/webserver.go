package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	log.Println("Starting our HTTP server at: " + port)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Println("Error occured starting HTTP Listener at port: ", port)
		log.Println("Error: ", err.Error())
	}
}
