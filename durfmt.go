package durfmt

import (
	"bytes"
	"slices"
	"strconv"
	"time"
)

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
	Day                       = 24 * Hour
	Week                      = 7 * Day
	Month                     = Year / 12
	Year                      = 8766 * Hour // 365.25 days
)

var (
	layoutTokens = []string{"y", "M", "w", "d", "h", "m", "s", "ms", "us", "ns"}
	layoutDurMap = map[string]time.Duration{
		"ns": Nanosecond,
		"us": Microsecond,
		"ms": Millisecond,
		"s":  Second,
		"m":  Minute,
		"h":  Hour,
		"d":  Day,
		"w":  Week,
		"M":  Month,
		"y":  Year,
	}
)

func String(layout string, dur time.Duration) (string, error) {
	layoutBytes := []byte(layout)
	for _, t := range layoutTokens {
		i := bytes.Index(layoutBytes, []byte(t))
		if i == -1 {
			continue
		}

		switch t {
		case "m":
			i = handleMinutes(layoutBytes, i)
		case "s":
			i = handleSeconds(layoutBytes, i)
		}
		if i == -1 {
			continue
		}

		unitCount := dur / layoutDurMap[t]
		dur %= layoutDurMap[t]
		unitCountStr := strconv.FormatInt(int64(unitCount), 10)
		layoutBytes = slices.Insert(layoutBytes, i, []byte(unitCountStr)...)
	}
	return string(layoutBytes), nil
}

func handleMinutes(layoutBytes []byte, i int) int {
	if i+1 >= len(layoutBytes) {
		return i
	}
	next := layoutBytes[i+1]
	if next == byte('s') {
		i = bytes.LastIndexByte(layoutBytes, byte('m'))
	}
	return i
}

func handleSeconds(layoutBytes []byte, i int) int {
	if i-1 < 0 {
		return i
	}
	prev := layoutBytes[i-1]
	if prev == byte('m') || prev == byte('u') || prev == byte('n') {
		i += bytes.IndexByte(layoutBytes[i+1:], byte('s')) + 1
		return handleSeconds(layoutBytes, i)
	}
	return i
}
