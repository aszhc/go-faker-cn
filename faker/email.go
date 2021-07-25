package faker

import (
	"fmt"
	"math/rand"
)

func Email() string {
	return email(globalFaker.Rand)
}

func email(r *rand.Rand) string {
	var email string
	randomNumber := fmt.Sprintf("%04d", Number(0, 9999))
	email = getRandValue(r, []string{"person", "en_US_first"}) + getRandValue(r, []string{"person", "en_US_last"}) + randomNumber
	return email
}
