package faker

import (
	crand "crypto/rand"
	"encoding/binary"
	rand "math/rand"
)

type Faker struct {
	Rand *rand.Rand
}

var globalFaker *Faker = New(0)

func New(seed int64) *Faker {
	if seed == 0 {
		binary.Read(crand.Reader, binary.BigEndian, &seed) // 生成一个随机的seed
	}
	return &Faker{Rand: rand.New(rand.NewSource(seed))}
}
