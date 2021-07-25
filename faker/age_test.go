package faker

import (
	"fmt"
)

func ExampleAge() {
	Seed(0)
	fmt.Println(Age())
	// Output: 47
}
