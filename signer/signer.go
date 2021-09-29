package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

const TH = 6 // 6 process: th = 0...5

// ExecutePipeline which provides us with pipelining of worker functions that do something
func ExecutePipeline(jobs ...job) {
	stdIn := make(chan interface{})

	wg := &sync.WaitGroup{}
	for _, someJob := range jobs {
		stdOut := make(chan interface{})
		wg.Add(1)
		go func(stdIn, stdOut chan interface{}, j job) {
			j(stdIn, stdOut)
			close(stdOut)
			wg.Done()
		}(stdIn, stdOut, someJob)
		stdIn = stdOut
	}
	wg.Wait()
}

// SingleHash considers the value crc32 (data) + "~" + crc32 (md5 (data)) (concatenation of two strings through ~), where data is what came to the input (in fact, numbers from the first function)
func SingleHash(stdIn, stdOut chan interface{}) {
	wg := &sync.WaitGroup{}
	for input := range stdIn {
		wg.Add(1)
		strMd5 := DataSignerMd5(strconv.Itoa(input.(int)))

		go func(str string) {
			defer wg.Done()
			str1 := make(chan string)
			str2 := make(chan string)
			go dataSignerCrc32(str, str1)
			go dataSignerCrc32(strMd5, str2)
			stdOut <- strings.Join([]string{<-str1, <-str2}, "~")
		}(strconv.Itoa(input.(int)))
	}
	wg.Wait()
}

func dataSignerCrc32(str string, out chan <-string) {
	out <- DataSignerCrc32(str)
}

// MultiHash counts the value crc32 (th + data), where th = 0..5, then takes the concatenation of the results in the order of calculation (0..5), where data is what came to the input
func MultiHash(stdIn, stdOut chan interface{}) {
	wg := &sync.WaitGroup{}
	for d := range stdIn {
		strSigner := make([]string, TH)
		channels := make([]chan string, TH)
		wg.Add(1)
		for v := 0; v < TH; v++ {
			channels[v] = make(chan string)
		}

		for c := range channels {
			go dataSignerCrc32(strconv.Itoa(c)+d.(string), channels[c])
		}

		go func() {
			for k,v := range channels {
				strSigner[k] = <-v
			}
			stdOut <- strings.Join(strSigner, "")
			wg.Done()
		}()
	}
	wg.Wait()
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