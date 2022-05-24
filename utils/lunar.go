package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nosixtools/solarlunar"
)

// Lunar 农历
type Lunar struct {
	Year  int
	Month int
	Day   int
	Leap  bool
}

// NewLunar 创建农历日期
func NewLunar(date string, leap bool) Lunar {
	segs := strings.Split(date, "-")
	year, month, day := 0, 0, 0
	if len(segs) >= 3 {
		dayI64, _ := strconv.ParseInt(segs[2], 10, 32)
		day = int(dayI64)
	}
	if len(segs) >= 2 {
		monthI64, _ := strconv.ParseInt(segs[1], 10, 32)
		month = int(monthI64)
	}
	if len(segs) >= 1 {
		yearI64, _ := strconv.ParseInt(segs[0], 10, 32)
		year = int(yearI64)
	}
	return Lunar{
		Year:  year,
		Month: month,
		Day:   day,
		Leap:  leap,
	}
}

func (l Lunar) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", l.Year, l.Month, l.Day)
}

// SolarToLunar 阳历转农历
func SolarToLunar(date time.Time) Lunar {
	lunarDate, leap := solarlunar.SolarToLuanr(date.Format("2006-01-02"))
	return NewLunar(lunarDate, leap)
}

// LunarToSolar 农历转阳历
func LunarToSolar(date Lunar) time.Time {
	solarDate := solarlunar.LunarToSolar(date.String(), date.Leap)
	solar, _ := time.Parse("2006-01-02", solarDate)
	return solar
}
