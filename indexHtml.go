package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed public/index.html
var indexHtml string

func init() {
	occurences := strings.Count(indexHtml, "6800")
	log.Println("Replacing 6800 with", externalPort, "in", occurences, "occurences")
	indexHtml = strings.Replace(indexHtml, "6800", externalPort, occurences)
}
