package utils

import (
	"time"
)

// ParseRFC3339Nano Parse from RFC3339Nano format to 2006-01-02 15:04:05
func ParseRFC3339Nano(date string) (dateParsed string, err error) {
	dt, err := time.Parse(time.RFC3339Nano, date)
	if err != nil {
		return "", err
	}

	dateParsed = dt.Format("2006-01-02 15:04:05")
	return dateParsed, nil
}

// ParseISO Parse from 2006-01-02 to 2006-01-02 15:04:05
func ParseISO(date string) (dateParsed string, err error) {
	dt, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	dateParsed = dt.Format("2006-01-02 15:04:05")
	return dateParsed, nil
}

// ParseDatetoTime ParseDate parse date string to format 2006-01-02 15:04:05
func ParseDatetoTime(date string) time.Time {
	if date == "" {
		return time.Now()
	}

	parsedDate, _ := time.Parse("2006-01-02 15:04:05", date)
	return parsedDate
}
