package bwt

import (
	"bytes"
	"testing"
)

// Forward("banana") must produce exactly "nnbaaa" with primary index 3.
func TestForwardBanana(t *testing.T) {
	got, idx := Forward([]byte("banana"))

	want := []byte("nnbaaa")
	if !bytes.Equal(got, want) {
		t.Fatalf("Forward(banana) output\nwant: %q\ngot:  %q", want, got)
	}
	if idx != 3 {
		t.Fatalf("Forward(banana) primary index: want 3, got %d", idx)
	}
}

// Output length must always equal input length.
func TestForwardLength(t *testing.T) {
	for _, in := range [][]byte{
		{},
		{'a'},
		[]byte("banana"),
		[]byte("mississippi"),
		[]byte("the quick brown fox"),
	} {
		got, _ := Forward(in)
		if len(got) != len(in) {
			t.Fatalf("length mismatch for %q: input %d, output %d",
				in, len(in), len(got))
		}
	}
}

// Output must be a permutation of the input: same bytes, same counts.
func TestForwardIsPermutation(t *testing.T) {
	for _, in := range [][]byte{
		{},
		{'a'},
		[]byte("banana"),
		[]byte("mississippi"),
		[]byte("aaaaaaaaaa"),
		[]byte("abcdefghij"),
	} {
		got, _ := Forward(in)

		countIn := [256]int{}
		countOut := [256]int{}
		for _, b := range in {
			countIn[b]++
		}
		for _, b := range got {
			countOut[b]++
		}
		if countIn != countOut {
			t.Fatalf("not a permutation for %q\nin:  %v\nout: %v",
				in, countIn, countOut)
		}
	}
}

// Primary index must be a valid row of the sorted table.
func TestForwardPrimaryIndexInRange(t *testing.T) {
	for _, in := range [][]byte{
		{},
		{'a'},
		[]byte("banana"),
		[]byte("mississippi"),
		[]byte("aaaaaaaaaa"),
	} {
		_, idx := Forward(in)
		if idx < 0 || idx >= len(in) {
			// Empty input has length 0, so idx must be 0 and the
			// check below must allow it. Handle empty as a special case.
			if len(in) == 0 && idx == 0 {
				continue
			}
			t.Fatalf("primary index out of range for %q: %d", in, idx)
		}
	}
}

// Empty input must not panic.
func TestForwardEmpty(t *testing.T) {
	got, idx := Forward([]byte{})
	if len(got) != 0 {
		t.Fatalf("Forward(empty) output: want empty, got %q", got)
	}
	if idx != 0 {
		t.Fatalf("Forward(empty) primary index: want 0, got %d", idx)
	}
}

// Single byte must produce that same byte.
func TestForwardSingleByte(t *testing.T) {
	got, idx := Forward([]byte{'x'})
	if !bytes.Equal(got, []byte{'x'}) {
		t.Fatalf("Forward(x): want %q, got %q", []byte{'x'}, got)
	}
	if idx != 0 {
		t.Fatalf("Forward(x) primary index: want 0, got %d", idx)
	}
}

// All-same input is the classic tie-breaking case.
func TestForwardAllSame(t *testing.T) {
	in := []byte("aaaa")
	got, _ := Forward(in)
	if !bytes.Equal(got, in) {
		t.Fatalf("Forward(aaaa): want %q, got %q", in, got)
	}
}

// A known non-trivial case.
func TestForwardMississippi(t *testing.T) {
	got, _ := Forward([]byte("mississippi"))
	// last column of the sorted rotations of "mississippi"
	want := []byte("pssmipissii")
	if !bytes.Equal(got, want) {
		t.Fatalf("Forward(mississippi)\nwant: %q\ngot:  %q", want, got)
	}
}