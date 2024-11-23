package main

import (
	"github.com/andretefras/fullcycle-go-labs-1-cloudrun/internal/presentation"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", presentation.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
