package handlers

import (
	"strings"
)

// обработка массива слов, полученных от POST запроса
func postWordHandle(handler *HttpHandler, words []string) {
	for _, word := range words {
		wordFound := false
		for keyWord, wordsArray := range handler.Storage {
			if len(word) != len(keyWord) {
				continue
			}
			if matchedWord := isAnagram(keyWord, word); matchedWord {
				wordFound = true
				if hasDuplicate := checkDuplicate(wordsArray, word); !hasDuplicate {
					handler.Storage[keyWord] = append(handler.Storage[keyWord], word)
				}
			}
		}
		if !wordFound {
			handler.Storage[word] = []string{word}
		}
	}
}

// поиск слов по GET запросу
func getWordHandle(handler *HttpHandler, getWord string) []string {
	for keyWord := range handler.Storage {
		if len(getWord) != len(keyWord) {
			continue
		}
		if matchedWord := isAnagram(keyWord, getWord); matchedWord {
			return handler.Storage[keyWord]
		}
	}
	return []string{}
}

// проверка слов на анаграмму
func isAnagram(keyWord, word string) bool {
	for _, runeCode := range keyWord {
		char := strings.ToLower(string(runeCode))
		wordCount := strings.Count(strings.ToLower(word), char)
		keyCount := strings.Count(strings.ToLower(keyWord), char)
		if keyCount != wordCount {
			return false
		}
	}
	return true
}

// проверка на дубликат
func checkDuplicate(wordsArray []string, word string) bool {
	for _, value := range wordsArray {
		if word == value {
			return true
		}
	}
	return false
}
