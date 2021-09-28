package main

import (
	"sort"
	"strconv"
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
		go func(stdIn, stdOut chan interface{}, j job) {
			j(stdIn, stdOut)
			close(stdOut)
			wg.Done()
		}(stdIn, stdOut, someJob)
		stdIn = stdOut
		stdOut = make(chan interface{})
	}
	wg.Wait()
}

// SingleHash considers the value crc32 (data) + "~" + crc32 (md5 (data)) (concatenation of two strings through ~), where data is what came to the input (in fact, numbers from the first function)
func SingleHash(stdIn chan interface{}, stdOut chan interface{}) {
	wg := &sync.WaitGroup{}
	for input := range stdIn {
		wg.Add(1)
		strMd5 := DataSignerMd5(strconv.Itoa(input.(int)))

		go func(str string) {
			wg1 := &sync.WaitGroup{}
			wg1.Add(2) // 2 process
			var str1, str2 string

			go func(data string) {
				str1 = DataSignerCrc32(data)
				wg1.Done()
			}(str)

			go func(data string) {
				str2 = DataSignerCrc32(strMd5)
				wg1.Done()
			}(str)
			wg1.Wait()
			defer wg.Done()
			stdOut <- strings.Join([]string{str1, str2}, "~")
		}(strconv.Itoa(input.(int)))
	}
	wg.Wait()
}

// MultiHash counts the value crc32 (th + data), where th = 0..5, then takes the concatenation of the results in the order of calculation (0..5), where data is what came to the input
func MultiHash(stdIn chan interface{}, stdOut chan interface{}) {
	wg := &sync.WaitGroup{}
	for d := range stdIn {
		wg.Add(1)
		go func(stdOut chan interface{}, str string) {
			strSigner := make([]string, 6)
			wg1 := &sync.WaitGroup{}
			wg1.Add(6) // 6 process: th = 0...5
			for v := 0; v <= 5; v++ {
				go func(res *string, data string, i string) {
					*res = DataSignerCrc32(i + data)
					wg1.Done()
				}(&strSigner[v], str, strconv.Itoa(v))
			}
			wg1.Wait()
			stdOut <- strings.Join(strSigner, "")
			wg.Done()
		}(stdOut, d.(string))
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