package handlers

import (
	"strings"
)

func postWordHandle(handler *HttpHandler, words []string){
	for _, word := range words{
		found := false
		for keyWord, wordsArray := range handler.Storage {
			match := true
			if len(word) != len(keyWord){
				continue
			}
			for _, runeCode := range keyWord{
				char := strings.ToLower(string(runeCode))
				wordCount := strings.Count(strings.ToLower(word), char)
				keyCount := strings.Count(strings.ToLower(keyWord), char)
				if keyCount != wordCount{
					match = false
					break
				}
			}
			if match{
				found = true
				hasDuplicate := false
				for _, value := range wordsArray {
					if word == value {
						hasDuplicate = true
						break
					}
				}
				if !hasDuplicate {
					handler.Storage[keyWord] = append(handler.Storage[keyWord], word)

				}
			}
		}
		if !found{
			handler.Storage[word] = []string{word}
		}
	}
}

func getWordHandle(handler *HttpHandler, getWord string)[]string{
	for keyWord, _ := range handler.Storage {
		match := true
		if len(getWord) != len(keyWord){
			continue
		}
			for _, runeCode := range keyWord{
				char := strings.ToLower(string(runeCode))
				wordCount := strings.Count(strings.ToLower(getWord), char)
				keyCount := strings.Count(strings.ToLower(keyWord), char)
				if keyCount != wordCount{
					match = false
					break
				}
			}
			if match{
				return handler.Storage[keyWord]
			}
		}
	return []string{}
}
