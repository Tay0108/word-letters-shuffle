package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"word-letters-shuffle/word"
)

// TODO: dopisać testy
// TODO: wystawić jako serwis na Heroku

func main() {

	rand.Seed(time.Now().UnixNano()) // to ensure unique shuffle each time it runs
	inputText := "According to a research at Cambridge University, it doesn't matter in what order the letters in a word are, the only important thing is that the first and last letter be at the right place. The rest can be a total mess and you can still read it without problem. This is because the human mind does not read every letter by itself but the word as a whole."

	inputTextSplitted := strings.Fields(inputText)
	outputTextSplitted := make([]string, 0, 100)

	for i := range inputTextSplitted {
		shuffledWord := word.ShuffleWordInnerials(inputTextSplitted[i])
		outputTextSplitted = append(outputTextSplitted, string(shuffledWord))
	}

	outputText := strings.Join(outputTextSplitted, " ")

	fmt.Println(outputText)

}
