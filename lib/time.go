package lib

import "time"

var (
	jst, _ = time.LoadLocation("Asia/Tokyo")
)

func ParseTime(str string) time.Time {
	layout := "2006-01-02T15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}
	return t
}

func Now() time.Time {
	now := time.Now().In(jst)
	return now
}

func TimeToString(t time.Time) string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}
