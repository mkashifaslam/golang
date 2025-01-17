package lib

import (
	"strings"
)

type Str2NumMap map[string]int64

func getStr2NumMap() Str2NumMap {
	var str2numMap = make(Str2NumMap)

	var (
		lNum = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		bNum = []string{"hundred", "thousand", "million", "billion", "trillion", "quadrillion"}
	)

	for i := 0; i < len(lNum); i++ {
		str2numMap[lNum[i]] = int64(i + 1)
	}

	var (
		multi  = int64(1000)
		bStart = int64(100)
		bLast  = multi
	)

	for i := 0; i < len(bNum); i++ {
		if i == 0 {
			str2numMap[bNum[i]] = bStart
		} else if i == 1 {
			str2numMap[bNum[i]] = multi
		} else {
			str2numMap[bNum[i]] = multi * bLast
			bLast = str2numMap[bNum[i]]
		}
	}

	return str2numMap
}

func Str2Num(str string) int64 {
	str2numMap := getStr2NumMap()

	strValues := strings.Split(str, " ")

	var result int64 = 1
	for _, strValue := range strValues {
		result *= str2numMap[strings.ToLower(strValue)]
	}

	return result
}
