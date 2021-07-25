package faker

import (
	"strconv"
)

// Age between 0 ~ 108
func Age() string {
	return age()
}

func age() string {
	age := strconv.Itoa(Number(0, 108))
	return age
}
