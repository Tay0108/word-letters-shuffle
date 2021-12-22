package word

import "math/rand"

func ShuffleWordInnerials(word string) string {
	if len(word) <= 3 {
		return word
	}
	runes := []rune(word)
	wordInnerials := runes[1 : len(runes)-1]
	rand.Shuffle(len(wordInnerials), func(i, j int) { wordInnerials[i], wordInnerials[j] = wordInnerials[j], wordInnerials[i] })
	shuffledRunes := append(wordInnerials, runes[len(runes)-1])
	shuffledRunes = append([]rune{runes[0]}, shuffledRunes...)

	return string(shuffledRunes)
}
