package faker

import (
	"fmt"
	"testing"
)

func TestAreaCode(t *testing.T) {
	Seed(0)
	//fmt.Println(getRandomNumber(globalFaker.Rand))
	//fmt.Println(getGenderNumber(globalFaker.Rand, "男"))
	//fmt.Println(getRanderGender(globalFaker.Rand))
	for i := 0; i < 100; i++ {
		fmt.Println(IdCard())
		//fmt.Println(BirthDay())
		//fmt.Println(GenderNumber())
	}
}
