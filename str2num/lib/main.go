package lib

import (
	"strings"
)

var (
	mSep = ","
	sSep = " "
)

func StrToNumber(str string) int64 {
	var strValues = make([]string, 0)

	if strings.Contains(str, mSep) {
		strValues = strings.Split(str, mSep)
	} else {
		strValues = append(strValues, str)
	}

	numMap := GetNumbersDict()

	var result int64
	for _, strValue := range strValues {
		result += calcResult(strings.Split(strValue, sSep), numMap)
	}

	return result
}

func calcResult(strValues []string, numMap Str2Num) int64 {
	var result int64
	var carry int64
	for _, strValue := range strValues {
		var strVal = strings.ToLower(strValue)
		var num = numMap[strings.ToLower(strValue)]

		if GetRank(Last, strVal) {
			carry *= num
		} else {
			carry += num
		}
	}

	if carry != 0 {
		result += carry
	}

	return result
}
