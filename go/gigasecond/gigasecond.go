// Package clause.
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4 // find the value in gigasecond_test.go

const GigaSecond = "1000000000s"

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	dur, _ := time.ParseDuration(GigaSecond)
	res := t.Add(dur)
	return res
}
