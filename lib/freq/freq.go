package freq

import (
	"strings"
	"sort"
)


// MapKeyByValue returns key by value (if it exists)
func MapKeyByValue(m map[string]int, value int) (key string, ok bool) {
	for k, v := range m {
	  if v == value { 
		key = k
		ok = true
		return key, ok
	  }
	}
	return key, ok
  }

// CleanupWord takes word with "." "," ";" ":" "\n" and return "clean" word
func CleanupWord(inWord string) string {
	outWord := strings.TrimSuffix(inWord, `
`)
	outWord = strings.TrimSuffix(outWord, ".")
	outWord = strings.TrimSuffix(outWord, ",")
	outWord = strings.TrimSuffix(outWord, ";")
	outWord = strings.TrimSuffix(outWord, ":")
	outWord = strings.TrimSuffix(outWord, "..")
    outWord = strings.ToLower(outWord)
	return outWord
}

// StringToCleanSlice splits text to words (cleaned from any of . , : ; \n ) and retuns it as slice of words in lower case
func StringToCleanSlice(text string) []string {
	var cleanSlice []string

	// get slice of strings without \n
	dirtyStringSlice := strings.Split(text, `
`)

    // get slice of strings (words) without " "
	for _, dirtyString := range dirtyStringSlice {
		dirtyWordSlice := strings.Split(dirtyString, " ")
		for _, dirtyWord := range dirtyWordSlice {
			cleanSlice = append(cleanSlice, CleanupWord(dirtyWord))
		}
	}
	
	return cleanSlice
}

// StatMapFromSlice iterate over slice of words and generates map with words counts
func StatMapFromSlice(words []string) map[string]int {
	statMap := make(map[string]int)
	for _, word := range words {
		wordCount, ok := statMap[word]
		if ok {
			wordCount++
			statMap[word] = wordCount
		} else {
			wordCount = 1
			statMap[word] = wordCount
		}
	}
    return statMap
}


// Top10 returns 10 most used words in form of slice
func Top10(statMap map[string]int) []string {
	var top10 []string

	type singleStat struct {
		word string
		count int
	}
	type myStat []singleStat

	var mySingleStat singleStat
	var myStatMap myStat

	for key, val := range statMap {
		mySingleStat.word = key
		mySingleStat.count = val
		myStatMap = append(myStatMap, mySingleStat)
	}
    // reverse sort
	sort.Slice(myStatMap, func(i, j int)bool {return myStatMap[i].count > myStatMap[j].count})

	for i, stat := range myStatMap {
		top10 = append(top10, stat.word)
		if i == 9 {
			return top10
		}
	}
	return top10
}