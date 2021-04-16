package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf(os.Stderr.Name())
		os.Exit(1)
	}

	baseurl := "http://tinyurl.com/api-create.php?url="

	lengthyUrl := os.Args[1]

	targetUrl := baseurl + lengthyUrl

	response, err := http.Get(targetUrl)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}

}
