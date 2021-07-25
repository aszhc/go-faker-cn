package faker

import "math/rand"

func Car() string {
	return car(globalFaker.Rand)
}

func car(r *rand.Rand) string {
	return getRandValue(r, []string{"carBrands", "brands"})
}
