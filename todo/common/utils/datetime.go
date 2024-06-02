package utils

import (
	"fmt"
	"time"
)

func StringDatetoTime(dateString string) (time.Time, error) {
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
func StartEndDate(startDate string, endDate string) (time.Time, time.Time, error) {
	if startDate == "" || endDate == "" {
		return time.Time{}, time.Time{}, fmt.Errorf("Invalid input, start date and end date are required")
	}
	start, err := StringDatetoTime(startDate)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("Invalid start date, date format should be yyyy-mm-dd")
	}

	end, err := StringDatetoTime(endDate)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("Invalid end date, date format should be yyyy-mm-dd")
	}
	end = end.AddDate(0, 0, 1)
	if end.Before(start) {
		return time.Time{}, time.Time{}, fmt.Errorf("End date should be greater than start date")
	}
	if start.AddDate(0, 0, 31).Before(end) {
		return time.Time{}, time.Time{}, fmt.Errorf("Date range should not be greater than 30 days")
	}
	return start, end, nil
}
func NowAddDays(days int) time.Time {
	day := time.Now().AddDate(0, 0, days)
	return time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.Local)
}
