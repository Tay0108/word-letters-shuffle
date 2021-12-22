package main

import "fmt"

// func shuffleWord(word string) []rune {
// 	wordInnerials := word[1 : len(word)-1]
// 	shuffledString := []rune(wordInnerials)
// 	rand.Shuffle(len(word), func(i, j int) { shuffledString[i], shuffledString[j] = shuffledString[j], shuffledString[i] })
// 	shuffledString = append(shuffledString, rune(word[len(word)-1]))
// 	// shuffledString.
// 	return shuffledString
// }

// TODO: dopisać testy
// TODO: wystawić jako serwis na Heroku

func main() {

	// inputText := ""
	// rand.Seed(time.Now().UnixNano())

	// inputTextSplitted := strings.Fields(inputText)
	// outputTextSplitted := make([]string, 0, 100)

	// for i := range inputTextSplitted {
	// 	shuffledWord := shuffleWord(inputTextSplitted[i])
	// 	outputTextSplitted = append(outputTextSplitted, string(shuffledWord))
	// }

	// fmt.Println(outputTextSplitted)

	testString := "literal"

	testString2 := testString[:1] + "d" + testString[2:]

	fmt.Println(testString2)
}
