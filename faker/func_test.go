package faker

import (
	"fmt"
	"testing"
)

func RangeTestStr(RangeTestFuncStr func() string) {
	for i := 0; i < 10; i++ {
		fmt.Println(RangeTestFuncStr())
	}
}

func RangeTestInt(RangeTestFuncInt func() int) {
	for i := 0; i < 10; i++ {
		fmt.Println(RangeTestFuncInt())
	}
}

func TestName(t *testing.T) {
	RangeTestStr(Name)
}

func TestMacAddress(t *testing.T) {
	RangeTestStr(MacAddress)
}

func TestIPAddress(t *testing.T) {
	RangeTestStr(IPv4Address)
	RangeTestStr(IPv6Address)
}

func TestHttpStatusCode(t *testing.T) {
	RangeTestInt(HTTPStatusCode)
}

func TestUserAgent(t *testing.T) {
	RangeTestStr(UserAgent)
}

func TestMobilePhone(t *testing.T) {
	RangeTestStr(MobilePhone)
}
