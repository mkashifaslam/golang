package lib

import (
	"strings"
)

func StrToNumber(str string) int64 {
	strValues := strings.Split(str, " ")
	numMap := GetNumbersDict()

	var result int64
	var carry int64
	for _, strValue := range strValues {
		var strVal = strings.ToLower(strValue)
		var num = numMap[strings.ToLower(strValue)]
		if GetRank(Last, strVal) {
			carry *= num
			result += carry
			carry = 0
		} else {
			carry += num
		}
	}

	if carry != 0 {
		result += carry
	}

	return result
}
