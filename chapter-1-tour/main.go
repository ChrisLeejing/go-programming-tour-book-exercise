package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Go", "Golang help info.")
	flag.StringVar(&name, "n", "Java", "Java help info.")
	flag.Parse()

	log.Printf("name: %s", name)
}
