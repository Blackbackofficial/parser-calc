package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

var th int = 6 // 6 process: th = 0...5

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
func SingleHash(stdIn chan interface{}, stdOut chan interface{}) {
	wg := &sync.WaitGroup{}
	for input := range stdIn {
		wg.Add(1)
		strMd5 := DataSignerMd5(strconv.Itoa(input.(int)))

		go func(str string) {
			wg1 := &sync.WaitGroup{}
			wg1.Add(2) // 2 process
			str1 := dataSignerCrc32(wg1, &str)
			str2 := dataSignerCrc32(wg1, &strMd5)
			wg1.Wait()
			defer wg.Done()
			stdOut <- strings.Join([]string{<-str1, <-str2}, "~")
		}(strconv.Itoa(input.(int)))
	}
	wg.Wait()
}

func dataSignerCrc32(wg1 *sync.WaitGroup, str *string) chan string {
	out := make(chan string)
	go func(data string, out chan <-string) {
		out <- DataSignerCrc32(*str)
	}(*str, out)
	wg1.Done()
	return out
}

// MultiHash counts the value crc32 (th + data), where th = 0..5, then takes the concatenation of the results in the order of calculation (0..5), where data is what came to the input
func MultiHash(stdIn chan interface{}, stdOut chan interface{}) {
	wg := &sync.WaitGroup{}
	for d := range stdIn {
		wg.Add(1)
		go func(stdOut chan interface{}, str string) {
			strSigner := make([]string, th)
			wg1 := &sync.WaitGroup{}
			wg1.Add(th)
			for v := 0; v < th; v++ {
				str1 := strconv.Itoa(v)+str
				go func(wg1 *sync.WaitGroup, str, out *string) {
					*out = DataSignerCrc32(*str)
					wg1.Done()
				}(wg1, &str1, &strSigner[v])
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