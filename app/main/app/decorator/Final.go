package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	cli := Decorator(http.DefaultClient,
		Authorization("sdef34DRG$%htTfDVve5wedDg4t$4%GRRtrGRFgffg"),
		Logging(log.New(os.Stdout, "client: ", log.LstdFlags)),
	)

	req, err := http.NewRequest("GET", "https://google.com?q=go", nil)
	if err != nil {
		log.Fatal("Failed to create request %v", err)
	}
	res, err := cli.Do(req)
	if err != nil {
		log.Fatal("Failed to Do request %v", err)
	}

	defer res.Body.Close

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Fatal(err)
	}
}
