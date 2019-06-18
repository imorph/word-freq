package freq

import (
	"sort"
	"strings"
)

// CleanupWord takes word with "." "," ";" ":" "\n" ")" "(" and return "clean" word
func cleanupWord(inWord string) string {
	outWord := strings.TrimSuffix(inWord, `
`)
	outWord = strings.TrimSuffix(outWord, ")")
	outWord = strings.TrimPrefix(outWord, "(")
	outWord = strings.TrimSuffix(outWord, ".")
	outWord = strings.TrimSuffix(outWord, ",")
	outWord = strings.TrimSuffix(outWord, ";")
	outWord = strings.TrimSuffix(outWord, ":")
	outWord = strings.TrimSuffix(outWord, "..")
	outWord = strings.ToLower(outWord)
	return outWord
}

// StringToCleanSlice splits text to words (cleaned from any of . , : ; \n ) and retuns it as slice of words in lower case
func stringToCleanSlice(text string) []string {
	var cleanSlice []string

	// get slice of strings without \n
	dirtyStringSlice := strings.Split(text, `
`)

	// get slice of strings (words) without " "
	for _, dirtyString := range dirtyStringSlice {
		dirtyWordSlice := strings.Split(dirtyString, " ")
		for _, dirtyWord := range dirtyWordSlice {
			cleanSlice = append(cleanSlice, cleanupWord(dirtyWord))
		}
	}
	return cleanSlice
}

// StatMapFromSlice iterate over slice of words and generates map with words counts
func statMapFromSlice(words []string) map[string]int {
	statMap := make(map[string]int)
	for _, word := range words {
		wordCount, _ := statMap[word]
		wordCount++
		statMap[word] = wordCount
	}
	return statMap
}

// Top10 returns 10 most used words in form of slice
func Top10(text string) []string {

	var top10 []string

	type singleStat struct {
		word  string
		count int
	}

	myStatMap := []singleStat{}

	statMap := statMapFromSlice(stringToCleanSlice(text))
	for key, val := range statMap {
		myStatMap = append(myStatMap, singleStat{key, val})
	}
	// reverse sort
	sort.Slice(myStatMap, func(i, j int) bool { return myStatMap[i].count > myStatMap[j].count })

	for i, stat := range myStatMap {
		top10 = append(top10, stat.word)
		if i == 9 {
			return top10
		}
	}
	return top10
}
