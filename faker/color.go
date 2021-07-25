package faker

import "math/rand"

func Color() string {
	return color(globalFaker.Rand)
}

func color(r *rand.Rand) string { return getRandValue(r, []string{"color", "safe"}) }
