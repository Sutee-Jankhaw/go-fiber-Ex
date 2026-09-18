package utils

import "time"

const DateFormat = "2006-01-02"

func ParseDate(date string) (time.Time, error) {
	return time.Parse(DateFormat, date)
}

func FormatDate(date time.Time) string {
	return date.Format(DateFormat)
}
