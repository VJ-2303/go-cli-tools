package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/VJ-2303/fuzzy"
)

func main() {
	var fileName string
	var searchTerm string

	flag.StringVar(&fileName, "file", "", "File name to search")
	flag.StringVar(&searchTerm, "search", "", "Term to Search")

	flag.Parse()

	if fileName == "" || searchTerm == "" {
		fmt.Println("filename and search term must be provided")
		os.Exit(1)
	}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	count, lines := fuzzy.FuzzyFind(file, searchTerm)
	fmt.Printf("%s appread %d times\n", searchTerm, count)
	for _, line := range lines {
		fmt.Printf("Appeared in line %d\n", line)
	}
}
