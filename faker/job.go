package faker

import "math/rand"

func Job() string {
	return job(globalFaker.Rand)
}

func job(r *rand.Rand) string {
	return getRandValue(r, []string{"job", "jobs"})
}
