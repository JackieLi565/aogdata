package aogdata

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const siteURL = "https://adventofcode.com"

func NewAOCData(year int, day int) (string, error) {
	token, err := getSessionToken()
	if err != nil {
		return "", err
	}

	url, err := getURL(year, day)
	if err != nil {
		return "", err
	}

	data, err := getRequestData(token, url)
	if err != nil {
		return "", err
	}

	return data, nil
}

func getSessionToken() (string, error) {
	err := godotenv.Load()

	if err != nil {
		return "", errors.New("error loading .env file")
	}

	return os.Getenv("AOC_SESSION"), nil
}

func getURL(year int, day int) (string, error) {
	currentYear := getCurrentYear(time.Now())
	isValidYear := year >= 2015 && year <= currentYear
	isValidDay := day >= 1 && day <= 25

	// checks to see if the newest problem has been released
	if year == currentYear {
		currentDay, err := getCurrentESTDay()
		if err != nil {
			return "", err
		}

		isValidCurrentDay := day >= 1 && day <= currentDay

		if isValidCurrentDay {
			isValidDay = isValidCurrentDay
		} else {
			return "", errors.New("problem hasn't been released yet")
		}
	}
	
	if isValidDay && isValidYear {
		return fmt.Sprintf("%s/%s/day/%s/input", siteURL, strconv.Itoa(year), strconv.Itoa(day)), nil
	}

	return "", errors.New("invalid date")
}

func getCurrentYear(currentTime time.Time) int {
	if currentTime.Month() == time.December && currentTime.Day() >= 1 {
		return currentTime.Year()
	} else {
		return currentTime.Year() - 1
	}
}

func getCurrentESTDay() (int, error) {
	et, err := time.LoadLocation("America/New_York")
	if err != nil {
		return 0, err
	}

	return time.Now().In(et).Day(), nil
}

func getRequestData(token string, url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Cookie", fmt.Sprintf("session=%s;", token))
	
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("error sending request")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		if resp.StatusCode == 400 {
			return "", errors.New("request error, session token invalid")
		} else {
			return "", errors.New("request error")
		}
	}

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", errors.New("gzip reading error")
	}

	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return "", errors.New("body reading error")
	}

	return string(body), nil
}
