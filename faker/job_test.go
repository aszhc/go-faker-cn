package faker

import (
	"fmt"
	"testing"
)

func TestJob(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Job())
	}
}
