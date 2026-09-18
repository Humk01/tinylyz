package bwt

import (
	"bytes"
	"testing"
)

// Inverse("nnbaaa", 3) must produce exactly "banana".
func TestInverseBanana(t *testing.T) {
	got := Inverse([]byte("nnbaaa"), 3)

	want := []byte("banana")
	if !bytes.Equal(got, want) {
		t.Fatalf("Inverse(nnbaaa, 3) output\nwant: %q\ngot:  %q", want, got)
	}
}

func TestInverseAAB(t *testing.T) {
	bwt, idx := Forward([]byte("aab"))

	if !bytes.Equal(bwt, []byte("baa")) {
		t.Fatalf("Forward(aab): want %q, got %q", []byte("baa"), bwt)
	}
	if idx != 0 {
		t.Fatalf("Forward(aab) primary index: want 0, got %d", idx)
	}

	got := Inverse(bwt, idx)
	want := []byte("aab")
	if !bytes.Equal(got, want) {
		t.Fatalf("Inverse(%q, %d)\nwant: %q\ngot:  %q", bwt, idx, want, got)
	}
}

// Inverse("pssmipissii", 4) must produce exactly "mississippi".
func TestInverseMississippi(t *testing.T) {
	got := Inverse([]byte("pssmipissii"), 4)

	want := []byte("mississippi")
	if !bytes.Equal(got, want) {
		t.Fatalf("Inverse(pssmipissii, 4) output\nwant: %q\ngot:  %q", want, got)
	}
}

// Output length must always equal input length.
func TestInverseLength(t *testing.T) {
	cases := []struct {
		in  []byte
		idx int
	}{
		{[]byte{}, 0},
		{[]byte{'a'}, 0},
		{[]byte("nnbaaa"), 3},
		{[]byte("pssmipissii"), 4},
	}
	for _, c := range cases {
		got := Inverse(c.in, c.idx)
		if len(got) != len(c.in) {
			t.Fatalf("length mismatch for %q: input %d, output %d",
				c.in, len(c.in), len(got))
		}
	}
}

// Forward followed by Inverse must return the original input.
// This is the strongest test: any bug in either side shows up here.
func TestInverseRoundTrip(t *testing.T) {
	for _, in := range [][]byte{
		{},
		{'a'},
		{'x'},
		[]byte("aa"),
		[]byte("ab"),
		[]byte("ba"),
		[]byte("aaaa"),
		[]byte("banana"),
		[]byte("mississippi"),
		[]byte("abcdefghij"),
		[]byte("the quick brown fox"),
	} {
		bwt, idx := Forward(in)
		got := Inverse(bwt, idx)
		if !bytes.Equal(got, in) {
			t.Fatalf("round-trip failed for %q\nbwt: %q\nidx: %d\ngot: %q",
				in, bwt, idx, got)
		}
	}
}

// Empty input must not panic and must return empty.
func TestInverseEmpty(t *testing.T) {
	got := Inverse([]byte{}, 0)
	if len(got) != 0 {
		t.Fatalf("Inverse(empty): want empty, got %q", got)
	}
}

// Single byte must produce that same byte.
func TestInverseSingleByte(t *testing.T) {
	got := Inverse([]byte{'x'}, 0)
	if !bytes.Equal(got, []byte{'x'}) {
		t.Fatalf("Inverse(x): want %q, got %q", []byte{'x'}, got)
	}
}

// All-same input is the classic tie-breaking case.
func TestInverseAllSame(t *testing.T) {
	in := []byte("aaaa")
	got := Inverse(in, 0)
	if !bytes.Equal(got, in) {
		t.Fatalf("Inverse(aaaa): want %q, got %q", in, got)
	}
}