package bwt

import (
	"bytes"
	"slices"
)

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
