package main

import (
	"os"

	"github.com/markmontymark/iedb_to_pubmed"
)

func main() {
	for _, path := range os.Args[1:] {
		iedb_to_pubmed.Process_file(path)
	}
}
