package main

import (
	"log"
	"net/http"

	"github.com/Mohammad-Katoozi/go_crash_course/18_testing/example4/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
