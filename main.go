package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NYTimes/gziphandler"
)

func main() {
	gz, err := gziphandler.NewGzipLevelAndMinSize(gzip.DefaultCompression, 0)
	if err != nil {
		panic(err)
	}

	http.Handle("/", gz(http.HandlerFunc(handler)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received a request")
	log.Println(r.Header)
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)
}
