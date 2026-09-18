package bwt

import (
	"bytes"
	"slices"
)

// creates the bwt
func Forward(word []byte) ([]byte, int) {
	primaryIndex := 0
	bwt_slice := [][]byte{}
	length := len(word)
	bwt := []byte{}

	for i := range length {
		temp := make([]byte, 0, length)
		temp = append(temp, word[i:]...)
		temp = append(temp, word[:i]...)
		bwt_slice = append(bwt_slice, temp)
	}

	slices.SortFunc(bwt_slice, bytes.Compare)

	for i, v := range bwt_slice {
		if bytes.Equal(bwt_slice[i], word) {
			primaryIndex = i
		}
		bwt = append(bwt, v[length-1])
	}
	return bwt, primaryIndex
}

// reverses the bwt transform
func Inverse(word []byte, num int) []byte {
    index := num
    length := len(word)
    sorted := slices.Clone(word)
    slices.Sort(sorted)
    final := make([]byte, length)
    start := make(map[byte]int)
    seen := make(map[byte]int)

    // build start once, from F
    for i := range length {
        c := sorted[i]
        if _, ok := start[c]; !ok {
            start[c] = i
        }
    }

    // walk, filling final from the back
    for i := range length {
        current := word[index]
        final[length-1-i] = current
        seen[current] += 1
        index = start[current] + seen[current] - 1
    }

    return final
}