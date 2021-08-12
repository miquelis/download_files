package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cavaliercoder/grab"
)

func main() {
	url := flag.String("url", "", "URL to download")
	filename := flag.String("file", ".", "Inform the file name")

	flag.Parse()

	if *url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	resp, err := grab.Get(*filename, *url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Download saved to", resp.Filename)

}
