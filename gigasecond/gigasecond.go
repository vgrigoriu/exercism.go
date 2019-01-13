// Package gigasecond solves the Gigasecond problem from Exercism.
package gigasecond

import "time"

var gigasecond time.Duration

func init() {
	// ignore errors
	gigasecond, _ = time.ParseDuration("1000000000s")
}

// AddGigasecond adds a gigasecond (10^9s) to t.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
