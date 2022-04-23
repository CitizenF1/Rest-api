package substr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

func GetSubStr(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodyString := string(body)
	response, max := MaxSubstring(bodyString)
	responseString := "Max len " + strconv.Itoa(max) + " " + response

	fmt.Fprint(w, responseString)
}

func MaxSubstring(s string) (string, int) {
	vChars := make(map[rune]int)
	var start, maxLength int
	var str []rune

	for i, ch := range []rune(s) {
		if lastI, ok := vChars[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		vChars[ch] = i
	}

	for k, _ := range vChars {
		str = append(str, k)
	}

	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})

	return string(str), maxLength
}
