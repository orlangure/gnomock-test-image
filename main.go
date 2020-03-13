package main

import (
	"net/http"
	"sync"
)

func main() {
	mux80 := http.NewServeMux()
	mux80.HandleFunc("/", echoHandler("80"))

	mux8080 := http.NewServeMux()
	mux8080.HandleFunc("/", echoHandler("8080"))

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		_ = http.ListenAndServe(":8080", mux8080)
	}()

	go func() {
		defer wg.Done()

		_ = http.ListenAndServe(":80", mux80)
	}()

	wg.Wait()
}

func echoHandler(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(msg))
	}
}
