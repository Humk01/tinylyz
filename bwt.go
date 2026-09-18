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
// reverses the bwt transform
func Inverse(word []byte, num int) []byte {
    length := len(word)
    sorted := slices.Clone(word)
    slices.Sort(sorted)
    final := make([]byte, length)

    // start[c] = first row in F where c appears
    start := make(map[byte]int)
    for i := range length {
        c := sorted[i]
        if _, ok := start[c]; !ok {
            start[c] = i
        }
    }

    // rank[i] = which occurrence of word[i] this row is in L,
    // counting rows from the top.
    rank := make([]int, length)
    count := make(map[byte]int)
    for i := range length {
        c := word[i]
        count[c]++
        rank[i] = count[c]
    }

    index := num
    for i := range length {
        current := word[index]
        final[length-1-i] = current
        index = start[current] + rank[index] - 1
    }

    return final
}