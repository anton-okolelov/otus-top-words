package top_words

import (
	"regexp"
	"sort"
	"strings"
)

func GetTop10Words(text string) []string {
	delimiterRegex := regexp.MustCompile(`\P{L}+`)
	words := delimiterRegex.Split(text, -1)

	counter := map[string]int{}

	for _, word := range words {
		if word != "" {
			counter[word] = counter[word] + 1
		}
	}

	type wordcount struct {
		word  string
		count int
	}

	countedWords := make([]wordcount, len(counter))
	i := 0
	for word, cnt := range counter {
		countedWords[i] = wordcount{word, cnt}
		i += 1
	}

	sort.Slice(countedWords, func(i int, j int) bool {
		ci := countedWords[i]
		cj := countedWords[j]
		return ci.count > cj.count ||
			ci.count == cj.count && strings.Compare(ci.word, cj.word) > 0 // для стабильности результата
	})

	var result []string

	for i, value := range countedWords {
		result = append(result, value.word)
		if i >= 9 {
			break
		}
	}

	return result
}
