package aogdata

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const RootUrl = "https://adventofcode.com"
const (
	Accept = "Accept"
	AcceptEncoding = "Accept-Encoding"
	Cookie = "Cookie"
) 

type config struct {
	year	uint
	day 	uint
}

func NewAdventConfig(year uint, day uint) *config {
	return &config{
		year: year,
		day: day,
	}
}

func GetPuzzleInput(config *config) (string) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	sessionToken := os.Getenv("AOC_SESSION")

	return fetchPuzzleInput(sessionToken, buildUrl(config))
}


func fetchPuzzleInput(token string, url string) (string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set(Accept, "text/plain")
	req.Header.Set(AcceptEncoding, "gzip")
	req.Header.Set(Cookie, fmt.Sprintf("session=%s;", token))
	
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)

		if res.StatusCode == http.StatusNotFound {
			panic("input problem does not exist")
		}
		if res.StatusCode != http.StatusOK {
			panic(fmt.Sprintf("unhandled status code - %d", res.StatusCode))
		}
	}

	reader, err := gzip.NewReader(res.Body)
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func buildUrl(config *config) string {
		return fmt.Sprintf("%s/%d/day/%d/input", RootUrl, config.year, config.day)
}
