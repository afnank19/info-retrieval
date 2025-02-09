package indexer

import (
	"os"
	"strings"
)

var STOPWORDS = []string{"a", "and", "the", "not", "or", "of", "is", "are", "at", "this", "if", "in", "on"}

func CreateIndex(fileList []string, basePath string, index map[string]string) {
	for _, filename := range fileList {
		fileStr := readFile(filename, basePath)
		fileStr = strings.ToLower(fileStr)

		tokens := tokenizer(fileStr)

		addTokensToIndex(tokens, filename, index)
	}
}

func readFile(filename, path string) string {
	data, err := os.ReadFile(path + filename)
	if err != nil {
		panic(err)
	}

	fileData := string(data)
	fileData = fileData + "\n" // Adding our own terminator to know if the string has ended

	return fileData
}

func tokenizer(str string) []string {
	var tokens []string
	var currChars []byte

	for i := 0; i < len(str); i++ {
		// 32 is ASCII for " " & 10 is for \n
		if str[i] == ' ' || str[i] == '\n' {
			if len(currChars) > 0 {
				currWord := string(currChars)
				if isStopword(STOPWORDS, currWord) {
					continue
				}
				tokens = append(tokens, string(currChars))
				currChars = []byte{}
			}
			continue
		}
		currChars = append(currChars, str[i])
	}

	return tokens
}

func isStopword(STOPWORDS []string, element string) bool {
	for _, stopword := range STOPWORDS {
		if element == stopword {
			return true
		}
	}

	return false
}

func addTokensToIndex(tokens []string, filename string, index map[string]string) {
	for _, token := range tokens {
		// If token already exist, don't do anything, go to next token
		currFilenames := index[token]
		if strings.Contains(currFilenames, filename) {
			continue
		}
		index[token] = currFilenames + filename + ","
	}
}