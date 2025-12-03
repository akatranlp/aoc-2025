package main

import (
	"aoc-lib/aoc"
	"aoc-lib/its"
	"io"
	"slices"
	"strconv"
)

type Day03 struct{}

var _ aoc.Problem = (*Day03)(nil)

func (*Day03) Part1(r io.Reader) int {
	var res int
	for row := range its.Filter(its.ReaderToIter(r), its.FilterEmptyLines) {
		numbers := slices.Collect(its.Map(slices.Values([]byte(row)), func(c byte) int {
			return int(c) - 48
		}))

		var maxIdx int
		for i := 1; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[maxIdx] {
				maxIdx = i
			}
		}

		max2Idx := maxIdx + 1
		for i := max2Idx; i < len(numbers); i++ {
			if numbers[i] > numbers[max2Idx] {
				max2Idx = i
			}
		}

		maxNum := numbers[maxIdx]*10 + numbers[max2Idx]
		res += maxNum
	}
	return res
}

func (*Day03) Part2(r io.Reader) int {
	var res int
	for row := range its.Filter(its.ReaderToIter(r), its.FilterEmptyLines) {
		numbers := []byte(row)

		maxIdxs := make([]int, 12)
		for i := range maxIdxs {
			if i != 0 {
				maxIdxs[i] = maxIdxs[i-1] + 1
			}
			for j := maxIdxs[i]; j < len(numbers)-(11-i); j++ {
				if numbers[j] > numbers[maxIdxs[i]] {
					maxIdxs[i] = j
				}
			}
		}
		resBytes := make([]byte, 12)
		for i := range maxIdxs {
			resBytes[i] = numbers[maxIdxs[i]]
		}
		maxNum, _ := strconv.Atoi(string(resBytes))
		res += maxNum
	}
	return res
}
