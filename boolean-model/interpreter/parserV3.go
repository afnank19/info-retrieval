package interpreter

import (
	"boolean-model/sliceutil"
	"fmt"
)

func ParseQuery(query string, index map[string]string, fileList []string) []string{
	tokenizedQuery := tokenizeByDelimiter(query, ' ')
	hasAnd, hasOr := setOperations(tokenizedQuery)

	foundError := checkSyntax(tokenizedQuery)
	if foundError {
		return []string{}
	}

	var result [][]string

	result = handleNOT(tokenizedQuery, fileList, result, index)
	fmt.Println("POST NOT PARSING -> ", result)

	if hasAnd {
		result = handleAND(result)
		fmt.Println("POST AND PARSING -> ", result)
	}

	if hasOr {
		result = handleOR(result)
		fmt.Println("POST OR PARSING -> ", result)
	}

	fmt.Println("Pre cleaned result: ", result)
	finalRes := cleanUpResult(result)

	fmt.Println("<-----END----->")
	return finalRes
}

func setOperations(tokenizedQuery []string) (bool, bool) {
	var hasAnd, hasOr bool = false, false

	for i := 0; i < len(tokenizedQuery); i++ {
		if tokenizedQuery[i] == "AND" {
			hasAnd = true
		}
		if tokenizedQuery[i] == "OR" {
			hasOr = true
		}
	}

	return hasAnd, hasOr
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

// I always find a case where my code doesn't work
// but this has given reliable answers
func handleAND(result [][]string) [][]string{
	var newResult [][]string

	for i := 0; i < len(result); i++ {
		if result[i][0] == "AND" && result[i-1][0] == "X" {
			fmt.Println("Solved an AND, found a new one", newResult, len(newResult))
			documents := sliceutil.Intersection(newResult[len(newResult) - 1], result[i+1])
			if len(documents) == 0 {
				// This is if the result is empty, we put E in that place, \
				// so when parsing OR, we don't run into index out of bounds errors
				documents = append(documents, "E") 
			}

			result[i+1] = []string{"X"}

			newResult[len(newResult) - 1] = documents
			i++
		}

		if result[i][0] == "AND" && result[i-1][0] != "X" {
			fmt.Println(result[i-1], " AND ", result[i+1])
			documents := sliceutil.Intersection(result[i-1], result[i+1])
			if len(documents) == 0 {
				documents = append(documents, "E")
			}

			result[i-1] = []string{"X"}
			result[i+1] = []string{"X"}

			newResult = append(newResult, documents)
			i++
		}

		if result[i][0] == "OR" {
			var operator []string

			operator = append(operator, result[i][0])
			newResult = append(newResult, operator)
		}
		fmt.Println("EVAL STATE: ", result)
	}

	fmt.Println("NEW RESULT STATE", newResult)
	// OR fix i guess :P
	for j := 0; j < len(result); j++ {
		// Not a fan of this condition
		if result[j][0] != "X" && result[min(j + 1, len(result) - 1)][0] == "OR" && result[j][0] != "AND" && result[j][0] != "NOT" {
			newResult = append([][]string{result[j]}, newResult...)
			continue
		} 

		if result[j][0] != "X" && result[j - 1][0] == "OR" && result[j][0] != "AND" && result[j][0] != "NOT" {
			newResult = append(newResult, result[j])
			continue
		} 
	}

	return newResult
}

func handleOR(result [][]string) [][]string {
	var newResult [][]string

	for i := 0; i < len(result); i++ {
		if result[i][0] == "OR" && result[i-1][0] == "X" {
			fmt.Println("Solved an OR, found a new one", newResult, len(newResult))

			documents := sliceutil.Union(newResult[len(newResult) - 1], result[i+1])
			result[i+1] = []string{"X"}

			newResult[len(newResult) - 1] = documents
			i++
		}

		if result[i][0] == "OR" && result[i-1][0] != "X" {
			fmt.Println(result[i-1], " OR ", result[i+1])

			var documents []string
			if result[i-1][0] == "E" && result[i+1][0] == "E" {
				documents = sliceutil.Union([]string{}, []string{})
			} else if result[i-1][0] == "E" {
				documents = sliceutil.Union([]string{}, result[i+1])
			}  else if result[i+1][0] == "E" {
				documents = sliceutil.Union(result[i-1], []string{})
			} else {
				documents = sliceutil.Union(result[i-1], result[i+1])
			}

			result[i-1] = []string{"X"}
			result[i+1] = []string{"X"}

			newResult = append(newResult, documents)
			i++
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

func cleanUpResult(result [][]string) []string{
	var newRes []string

	temp := result[0]

	if temp[0] == "E" {
		temp = []string{}
	}

	newRes = temp

	return newRes
}