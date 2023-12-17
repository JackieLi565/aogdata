package main

import (
	"testing"
	"time"
)

func TestGetURL(t *testing.T) {
	testCases := []struct {
		year         int
		day          int
		expectedURL  string
	}{
		{2022, 5, "https://adventofcode.com/2022/day/5/input"},
		{2010, 10, ""},
		{2022, 0, "" },
		{2010, 0, "" },
	}

	for _, testCase := range testCases {
		url, _ := getURL(testCase.year, testCase.day)

		if url != testCase.expectedURL {
			t.Errorf("For year %d and day %d, expected URL %s, but got %s", testCase.year, testCase.day, testCase.expectedURL, url)
		}
	}
}

func TestGetCurrentYear(t *testing.T) {
	testCases := []struct {
		currentTime 	time.Time
		expectedYear  int
	}{
		{time.Date(2022, time.December, 5, 0, 0, 0, 0, time.UTC), 2022},
		{time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC), 2021},
	}
	
	for _, testCase := range testCases {
		year := getCurrentYear(testCase.currentTime)

		if year != testCase.expectedYear {
			t.Errorf("For year %d and month %s, expected year %d, but got %d", testCase.currentTime.Year(), testCase.currentTime.Month(), testCase.expectedYear, year)
		}
	}
}
