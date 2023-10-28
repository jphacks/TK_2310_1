package lib

import "time"

func ParseTime(str string) time.Time {
	layout := "2006-01-02T15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}
	return t
}
