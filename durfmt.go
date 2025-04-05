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
	tokenToDur   = map[string]time.Duration{
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
		tPos := bytes.Index(layoutBytes, []byte(t))
		if tPos == -1 {
			continue
		}

		switch t {
		case "m":
			tPos = resolveMinutesPos(layoutBytes, tPos)
		case "s":
			tPos = resolveSecondsPos(layoutBytes, tPos)
		}
		if tPos == -1 {
			continue
		}

		unitCount := dur / tokenToDur[t]
		dur %= tokenToDur[t]
		unitCountStr := strconv.FormatInt(int64(unitCount), 10)
		layoutBytes = slices.Insert(layoutBytes, tPos, []byte(unitCountStr)...)
	}
	return string(layoutBytes), nil
}

func resolveMinutesPos(layoutBytes []byte, tPos int) int {
	if tPos+1 >= len(layoutBytes) {
		return tPos
	}
	nextByte := layoutBytes[tPos+1]
	if nextByte == byte('s') {
		tPos = bytes.LastIndexByte(layoutBytes, byte('m'))
	}
	return tPos
}

func resolveSecondsPos(layoutBytes []byte, tPos int) int {
	if tPos-1 < 0 {
		return tPos
	}
	prevByte := layoutBytes[tPos-1]
	if prevByte == byte('m') || prevByte == byte('u') || prevByte == byte('n') {
		nextSTokenPos := bytes.IndexByte(layoutBytes[tPos+1:], byte('s'))
		if nextSTokenPos == -1 {
			return -1
		}
		tPos += nextSTokenPos + 1
		return resolveSecondsPos(layoutBytes, tPos)
	}
	return tPos
}
