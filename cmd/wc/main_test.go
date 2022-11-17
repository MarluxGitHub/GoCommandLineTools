package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n")

	exp := 3


	res := count(b, Flags{lines: false, bytes: false})

	if res != exp {
		t.Errorf("expected %d, got %d", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n word4")

	exp := 2

	res := count(b, Flags{lines: true, bytes: false})

	if res != exp {
		t.Errorf("expected %d, got %d", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n word4")

	exp := 24

	res := count(b,Flags{lines: false, bytes: true})

	if res != exp {
		t.Errorf("expected %d, got %d", exp, res)
	}
}