package main

import (
	"cyoa"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "json file")
	flag.Parse()
	fmt.Printf("%s\n", *filename)

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(file)

	h := cyoa.NewHandler(story)

	fmt.Printf("Starting the server on port: %d\n", *port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), h)

}
