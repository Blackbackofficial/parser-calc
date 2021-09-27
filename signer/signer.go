package main

import (
	"sort"
	"strings"
)

// ExecutePipeline which provides us with pipelining of worker functions that do something
func ExecutePipeline(jobs ...job) {

}

// SingleHash considers the value crc32 (data) + "~" + crc32 (md5 (data)) (concatenation of two strings through ~), where data is what came to the input (in fact, numbers from the first function)
func SingleHash(stdIn, out chan interface{}) {

}

// MultiHash counts the value crc32 (th + data)) (concatenation of a digit reduced to a string and a string), where th = 0..5, then takes the concatenation of the results in the order of calculation (0..5), where data is what came to the input
func MultiHash(stdIn, out chan interface{}) {

}


// CombineResults gets all results, sorts (https://golang.org/pkg/sort/), concatenates the sorted result with _ (underscore) into one line
func CombineResults(stdIn, stdOut chan interface{}) {
	var concat []string
	for v := range stdIn {
		concat = append(concat, v.(string))
	}

	sort.Strings(concat)
	stdOut <- strings.Join(concat, "_")
}