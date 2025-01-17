package lib

import "slices"

type Str2Num map[string]int64

var (
	str2NumDic = make(Str2Num)
	fNum       = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	mNum       = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	eNum       = []string{"thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	lNum       = []string{"hundred", "thousand", "million", "billion", "trillion", "quadrillion"}
)

type Rank string

var (
	First  Rank = "fn"
	Middle Rank = "mn"
	Extend Rank = "en"
	Last   Rank = "ln"
)

func GetRank(rank Rank, n string) bool {
	switch rank {
	case First:
		return slices.Contains(fNum, n)
	case Middle:
		return slices.Contains(mNum, n)
	case Extend:
		return slices.Contains(eNum, n)
	case Last:
		return slices.Contains(lNum, n)
	default:
		return false
	}
}

func GetNumbersDict() Str2Num {
	rankF()
	rankM()
	rankE()
	rankL()
	return str2NumDic
}

func rankF() {
	for i := 0; i < len(fNum); i++ {
		str2NumDic[fNum[i]] = int64(1 + i)
	}
}

func rankM() {
	for i := 0; i < len(mNum); i++ {
		str2NumDic[mNum[i]] = int64(i + 10)
	}
}

func rankE() {
	for i := 0; i < len(eNum); i++ {
		str2NumDic[eNum[i]] = int64((i+1)*10 + 20)
	}
}

func rankL() {
	var (
		multi  = int64(1000)
		bStart = int64(100)
		bLast  = multi
	)

	for i := 0; i < len(lNum); i++ {
		if i == 0 {
			str2NumDic[lNum[i]] = bStart
		} else if i == 1 {
			str2NumDic[lNum[i]] = multi
		} else {
			str2NumDic[lNum[i]] = multi * bLast
			bLast = str2NumDic[lNum[i]]
		}
	}
}
