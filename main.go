package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	env1 := os.Getenv("GNOMOCK_TEST_1")
	env2 := os.Getenv("GNOMOCK_TEST_2")

	fmt.Println("received args:", os.Args[1:])
	fmt.Printf("starting with env1 = '%s', env2 = '%s'\n", env1, env2)

	mux80 := http.NewServeMux()
	mux80.HandleFunc("/", echoHandler("80"))
	mux80.HandleFunc("/env1", echoHandler(env1))

	mux8080 := http.NewServeMux()
	mux8080.HandleFunc("/", echoHandler("8080"))
	mux8080.HandleFunc("/env2", echoHandler(env2))

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
