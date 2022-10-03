package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

//go run main.go https://qna.habr.com/q/506775

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 || len(args) > 1 {
		return
	}
	url := args[0]
	httpClient := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}
