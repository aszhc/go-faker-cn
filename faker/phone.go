package faker

import (
	"fmt"
	"math/rand"
)

func MobilePhone() string {
	return mobilePhone(globalFaker.Rand)
}

func mobilePhone(r *rand.Rand) string {
	phonePrefix := getRandValue(r, []string{"phone", "zh_CN_mobile"})
	phone8Num := fmt.Sprintf("%08d", Number(1, 99999999))
	return phonePrefix + phone8Num
}
