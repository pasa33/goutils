package datetime_utils

import (
	"fmt"
	"strings"
	"time"
)

func TimestampNow() int {
	return int(time.Now().UnixMilli())
}

func TimestampNowStr() string {
	return fmt.Sprint(time.Now().UnixMilli())
}

func TStoUTC(ts int) string {
	t := time.UnixMilli(int64(ts))
	return t.Format("2006-01-02")
}

func UTCtoTS(utc string) int {
	utc = strings.Split(utc, " ")[0]
	t, err := time.Parse("2006-01-02", utc)
	if err != nil {
		return 0
	}
	return int(t.UnixMilli())
}
