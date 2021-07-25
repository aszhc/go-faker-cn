package faker

import (
	"fmt"
	"math/rand"
	"strconv"
)

func IPv4Address() string {
	return ipv4Address(globalFaker.Rand)
}

func ipv4Address(r *rand.Rand) string {
	num := func() int { return r.Intn(256) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

func IPv6Address() string {
	return ipv6Address(globalFaker.Rand)
}

func ipv6Address(r *rand.Rand) string {
	num := func() int { return r.Intn(65536) }
	return fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x", num(), num(), num(), num(), num(), num(), num(), num()) // %x 16进制
}

func MacAddress() string { return macAddress(globalFaker.Rand) }

func macAddress(r *rand.Rand) string {
	num := 255

	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", r.Intn(num), r.Intn(num), r.Intn(num), r.Intn(num), r.Intn(num), r.Intn(num))
}

// HTTPStatusCode will generate a random status code
func HTTPStatusCode() int { return httpStatusCode(globalFaker.Rand) }

func httpStatusCode(r *rand.Rand) int {
	return getRandIntValue(r, []string{"status_code", "general"})
}

// UserAgent will generate a random broswer user agent
func UserAgent() string { return userAgent(globalFaker.Rand) }

func userAgent(r *rand.Rand) string {
	randNum := randIntRange(r, 0, 4)
	switch randNum {
	case 0:
		return chromeUserAgent(r)
	case 1:
		return firefoxUserAgent(r)
	case 2:
		return safariUserAgent(r)
	case 3:
		return operaUserAgent(r)
	default:
		return chromeUserAgent(r)
	}
}

func chromeUserAgent(r *rand.Rand) string {
	randNum1 := strconv.Itoa(randIntRange(r, 531, 536)) + strconv.Itoa(randIntRange(r, 0, 2))
	randNum2 := strconv.Itoa(randIntRange(r, 36, 40))
	randNum3 := strconv.Itoa(randIntRange(r, 800, 899))
	return "Mozilla/5.0 " + "(" + randomPlatform(r) + ") AppleWebKit/" + randNum1 + " (KHTML, like Gecko) Chrome/" + randNum2 + ".0." + randNum3 + ".0 Mobile Safari/" + randNum1
}

func firefoxUserAgent(r *rand.Rand) string {
	ver := "Gecko/" + date(r).Format("2006-02-01") + " Firefox/" + strconv.Itoa(randIntRange(r, 35, 37)) + ".0"
	platforms := []string{
		"(" + windowsPlatformToken(r) + "; " + "en-US" + "; rv:1.9." + strconv.Itoa(randIntRange(r, 0, 3)) + ".20) " + ver,
		"(" + linuxPlatformToken(r) + "; rv:" + strconv.Itoa(randIntRange(r, 5, 8)) + ".0) " + ver,
		"(" + macPlatformToken(r) + " rv:" + strconv.Itoa(randIntRange(r, 2, 7)) + ".0) " + ver,
	}

	return "Mozilla/5.0 " + randomString(r, platforms)
}

func safariUserAgent(r *rand.Rand) string {
	randNum := strconv.Itoa(randIntRange(r, 531, 536)) + "." + strconv.Itoa(randIntRange(r, 1, 51)) + "." + strconv.Itoa(randIntRange(r, 1, 8))
	ver := strconv.Itoa(randIntRange(r, 4, 6)) + "." + strconv.Itoa(randIntRange(r, 0, 2))

	mobileDevices := []string{
		"iPhone; CPU iPhone OS",
		"iPad; CPU OS",
	}

	platforms := []string{
		"(Windows; U; " + windowsPlatformToken(r) + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + macPlatformToken(r) + " rv:" + strconv.Itoa(randIntRange(r, 4, 7)) + ".0; en-US) AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + randomString(r, mobileDevices) + " " + strconv.Itoa(randIntRange(r, 7, 9)) + "_" + strconv.Itoa(randIntRange(r, 0, 3)) + "_" + strconv.Itoa(randIntRange(r, 1, 3)) + " like Mac OS X; " + "en-US" + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + strconv.Itoa(randIntRange(r, 3, 5)) + ".0.5 Mobile/8B" + strconv.Itoa(randIntRange(r, 111, 120)) + " Safari/6" + randNum,
	}

	return "Mozilla/5.0 " + randomString(r, platforms)
}

func operaUserAgent(r *rand.Rand) string {
	platform := "(" + randomPlatform(r) + "; en-US) Presto/2." + strconv.Itoa(randIntRange(r, 8, 13)) + "." + strconv.Itoa(randIntRange(r, 160, 355)) + " Version/" + strconv.Itoa(randIntRange(r, 10, 13)) + ".00"

	return "Opera/" + strconv.Itoa(randIntRange(r, 8, 10)) + "." + strconv.Itoa(randIntRange(r, 10, 99)) + " " + platform
}

// linuxPlatformToken will generate a random linux platform
func linuxPlatformToken(r *rand.Rand) string {
	return "X11; Linux " + getRandValue(r, []string{"computer", "linux_processor"})
}

// macPlatformToken will generate a random mac platform
func macPlatformToken(r *rand.Rand) string {
	return "Macintosh; " + getRandValue(r, []string{"computer", "mac_processor"}) + " Mac OS X 10_" + strconv.Itoa(randIntRange(r, 5, 9)) + "_" + strconv.Itoa(randIntRange(r, 0, 10))
}

// windowsPlatformToken will generate a random windows platform
func windowsPlatformToken(r *rand.Rand) string {
	return getRandValue(r, []string{"computer", "windows_platform"})
}

// randomPlatform will generate a random platform
func randomPlatform(r *rand.Rand) string {
	platforms := []string{
		linuxPlatformToken(r),
		macPlatformToken(r),
		windowsPlatformToken(r),
	}

	return randomString(r, platforms)
}

func randomString(r *rand.Rand, a []string) string {
	size := len(a)
	if size == 0 {
		return ""
	}
	if size == 1 {
		return a[0]
	}
	return a[r.Intn(size)]
}
