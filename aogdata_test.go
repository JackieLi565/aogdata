package aogdata

import (
	"fmt"
	"testing"
	"time"
)

func TestGetMostRecentYear(t *testing.T) {
	testCases := []struct {
		currentTime 	time.Time
		expectedYear  int
	}{
		{time.Date(2022, time.December, 5, 0, 0, 0, 0, time.UTC), 2022},
		{time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC), 2021},
	}
	
	for _, testCase := range testCases {
		year := getMostRecentYear(testCase.currentTime)

		if year != testCase.expectedYear {
			t.Errorf("For year %d and month %s, expected year %d, but got %d", testCase.currentTime.Year(), testCase.currentTime.Month(), testCase.expectedYear, year)
		}
	}
}

func TestValidateDate(t *testing.T) {
	testCases := []struct {
		year					int
		day						int
		expectedYear	string
		expectedDay		string
	}{
		{year: 2014, day: 2, expectedYear: "", expectedDay: ""},
		{year: 2022, day: 20, expectedYear: "2022", expectedDay: "20"},
	}

	for _, testCase := range testCases {
		year, day, _ := validateDate(testCase.year, testCase.day)

		if year != testCase.expectedYear && day != testCase.expectedDay {
			t.Errorf("For %d and day %d, expected year %s and day %s, but got %s and %s", testCase.year, testCase.day, testCase.expectedYear, testCase.expectedDay, year, day)
		}
	}
}

func TestGetURL(t *testing.T) {
	testCases := []struct {
		year					string
		day						string
		expectedURL		string
	}{
		{year: "2015", day: "4", expectedURL: fmt.Sprintf("%s/%s/day/%s/input", RootURL, "2015", "4") },
		{year: "2022", day: "20", expectedURL: fmt.Sprintf("%s/%s/day/%s/input", RootURL, "2022", "20") },
	}

	for _, testCase := range testCases {
		url := getURL(testCase.year, testCase.day)

		if testCase.expectedURL != url {
			t.Errorf("For year and day %s, %s, expected url %s, but got %s", testCase.year, testCase.day, testCase.expectedURL, url)
		}
	}
}
