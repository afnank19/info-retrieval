package interpreter

import (
	"boolean-model/sliceutil"
	"fmt"
)

func ParseQuery(query string, index map[string]string, fileList []string) {
	// First pass, convert all terms to their respective index values
	// Also in the first pass, we will solve NOT aswell.
	// Second pass, we solve all ANDs, rebuild the array again
	// finally, solve all ORs, if any

	tokenizedQuery := tokenizeByDelimiter(query, ' ')

	var result [][]string

	result = handleNOT(tokenizedQuery, fileList, result, index)
	fmt.Println("POST NOT PARSING -> ", result)
	result = handleAND(fileList, result, index)
	fmt.Println("POST AND PARSING -> ", result)

}

func handleNOT(tokenizedQuery, fileList []string, result [][]string, index map[string]string) [][]string {
	for j := 0; j < len(tokenizedQuery); j++ {
		if tokenizedQuery[j] == "AND" || tokenizedQuery[j] == "OR" {
			var operator []string

			operator = append(operator, tokenizedQuery[j])
			result = append(result, operator)

			continue
		}

		if tokenizedQuery[j] == "NOT" {
			docStr := index[tokenizedQuery[j+1]]
			documents := tokenizeByDelimiter(docStr, ',')
			// fmt.Println("Prev docs b4 NOT: ", documents)
			documents = sliceutil.Difference(fileList, documents) // Difference from the universal set

			result = append(result, documents)
			j++ // we skip the next token, because we already used that token i.e NOT [TERM]

			continue
		}

		docStr := index[tokenizedQuery[j]]
		documents := tokenizeByDelimiter(docStr, ',')

		result = append(result, documents)
	}

	return result
}

func handleAND(fileList []string, result [][]string, index map[string]string) [][]string{
	var newResult [][]string

	for i := 0; i < len(result); i++ {
		// i-1 could have a flag -> probably need something from newResult
		// i-1 could not have a flag -> use the current documents
		if result[i][0] == "AND" {
			fmt.Println(result[i-1], " AND ", result[i+1])
			documents := sliceutil.Intersection(result[i-1], result[i+1])

			newResult = append(newResult, documents)
		}

		if result[i][0] == "OR" {
			var operator []string

			operator = append(operator, result[i][0])
			newResult = append(newResult, operator)
		}
	}

	return newResult
}

func tokenizeByDelimiter(str string, delimiter byte) []string {
	var tokens []string
	var currChars []byte
	str = str + "\n"

	for i := 0; i < len(str); i++ {
		if str[i] == delimiter || str[i] == '\n' {
			if len(currChars) > 0 {
				tokens = append(tokens, string(currChars))
				currChars = []byte{}
			}
			continue
		}
		currChars = append(currChars, str[i])
	}

	return tokens
}