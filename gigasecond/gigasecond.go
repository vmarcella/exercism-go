// Calculate the moment when someone had lived for a gigasecond
package gigasecond

// import path for the time package from the standard library
import "time"

// Calculate the moment someone has lived for a gigasecond
// Given a starting date
func AddGigasecond(t time.Time) time.Time {
	duration := time.Duration(1000000000000000000)
	t = t.Add(duration)
	return t
}
