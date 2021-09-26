package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaseFunction(t *testing.T) {
	testCases := []struct {
		name        		string
		params		 		Params
		inputCase			[]string
		outputCaseS			[]string // Flag -f -s
		outputCaseIU		[]int // Flag -d && -i && -u && default
		outputCaseCount		[]CountU
		input				[]string
	} {
		{
			name:        	"Flag -c",
			params: 		Params{C: true},
			input:       	[]string{"I love music.", "I love music.", "i love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			outputCaseCount:[]CountU{{count: 2, num: 0}, {count:1, num:2}, {count: 1, num: 3}, {count: 2, num: 4}, {count: 1, num: 6}},
			outputCaseIU: 	[]int{},
		}, {
			name:        	"Flag -c with -i",
			params: 		Params{C: true, I: true},
			input:       	[]string{"I love music.", "I love music.", "i love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			outputCaseCount:[]CountU{{count: 3, num: 0}, {count: 1, num: 3}, {count: 2, num: 4}, {count: 1, num: 6}},
			outputCaseIU: 	[]int{},
		}, {
			name:        	"Flag -d",
			params: 		Params{D: true},
			input:       	[]string{"i love music.", "I love music.", "i Love music.", "", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			outputCaseIU: 	[]int{4, 6},
		}, {
			name:        	"Flag -d with -i",
			params: 		Params{D: true, I: true},
			input:       	[]string{"i love music.", "I love music.", "i Love music.", "", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			outputCaseIU: 	[]int{1, 4, 6},
		}, {
			name:        	"Flag -u",
			params: 		Params{U: true},
			input:       	[]string{"i love music.", "I love music.", "i Love music.", " ", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			outputCaseIU: 	[]int{1, 6},
		}, {
			name:        	"Flag -u with -i",
			params: 		Params{U: true, I: true},
			input:       	[]string{"i love music.", "I love music.", "i Love music.", "", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			outputCaseIU: 	[]int{3, 5, 7},
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.params.C {
				uniq := cUnique(tc.input, tc.params)
				assert.Equal(t, uniq, tc.outputCaseCount)
			} else if tc.params.D {
				uniq := dRepeated(tc.input, tc.params)
				assert.Equal(t, uniq, tc.outputCaseIU)
			} else if tc.params.U {
				uniq := uUnique(tc.input, tc.params)
				assert.Equal(t, uniq, tc.outputCaseIU)
			}
		})
	}
}