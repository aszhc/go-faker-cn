package faker

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"

	"github.com/aszhc/go-faker-cn/data"
)

func Seed(seed int64) {
	if seed == 0 {
		binary.Read(crand.Reader, binary.BigEndian, &seed) // 生成随机的seed
		globalFaker.Rand.Seed(seed)
	} else {
		globalFaker.Rand.Seed(seed) // 使用传入的seed	 非真正随机数
	}
}

func intDataCheck(dataVal []string) bool {
	if len(dataVal) != 2 {
		return false
	}

	_, checkOk := data.IntData[dataVal[0]]
	if checkOk {
		_, checkOk = data.IntData[dataVal[0]][dataVal[1]]
	}

	return checkOk
}

// Check if in lib
func dataCheck(dataVal []string) bool {
	var checkOk bool

	if len(dataVal) == 2 {
		_, checkOk = data.Data[dataVal[0]]
		if checkOk {
			_, checkOk = data.Data[dataVal[0]][dataVal[1]]
		}
	}

	return checkOk
}

// Get Random Value
func getRandValue(r *rand.Rand, dataVal []string) string {
	if !dataCheck(dataVal) {
		return ""
	}
	return data.Data[dataVal[0]][dataVal[1]][r.Intn(len(data.Data[dataVal[0]][dataVal[1]]))]
}

func getRandIntValue(r *rand.Rand, dataVal []string) int {
	if !intDataCheck(dataVal) {
		return 0
	}
	return data.IntData[dataVal[0]][dataVal[1]][r.Intn(len(data.IntData[dataVal[0]][dataVal[1]]))]
}

// 生成介于min和max之间的随机int数字
func randIntRange(r *rand.Rand, min int, max int) int {
	if min == max {
		return min
	}
	return r.Intn((max+1)-min) + min
}
