package main

import (
	"fmt"
	"math"
	"time"
)

const (
	lexoSize = 10
	MinLexo  = '0'
	MidLexo  = 'U'
	MaxLexo  = 'z'
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

func GetRankBetween(firstRank string, secondRank string) string {
	for len(firstRank) != len(secondRank) {
		if len(firstRank) > len(secondRank) {
			secondRank += string(MinLexo)
		} else {
			firstRank += string(MinLexo)
		}
	}
	diff := calcTotalDiff(firstRank, secondRank)
	newRank := ""
	if diff <= 1 {
		newRank = firstRank + string(MidLexo)
	} else {
		diff = diff / 2
		newRank = genRankByDiff(firstRank, diff)
	}
	return newRank
}

func estimateEndRank(taskNumber int64, startRank string) string {
	endRank := ""
	for i := 0; i < len(startRank); i++ {
		endRank += string(MaxLexo)
	}
	for diff := int64(0); diff <= taskNumber; diff = calcTotalDiff(startRank, endRank) {
		startRank += string(MinLexo)
		endRank += string(MaxLexo)
	}
	return endRank
}

func generateRanks(taskNumber int64, startRank string) []string {
	endRank := estimateEndRank(taskNumber, startRank)
	for len(startRank) != len(endRank) {
		startRank += string(MinLexo)
	}
	totalDiff := calcTotalDiff(startRank, endRank)
	diff := totalDiff / (taskNumber + 1)
	ranks := make([]string, 0, taskNumber)
	for rank := genRankByDiff(startRank, diff); int64(len(ranks)) < taskNumber; rank = genRankByDiff(rank, diff) {
		ranks = append(ranks, rank)
	}
	return ranks
}

func main() {
	startTime := time.Now()
	ranks := generateRanks(100000, "0")
	endTime := time.Now()
	fmt.Printf("%+v\n", ranks)
	fmt.Printf("%d\n", len(ranks))
	fmt.Printf("%d", endTime.Sub(startTime).Milliseconds())
}
