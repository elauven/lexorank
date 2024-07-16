package lexorank

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	MinLexo  = '0'
	MaxLexo  = 'z'
	MidLexo  = (MinLexo + MaxLexo) / 2
	lexoSize = MaxLexo - MinLexo + 1
)

func genLexoCodes(data string) []int64 {
	codes := make([]int64, 0, len(data))
	for _, char := range data {
		codes = append(codes, int64(char))
	}
	return codes
}

func reverse(data string) string {
	reversedData := ""
	for _, char := range data {
		reversedData = string(char) + reversedData
	}
	return reversedData
}

func calcTotalDiff(firstRank string, secondRank string) int64 {
	firstPositionCodes := genLexoCodes(firstRank)
	secondPositionCodes := genLexoCodes(secondRank)
	var diff int64 = 0
	for idx := len(firstRank) - 1; idx >= 0; idx-- {
		firstCode := firstPositionCodes[idx]
		secondCode := secondPositionCodes[idx]
		if secondCode < firstCode {
			secondCode += lexoSize
			secondPositionCodes[idx-1] -= 1
		}
		powRes := int64(math.Pow(float64(lexoSize), float64(len(firstRank)-idx-1)))
		diff += (secondCode - firstCode) * powRes
	}
	return diff
}

func genRankByDiff(firstRank string, diff int64) string {
	newRank := ""
	var offset int64 = 0
	for idx := 0; idx < len(firstRank); idx++ {
		diffInSymbols := diff / int64(math.Pow(lexoSize, float64(idx))) % lexoSize
		newRankCode := int64(firstRank[len(firstRank)-idx-1]) + diffInSymbols + offset
		offset = 0
		if newRankCode > MaxLexo {
			offset++
			newRankCode -= lexoSize
		}
		newRank += string(rune(newRankCode))
	}
	newRank = reverse(newRank)
	return newRank
}

func validateRank(rank string) bool {
	for _, ch := range []byte(rank) {
		if ch < MinLexo || ch > MaxLexo {
			return false
		}
	}
	return true
}

func Rank(prev string, next string) (string, error) {
	for len(prev) != len(next) {
		if len(prev) > len(next) {
			next += string(MinLexo)
		} else {
			prev += string(MinLexo)
		}
	}
	if !validateRank(prev) || !validateRank(next) {
		return "", errors.New("invalid lexorank")
	}
	if prev > next {
		return "", errors.New("prev rank is greater than next rank")
	}
	diff := calcTotalDiff(prev, next)
	newRank := ""
	if diff <= 1 {
		newRank = prev + string(MidLexo)
	} else {
		diff = diff / 2
		newRank = genRankByDiff(prev, diff)
	}
	return newRank, nil
}

func RankN(prev, next string, n int) ([]string, error) {
	mid, err := Rank(prev, next)
	if err != nil {
		return nil, err
	}
	suffixRankLen := len(strconv.Itoa(n))
	suffixRankFormat := fmt.Sprintf("%%0%dd", suffixRankLen)
	ranks := make([]string, 0, n)
	for i := 0; i < n; i++ {
		ranks = append(ranks, mid+fmt.Sprintf(suffixRankFormat, i))
	}
	return ranks, nil
}
