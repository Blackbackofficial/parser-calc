package internal

import (
	"log"
	"strings"
)

type CountU struct {
	count 	int
	num		int
}

// Flag -c && -i
func cUnique(arrStr []string, params Params) []CountU {
	var count int
	var cUniq []CountU
	var last string
	for k, v := range arrStr {
		if (params.I && strings.EqualFold(last, v)) || (!params.I && last == v) { // flag -i
			count++
			cUniq[len(cUniq)-1].count += 1
		} else {
			count = 1
			countU := CountU { count: count, num: k }
			cUniq = append(cUniq, countU)
		}

		last = v
	}
	return cUniq
}

// Flag -d && -i
func dRepeated(arrStr []string, params Params) []int {
	var dPosition []int
	var last string
	var repeat bool
	for k, v := range arrStr {
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

// Flag -u && -i
func uUnique(arrStr []string, params Params) []int {
	var uPosition []int
	var last string
	var repeat bool
	var first bool
	for k, v := range arrStr {
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

// Flag -f
func cutStrF(arrStr []string, numField int) []string {
	if numField < 0  {
		log.Fatalln("Incorrect num field")
	}

	var newSlice []string
	for _, v := range arrStr {
		var str string
		s := strings.Split(v, " ")
		for k, v := range s {
			if k < numField || v == ""{
				continue
			}
			if len(s)-1 == k {
				str += v
			} else {
				str += v + " "
			}
		}
		newSlice = append(newSlice, str)
	}
	return newSlice
}

// Flag -s
func cutCharS(arrStr []string, numChar int) []string {
	if numChar < 0  {
		log.Fatalln("Incorrect num field")
	}

	var newSlice []string
	for _, v := range arrStr {
		var str string
		s := strings.Split(v, "")
		for k, v := range s {
			if k < numChar || v == ""{
				continue
			}
			str += v
		}
		newSlice = append(newSlice, str)
	}
	return newSlice
}

// No flags -> default
func defaultF(arrStr []string, params Params) []int {
	var position []int
	var last string
	var first bool
	for k, v := range arrStr {
		if (params.I && strings.EqualFold(last, v)) || (!params.I && last == v) && first { // flag -i
			continue
		}
		first = true
		position = append(position, k)
		last = v
	}
	return position
}