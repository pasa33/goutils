package datetime_utils

import (
	"fmt"
	"math/rand/v2"
	"strconv"
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

// just a basedate for gen random dates for botting
type BaseDate struct {
	day   int
	month int
	year  int
}

func NewBaseDate(d, m, y int) BaseDate {
	if d < 0 || d > 31 {
		d = 1
	}
	if m < 0 || d > 12 {
		m = 1
	}
	if y < 1900 || d > 3000 {
		y = 1900
	}
	return BaseDate{
		day:   d,
		month: m,
		year:  y,
	}
}

func TimeToBaseDate(t time.Time) BaseDate {
	return BaseDate{
		day:   t.Day(),
		month: int(t.Month()),
		year:  t.Year(),
	}
}

func (b *BaseDate) Day() int {
	return b.day
}
func (b *BaseDate) DayWith0() string {
	return fmt.Sprintf("%02d", b.day)
}
func (b *BaseDate) DayWitout0() string {
	return strings.TrimLeft(strconv.Itoa(b.day), "0")
}

func (b *BaseDate) Month() int {
	return b.month
}
func (b *BaseDate) MonthWith0() string {
	return fmt.Sprintf("%02d", b.month)
}
func (b *BaseDate) MonthWitout0() string {
	return strings.TrimLeft(strconv.Itoa(b.month), "0")
}

func (b *BaseDate) Year() int {
	return b.year
}
func (b *BaseDate) Year2() string {
	y := strconv.Itoa(b.year)
	if len(y) < 2 {
		return y
	}
	return y[len(y)-2:]
}
func (b *BaseDate) Year4() string {
	return strconv.Itoa(b.year)
}

func (b *BaseDate) toTime() time.Time {
	return time.Date(b.year, time.Month(b.month), b.day, 0, 0, 0, 0, time.UTC)
}

func RandomDateBetween(a, b BaseDate) BaseDate {
	startUnix := a.toTime().Unix()
	endUnix := b.toTime().Unix()

	randomUnix := rand.Int64N(endUnix-startUnix) + startUnix

	return TimeToBaseDate(time.Unix(randomUnix, 0))
}
