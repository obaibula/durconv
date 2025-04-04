package durfmt

import (
	"errors"
	"math"
	"strconv"
	"strings"
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

var precedenseMap = map[rune]int{
	'h': 1,
	'd': 2,
	'm': 3,
}

func String(layout string, dur time.Duration) (string, error) {
	var b strings.Builder
	prec := math.MaxInt
	for _, r := range layout {
		switch r {
		case 'h':
			err := validatePrecedense(r, &prec)
			if err != nil {
				return "", err
			}
			h := dur / Hour
			dur %= Hour
			b.WriteString(strconv.FormatInt(int64(h), 10))
		case 'd':
			err := validatePrecedense(r, &prec)
			if err != nil {
				return "", err
			}
			d := dur / Day
			dur %= Day
			b.WriteString(strconv.FormatInt(int64(d), 10))
		case 'm':
			err := validatePrecedense(r, &prec)
			if err != nil {
				return "", err
			}
			m := dur / Month
			dur %= Month
			b.WriteString(strconv.FormatInt(int64(m), 10))
		}
		b.WriteRune(r)
	}
	return b.String(), nil
}

func validatePrecedense(r rune, prec *int) error {
	rI := precedenseMap[r]
	if rI > *prec {
		return errors.New("invalid precedense")
	}
	*prec = rI
	return nil
}
