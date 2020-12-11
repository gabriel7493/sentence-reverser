package main

import "testing"

var TestCases = []struct {
	Str string
	WordExpected string
	SentenceExpected string
}{
	{"a", "a", "a"},
	{"word", "drow", "word"},
	{"word1 word2", "2drow 1drow", "word2 word1"},
	{"sentence, with. punctuation;", ";noitautcnup .htiw ,ecnetnes", "punctuation; with. sentence,"},
}

func TestReverseString(t *testing.T) {
	for _, v := range TestCases {
		t.Run(v.Str, func(t *testing.T) {
			get := ReverseString(v.Str)
			if get != v.WordExpected {
				t.Errorf("Expected: %v, got: %v", v.WordExpected, get)
			}
		})
	}
}

func TestReverseSentence(t *testing.T) {
	for _, v := range TestCases {
		t.Run(v.Str, func(t *testing.T) {
			get := ReverseSentence(v.Str, ReverseString)
			if get != v.SentenceExpected {
				t.Errorf("Expected: %v, got: %v", v.SentenceExpected, get)
			}
		})
	}
}

