package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"ex3/story"
)

func main() {
	fileName := flag.String("file", "gopher.json", "source file that contains the story")
	port := flag.Int("port", 8080, "application port")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	enStory, err := story.DecodeJson(f)
	if err != nil {
		log.Fatal(err)
	}

	handler := story.NewHandler(enStory)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	fmt.Printf("Listening on port:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
