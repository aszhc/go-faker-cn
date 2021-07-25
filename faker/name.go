package faker

import "math/rand"

func Name() string {
	return name(globalFaker.Rand)
}

func name(r *rand.Rand) string {
	firstName := getRandValue(r, []string{"person", "zh_CN_first"})
	lastName := getRandValue(r, []string{"person", "zh_CN_last"})
	return lastName + firstName
}
