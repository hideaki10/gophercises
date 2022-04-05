package main

import (
	"cyoa"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "json file")
	flag.Parse()
	fmt.Printf("%s\n", *filename)

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(file)
	var story cyoa.Story

	err = d.Decode(&story)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
