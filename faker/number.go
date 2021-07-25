package faker

import "math/rand"

func Number(min int, max int) int {
	return number(globalFaker.Rand, min, max)
}

func number(r *rand.Rand, min, max int) int {
	return randIntRange(r, min, max)
}
