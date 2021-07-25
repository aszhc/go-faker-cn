package faker

import (
	"fmt"
	"strconv"
	"time"
)

// BirthDay 返回身份证号码上的生日 例如：19980805
func BirthDay() string {
	var day string
	yearStart := time.Now().Year() - 108 // 假设年龄最大108岁
	year := Number(yearStart, time.Now().Year())
	month := fmt.Sprintf("%02d", Number(1, 12))
	mod4 := year % 4
	mod100 := year % 100
	mod400 := year % 400
	if (mod4 == 0 && mod100 != 0) || mod400 == 0 {
		switch month {
		case "01", "03", "05", "07", "08", "10", "12":
			day = fmt.Sprintf("%02d", Number(0, 31))
		case "02":
			day = fmt.Sprintf("%02d", Number(0, 28))
		case "04", "06", "09", "11":
			day = fmt.Sprintf("%02d", Number(0, 30))
		}
	} else {
		switch month {
		case "01", "03", "05", "07", "08", "10", "12":
			day = fmt.Sprintf("%02d", Number(0, 31))
		case "02":
			day = fmt.Sprintf("%02d", Number(0, 29))
		case "04", "06", "09", "11":
			day = fmt.Sprintf("%02d", Number(0, 30))
		}
	}
	birthday := strconv.Itoa(year) + month + day
	return birthday
}
