package faker

import (
	"math/rand"
	"strconv"
)

// 中国居民身份证
// 【六位数字地址码】【八位数字出生日期码】【二位数字顺序码】【性别码】【一位校验码】

func AreaCode() string {
	return area(globalFaker.Rand)
}

func area(r *rand.Rand) string {
	return getRandValue(r, []string{"areaCode", "areaCode"})
}

func RandomNumber() string {
	return getRandomNumber(globalFaker.Rand)
}

func getRandomNumber(r *rand.Rand) string {
	var number string
	first := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	second := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	number = first[r.Intn(len(first))] + second[r.Intn(len(second))]
	return number
}

func GenderNumber() string {
	return getGenderNumber(globalFaker.Rand, Gender())
}

// getGenderNumber 生成性别码
func getGenderNumber(r *rand.Rand, g string) string {
	var sex string
	var male = []string{"1", "3", "5", "7", "9"}
	var female = []string{"0", "2", "4", "6", "8"}
	switch g {
	case "男":
		sex = male[r.Intn(len(male))]
	case "女":
		sex = female[r.Intn(len(female))]
	default:
		panic("性别错误, 未知性别")
	}
	return sex
}

func Gender() string {
	return getRandomGender(globalFaker.Rand)
}

func getRandomGender(r *rand.Rand) string {
	var gender = []string{"男", "女"}
	return gender[r.Intn(len(gender))]
}

// 根据前17位数字获取校验码
func getValidateNumber(number string) string {
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}        // 十七位数字本体码权重
	validate := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"} // mod11,对应校验码字符值
	sum := 0
	for i := 0; i < len(number); i++ {
		temp, _ := strconv.Atoi(string(number[i]))
		temp = temp * weight[i]
		sum = sum + temp
	}
	mode := sum % 11
	return validate[mode]
}

func IdCard() string {
	areaCode := AreaCode()
	birthday := BirthDay()
	rNum := RandomNumber()
	sex := GenderNumber()
	num17 := areaCode + birthday + rNum + sex
	validate := getValidateNumber(num17)
	idCard := num17 + validate
	return idCard
}
