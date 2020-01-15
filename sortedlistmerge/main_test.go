package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func TestTableSortListMerge(t *testing.T) {
	var tests = []struct {
		input    string
		data     string
		expected string
	}{
		{"<nil>", "<nil>", "<nil>\n"},
		{"<nil>", "2-> 2-> 4-> 9-> 12-> 12-> 19-> 20-> <nil>", "2 -> 2 -> 4 -> 9 -> 12 -> 12 -> 19 -> 20 -> <nil>\n"},
		{"4-> 4-> 6-> 9-> 13-> 18-> 20-> 20-> <nil>", "<nil>", "4 -> 4 -> 6 -> 9 -> 13 -> 18 -> 20 -> 20 -> <nil>\n"},
		{"0-> 2-> 11-> 30-> 54-> 56-> 70-> 79-> 99-> <nil>", "23-> 28-> 38-> 67-> 67-> 79-> 95-> 97-> <nil>", "0 -> 2 -> 11 -> 23 -> 28 -> 30 -> 38 -> 54 -> 56 -> 67 -> 67 -> 70 -> 79 -> 79 -> 95 -> 97 -> 99 -> <nil>\n"},
		{"0-> 7-> 39-> 92-> 97-> 93-> 91-> 28-> 64-> <nil>", "80-> 23-> 27-> 30-> 85-> 81-> 75-> 70-> <nil>", "0 -> 7 -> 23 -> 27 -> 28 -> 30 -> 39 -> 64 -> 70 -> 75 -> 80 -> 81 -> 85 -> 91 -> 92 -> 93 -> 97 -> <nil>\n"},
		{"0-> 3-> 8-> 8-> 13-> 19-> 34-> 38-> 46-> <nil>", "7-> 39-> 45-> 53-> 59-> 70-> 76-> 79-> <nil>", "0 -> 3 -> 7 -> 8 -> 8 -> 13 -> 19 -> 34 -> 38 -> 39 -> 45 -> 46 -> 53 -> 59 -> 70 -> 76 -> 79 -> <nil>\n"},
	}
	for j, test := range tests {
		var link1 *NodeI
		var link2 *NodeI
		link1 = ParceInput(test.input, link1)
		link2 = ParceInput(test.data, link2)
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		PrintList(SortedListMerge(link1, link2))
		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		strout := ""
		for _, word := range out {
			strout = strout + string(word)
		}
		if strout != test.expected {
			t.Error("test failed:", "expected:", test.expected, "received:", strout, "number:", j+1)
		}

	}

}
func ParceInput(inp string, link *NodeI) *NodeI {
	strAtoi := ""
	var strArr []string
	check := false
	for _, word := range inp {
		check = false

		if word >= '0' && word <= '9' {
			strAtoi = strAtoi + string(word)
			check = true
		}
		if !check && strAtoi != "" {
			strArr = append(strArr, strAtoi)
			strAtoi = ""
			check = false
		}
	}
	number := 0
	for _, word := range strArr {
		number, _ = strconv.Atoi(string(word))
		link = listPushBack(link, number)
	}

	return link
}