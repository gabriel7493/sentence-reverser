package main

import (
	"fmt"
	"os"
	"strings"
)

func ReverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func ReverseSentence(sentence string, stringReverser func (string) string) (result string) {
	if stringReverser == nil {
		return
	}
	words := strings.Split(stringReverser(sentence), " ")
	for i, v := range words {
		result += stringReverser(v)
		if i < len(words) - 1 {
			result += " "
		}
	}
	return
}

func sentenceProcessor() {
	args := os.Args
	length := len(args)
	if length < 2 {
		fmt.Println("You should provide a valid sentence to reverse")
		os.Exit(0)
	} else if length > 2 {
		fmt.Println("As this program just process first argument, it will be the only one to be processed...")
	}
	fmt.Println(ReverseSentence(args[1], ReverseString))
}

func main() {
	sentenceProcessor()
}
