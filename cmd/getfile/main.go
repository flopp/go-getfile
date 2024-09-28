package main

import (
	"fmt"
	"os"

	"github.com/flopp/go-getfile"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("USAGE: getfile URL TARGETFILE")
		os.Exit(1)
	}
	url := os.Args[1]
	targetFile := os.Args[2]

	client := getfile.NewClient()

	if err := client.GetIfNotExists(url, targetFile); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Downloaded %s to %s\n", url, targetFile)
}
