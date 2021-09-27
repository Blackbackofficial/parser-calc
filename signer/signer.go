package main

import (
	"sort"
	"strings"
	"sync"
)

// ExecutePipeline which provides us with pipelining of worker functions that do something
func ExecutePipeline(jobs ...job) {
	stdIn := make(chan interface{})
	stdOut := make(chan interface{})

	wg := &sync.WaitGroup{}
	for _, someJob := range jobs {
		wg.Add(1)
		go func(wg *sync.WaitGroup, in, out chan interface{}, j job) {
			// TODO: mb defer??
			close(out)
			wg.Done()
			j(in, out)
		}(wg, stdIn, stdOut, someJob)
		stdIn = stdOut
		stdOut = make(chan interface{})
	}
	wg.Wait()
}

// SingleHash considers the value crc32 (data) + "~" + crc32 (md5 (data)) (concatenation of two strings through ~), where data is what came to the input (in fact, numbers from the first function)
func SingleHash(stdIn, out chan interface{}) {

}

// MultiHash counts the value crc32 (th + data), where th = 0..5, then takes the concatenation of the results in the order of calculation (0..5), where data is what came to the input
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