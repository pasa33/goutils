package datetime

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

// 28 November 2018 09:01
//
// t = 09:01
//
// T = 09:01:00
//
// d = 28/11/2018
//
// D = 28 November 2018
//
// f = 28 November 2018 09:01
//
// F = Wednesday, 28 November 2018 09:01
//
// R = 3 years ago
func TimeToDiscordWh(t time.Time, timeStyle ...string) string {
	//https://gist.github.com/LeviSnoot/d9147767abeef2f770e9ddcd91eb85aa
	if t.IsZero() {
		return "-"
	}
	if len(timeStyle) == 0 {
		return fmt.Sprintf("<t:%d>", t.Unix())
	}
	s := ""
	for _, v := range timeStyle {
		s += fmt.Sprintf("<t:%d:%s>", t.Unix(), v)
	}
	return s
}
