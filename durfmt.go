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

const layoutTokens = "yMwdhms"

var layoutDurMap = map[rune]time.Duration{
	's': Second,
	'm': Minute,
	'h': Hour,
	'd': Day,
	'w': Week,
	'M': Month,
	'y': Year,
}

func String(layout string, dur time.Duration) (string, error) {
	buf := []byte(layout)
	for _, r := range layoutTokens {
		i := bytes.IndexRune(buf, r)
		if i == -1 {
			continue
		}
		n := dur / layoutDurMap[r]
		dur %= layoutDurMap[r]
		conc := strconv.FormatInt(int64(n), 10)
		buf = slices.Insert(buf, i, []byte(conc)...)
	}
	return string(buf), nil
}
