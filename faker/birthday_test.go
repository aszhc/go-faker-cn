package faker

import "fmt"

func ExampleBirthDay() {
	Seed(0)
	fmt.Println(BirthDay())
	// Output: 19600406
}
