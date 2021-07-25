package faker

import (
	"fmt"
	"testing"
)

//func ExampleColor() {
//	Seed(0)
//	fmt.Println(Color())
//	// Output: Hello
//}

func TestColor(t *testing.T) {
	Seed(0)
	fmt.Println(Color())
}
