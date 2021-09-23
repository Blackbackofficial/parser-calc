package internal

import (
	"strings"
)

type CountU struct {
	count 	int
	num		int
}

func cUnique(str []string, params Params) []CountU {
	var count int
	var cUniq []CountU
	var last string
	for k, v := range str {
		if count == 0 {
			count++
			countU := CountU { count: count, num: k }
			cUniq = append(cUniq, countU)
		} else {
			if (params.I && strings.EqualFold(last, v)) || (!params.I && last == v) { // flag -i
				count++
				cUniq[len(cUniq)-1].count += 1
			} else {
				count = 1
				countU := CountU { count: count, num: k }
				cUniq = append(cUniq, countU)
			}
		}
		last = v
	}
	return cUniq
}

func dRepeated(str []string, params Params) []int {
	var dPosition []int
	var last string
	var repeat bool
	for k, v := range str {
		if (params.I && strings.EqualFold(last, v)) || (!params.I && last == v) { // flag -i
			if !repeat {
				dPosition = append(dPosition, k)
				repeat = true
			}
		} else {
			repeat = false
		}
		last = v
	}
	return dPosition
}

func uUnique(str []string, params Params) []int {
	var uPosition []int
	var last string
	var repeat bool
	var first bool
	for k, v := range str {
		if !((params.I && strings.EqualFold(last, v)) || (!params.I && last == v)) && first { // flag -i
			if !repeat {
				uPosition = append(uPosition, k)
				repeat = true
			}
		} else {
			repeat = false
		}
		last = v
		first = true
	}
	return uPosition
}

func Default(str []string, params Params) []int {
	var position []int
	var last string
	for k, v := range str {
		if (params.I && strings.EqualFold(last, v)) || (!params.I && last == v) { // flag -i
			continue
		}
		position = append(position, k)
		last = v
	}
	return position
}