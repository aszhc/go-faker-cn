package faker

import (
	"math/rand"
	"time"
)

// Date will generate a random time.Time struct
func Date() time.Time { return date(globalFaker.Rand) }

func date(r *rand.Rand) time.Time {
	return time.Date(year(r), time.Month(number(r, 1, 12)), day(r), hour(r), minute(r), second(r), nanoSecond(r), time.UTC)
}

func year(r *rand.Rand) int { return number(r, 1900, time.Now().Year()) }

func day(r *rand.Rand) int { return number(r, 1, 31) }

func hour(r *rand.Rand) int { return number(r, 0, 23) }

func minute(r *rand.Rand) int { return number(r, 0, 59) }

func second(r *rand.Rand) int { return number(r, 0, 59) }

func nanoSecond(r *rand.Rand) int { return number(r, 0, 999999999) }
