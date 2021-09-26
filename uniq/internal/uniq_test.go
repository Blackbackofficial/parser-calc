package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaseFunction(t *testing.T) {
	testCases := []struct {
		name        		string
		params		 		Params
		outputCaseIU		[]int // Flag -d && -i && -u && default
		outputCaseCount		[]CountU // Flag -c && -i
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
			name:        	"Flag -c space",
			params: 		Params{C: true},
			input:       	[]string{" "},
			outputCaseCount:[]CountU{{count:1, num:0}},
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
			name:        	"Flag -d space",
			params: 		Params{D: true, I: true},
			input:       	[]string{""},
			outputCaseIU: 	[]int{0},
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
		}, {
			name:        	"Flag -u space",
			params: 		Params{U: true},
			input:       	[]string{""},
			outputCaseIU: 	nil,
		}, {
			name:        	"Default",
			params: 		Params{},
			input:       	[]string{"I love music.", "I love music.", "i Love music.", "", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			outputCaseIU: 	[]int{0, 2, 3, 5, 7},
		}, {
			name:        	"Default with -i",
			params: 		Params{},
			input:       	[]string{"i love music.", "I love music.", "i Love music.", "", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			outputCaseIU: 	[]int{0, 1, 2, 3, 5, 7},
		}, {
			name:        	"Default",
			params: 		Params{},
			input:       	[]string{""},
			outputCaseIU: 	[]int{0},
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
			} else if !tc.params.U && !tc.params.D && !tc.params.C {
				uniq := defaultF(tc.input, tc.params)
				assert.Equal(t, uniq, tc.outputCaseIU)
			}
		})
	}
}

func TestCutStrings(t *testing.T) {
	testCases := []struct {
		name        		string
		params		 		Params
		inputCase			[]string
		outputCaseS			[]string // Flag -f -s
		input				[]string
	} {
		{
			name:        	"Flag -f num 0",
			params: 		Params{NumFields: 0},
			input:       	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			outputCaseS:	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
		}, {
			name:        	"Flag -f num 1",
			params: 		Params{NumFields: 1},
			input:       	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			outputCaseS:	[]string{"love music.", "love music.", "love music.", "love music of Kartik.", "love music of Kartik.", ""},
		}, {
			name:        	"Flag -f num 10",
			params: 		Params{NumFields: 10},
			input:       	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			outputCaseS:	[]string{"", "", "", "", "", ""},
		}, {
			name:        	"Flag -s num 0",
			params: 		Params{NumChars: 0},
			input:       	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			outputCaseS:	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
		}, {
			name:        	"Flag -s num 1",
			params: 		Params{NumChars: 1},
			input:       	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			outputCaseS:	[]string{"e love music.", " love music.", "hey love music.", " love music of Kartik.", "e love music of Kartik.", "hanks."},
		}, {
			name:        	"Flag -s num 50",
			params: 		Params{NumChars: 50},
			input:       	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			outputCaseS:	[]string{"", "", "", "", "", ""},
		}, {
			name:        	"Flag -f num 0 + -s num 1",
			params: 		Params{NumFields:0, NumChars: 1},
			input:       	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			outputCaseS:	[]string{"e love music.", " love music.", "hey love music.", " love music of Kartik.", "e love music of Kartik.", "hanks."},
		}, {
			name:        	"Flag -f num 1 + -s num 3",
			params: 		Params{NumFields: 1, NumChars: 3},
			input:       	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			outputCaseS:	[]string{"e music.", "e music.", "e music.", "e music of Kartik.", "e music of Kartik.", ""},
		}, {
			name:        	"Flag -f num 3 + -s num 5",
			params: 		Params{NumFields: 3, NumChars: 5},
			input:       	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			outputCaseS:	[]string{"", "", "", "rtik.", "rtik.", ""},
		}, {
			name:        	"Flag -f num 3 + -s num 10",
			params: 		Params{NumFields: 3, NumChars: 10},
			input:       	[]string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			outputCaseS:	[]string{"", "", "", "", "", ""},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.params.NumFields > 0 && tc.params.NumChars == 0 {
				uniq := cutStrF(tc.input, tc.params)
				assert.Equal(t, uniq, tc.outputCaseS)
			} else if tc.params.NumChars > 0 && tc.params.NumChars == 0 {
				uniq := cutCharS(tc.input, tc.params)
				assert.Equal(t, uniq, tc.outputCaseS)
			}
			if tc.params.NumFields > 0 && tc.params.NumChars > 0 {
				newSlice := cutStrF(tc.input, tc.params)
				slice := cutCharS(newSlice, tc.params)
				assert.Equal(t, slice, tc.outputCaseS)
			}
		})
	}
}

func TestUniqueCases(t *testing.T) {
	testCases := []struct {
		name        		string
		params		 		Params
		inputCase			[]string
		output				string
		input				string
	} {
		{
			name:        	"Flag -d success",
			params: 		Params{D: true},
			input:       	"We love music.\nI love music.\nThey love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.",
			output:			"",
		}, {
			name:        	"Flag -d",
			params: 		Params{D: true},
			input:       	"I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.",
			output:			"I love music.\nI love music of Kartik.\n",
		}, {
			name:        	"Flag -c",
			params: 		Params{C: true},
			input:       	"I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.",
			output:			"3 I love music.\n1 \n2 I love music of Kartik.\n1 Thanks.\n1 I love music of Kartik.\n",
		}, {
			name:        	"Flag -u",
			params: 		Params{U: true},
			input:       	"I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\nI love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.",
			output:			"\nThanks.\n",
		}, {
			name:        	"Flag -i",
			params: 		Params{I: true},
			input:       	"I LOVE MUSIC.\nI love music.\nI LoVe MuSiC.\n\nI love MuSIC of Kartik.\nI love music of kartik.\nThanks.\nI love music of kartik.\nI love MuSIC of Kartik.",
			output:			"I LOVE MUSIC.\n\nI love MuSIC of Kartik.\nThanks.\nI love music of kartik.\n",
		}, {
			name:        	"Flag -f num 1",
			params: 		Params{NumFields: 1},
			input:       	"We love music.\nI love music.\nThey love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.\n",
			output:			"We love music.\n\nI love music of Kartik.\nThanks.\n",
		}, {
			name:        	"Flag -s num 1",
			params: 		Params{NumChars: 1},
			input:       	"I love music.\nA love music.\nC love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.\n",
			output:			"I love music.\n\nI love music of Kartik.\nWe love music of Kartik.\nThanks.\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uniq, _ := startFlags(tc.input, tc.params)
			assert.Equal(t, uniq, tc.output)
		})
	}
}