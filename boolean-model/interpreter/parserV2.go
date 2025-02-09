// Deprecated code, NOT USED ANYWHERE
package interpreter

import (
	"boolean-model/sliceutil"
	"fmt"
)

// Parses without precendence
// Add NOT as well,
func queryIndex(query string, index map[string]string) {
	tokenizedQuery := tokenizeByDelimiter(query, ' ')

	var res []string
	var b []string
	var currentOp string
	for _, token := range tokenizedQuery {
		fmt.Print("CurrToken: " + token + "\n")
		if token != "AND" && token != "OR" {
			docStr := index[token]
			documents := tokenizeByDelimiter(docStr, ',')
			if currentOp == "" {
				res = documents
			} else {
				b = documents
			}
		} else { // maybe check if res is not empty
			currentOp = token
		}

		if currentOp != "" && len(res) > 0 && len(b) > 0 {
			if currentOp == "AND" {
				fmt.Print(res)
				fmt.Print(" INTERSECTION ")
				fmt.Print(b)
				fmt.Println()
				res = sliceutil.Intersection(res, b)

				currentOp = ""
				b = []string{}
			}

			if currentOp == "OR" {
				fmt.Print(res)
				fmt.Print(" UNION ")
				fmt.Print(b)
				fmt.Println()
				res = sliceutil.Union(res, b)

				currentOp = ""
				b = []string{}
			}
		}
	}

	fmt.Println(res)
}

// Based on precedence, although i need to test more cases
// There are some issues, especially with the order
func queryIndexV2(query string, index map[string]string, fileList []string) {
	tokenizedQuery := tokenizeByDelimiter(query, ' ')

	var res []string
	var b []string
	const EVAL_FLAG = "X" // Evaluated flag

	if len(tokenizedQuery) == 1 {
		// Rather than making the parsing func more complex
		// we can do a simple parse instead on a single term
		docStr := index[tokenizedQuery[0]]
		documents := tokenizeByDelimiter(docStr, ',')
		res = documents

		fmt.Println("<-----END-SIMPLE-PARSE----->")
		fmt.Println(res) // Will soon return this instead of a log
	}

	var operators = []string{"NOT", "AND", "OR"} // Sorted by precedence
	for i := 0; i < len(operators); i++ {
		// Once for each AND, OR
		for j := 0; j < len(tokenizedQuery); j++ {
			if tokenizedQuery[j] == operators[i] {
				fmt.Println("EVAL STATE: ", tokenizedQuery)

				// This logic handles NOT at the beginning only
				if j == 0 && tokenizedQuery[j] == "NOT" {
					docStr := index[tokenizedQuery[j+1]]
					tokenizedQuery[min(j+1, len(tokenizedQuery))] = EVAL_FLAG
					documents := tokenizeByDelimiter(docStr, ',')

					res = documents

					res, b = executeOperations(operators, res, b, fileList, i, "NOT")
					continue // Probably needs to execute operations before this
				}

				// This handles AND NOT
				if j != 0 && tokenizedQuery[j] == "NOT" && tokenizedQuery[j-1] == "AND" {
					res, b = parseAndNot(tokenizedQuery, res, b, j, index, EVAL_FLAG)
					res, b = executeOperations(operators, res, b, fileList, i, "AND_NOT")

					continue
				}

				// This handles OR NOT
				
				if tokenizedQuery[j] == "OR" && tokenizedQuery[j+1] == "NOT" {
					res, b = parseOrNot(tokenizedQuery, res, b, j, index, EVAL_FLAG)
					res, b = executeOperations(operators, res, b, fileList, i,"OR_NOT")

					continue
				}

				// I believe this can handle AND and OR
				if tokenizedQuery[j] == "AND" {
					leftToken := tokenizedQuery[j-1]
					rightToken := tokenizedQuery[min(j+1, len(tokenizedQuery))]
					var docStr string
	
					if len(res) == 0 {
						docStr := index[leftToken]
						tokenizedQuery[j-1] = EVAL_FLAG // Should change it to some other flag
						documents := tokenizeByDelimiter(docStr, ',')
						res = documents
					}
	
					if tokenizedQuery[j-1] == EVAL_FLAG {
						docStr = index[rightToken]
						tokenizedQuery[min(j+1, len(tokenizedQuery))] = EVAL_FLAG
					} else if tokenizedQuery[min(j+1, len(tokenizedQuery))] == EVAL_FLAG {
						docStr = index[leftToken]
						tokenizedQuery[min(j-1, len(tokenizedQuery))] = EVAL_FLAG
					} else {
						fmt.Println("Unhandled", tokenizedQuery[j])
					}
					documents := tokenizeByDelimiter(docStr, ',')
					b = documents
					res, b = executeOperations(operators, res, b, fileList, i)

					continue
				}

				if tokenizedQuery[j] == "OR" {
					leftToken := tokenizedQuery[j-1]
					rightToken := tokenizedQuery[min(j+1, len(tokenizedQuery))]
					var docStr string
	
					if len(res) == 0 {
						docStr := index[leftToken]
						tokenizedQuery[j-1] = EVAL_FLAG // Should change it to some other flag
						documents := tokenizeByDelimiter(docStr, ',')
						res = documents
					}
	
					if tokenizedQuery[j-1] == EVAL_FLAG {
						docStr = index[rightToken]
						tokenizedQuery[min(j+1, len(tokenizedQuery))] = EVAL_FLAG
					} else if tokenizedQuery[min(j+1, len(tokenizedQuery))] == EVAL_FLAG {
						docStr = index[leftToken]
						tokenizedQuery[min(j-1, len(tokenizedQuery))] = EVAL_FLAG
					} else {
						fmt.Println("Unhandled")
					}
					documents := tokenizeByDelimiter(docStr, ',')
					b = documents
					res, b = executeOperations(operators, res, b, fileList, i, "OR")

					continue
				}
			}
		}
	}

	fmt.Println("<-----END----->")
	fmt.Println(res)
}

func parseOrNot(tokenizedQuery []string, res, b []string, j int, index map[string]string, EVAL_FLAG string) ([]string, []string){
	// fmt.Println(len(tokenizedQuery), j)
	leftToken := tokenizedQuery[j-1]
	rightToken := tokenizedQuery[min(j+2, len(tokenizedQuery))]
	var docStr string

	// fmt.Println(leftToken, rightToken)

	if len(res) == 0 {
		docStr := index[leftToken]
		tokenizedQuery[j-1] = EVAL_FLAG // Should change it to some other flag
		documents := tokenizeByDelimiter(docStr, ',')
		res = documents
	}

	if tokenizedQuery[j-1] == EVAL_FLAG {
		docStr = index[rightToken]
		tokenizedQuery[min(j+2, len(tokenizedQuery))] = EVAL_FLAG
	} else if tokenizedQuery[min(j+2, len(tokenizedQuery))] == EVAL_FLAG {
		docStr = index[leftToken]
		tokenizedQuery[min(j-1, len(tokenizedQuery))] = EVAL_FLAG
	}
	tokenizedQuery[j] = "OP_EXEC" // Mark the AND as executed so it doesn't get run when we pass through for AND
	documents := tokenizeByDelimiter(docStr, ',')
	b = documents

	// fmt.Print("EVALUATION STATUS: ", tokenizedQuery, "\n")
	return res, b
}

func parseAndNot(tokenizedQuery []string, res, b []string, j int, index map[string]string, EVAL_FLAG string) ([]string, []string){
	leftToken := tokenizedQuery[j-2]
	rightToken := tokenizedQuery[min(j+1, len(tokenizedQuery))]
	var docStr string

	// fmt.Println(leftToken, rightToken)

	if len(res) == 0 {
		docStr := index[leftToken]
		tokenizedQuery[j-2] = EVAL_FLAG // Should change it to some other flag
		documents := tokenizeByDelimiter(docStr, ',')
		res = documents
	}

	if tokenizedQuery[j-2] == EVAL_FLAG {
		docStr = index[rightToken]
		tokenizedQuery[min(j+1, len(tokenizedQuery))] = EVAL_FLAG
	} else if tokenizedQuery[min(j+1, len(tokenizedQuery))] == EVAL_FLAG {
		docStr = index[leftToken]
		tokenizedQuery[min(j-2, len(tokenizedQuery))] = EVAL_FLAG
	}
	tokenizedQuery[j-1] = "OP_EXEC" // Mark the AND as executed so it doesn't get run when we pass through for AND
	documents := tokenizeByDelimiter(docStr, ',')
	b = documents

	// fmt.Print("EVALUATION STATUS: ", tokenizedQuery, "\n")
	return res, b
}

// Slices are being passed by value
func executeOperations(operators, res, b, fileList []string, i int, customOperation ...string) ([]string, []string) {
	if len(res) > 0 && len(b) == 0 && len(customOperation) > 0 && customOperation[0] == "NOT" {
		fmt.Println(res, " //  ", b, " //curr-op: ", operators[i])
		if operators[i] == "NOT" {
			fmt.Print("RUNNING -> ", fileList, " DIFFERENCE ", res, "\n")
			res = sliceutil.Difference(fileList, res)

			return res, b
		}
	}

	// Not sure if this condition is required now that i changed the parsing structure a little bit
	if len(res) > 0 && len(b) > 0 || true {
		if operators[i] == "NOT" && len(customOperation) > 0 && customOperation[0] == "AND_NOT" {
			fmt.Print("RUNNING -> ", res, " DIFFERENCE ", b, "\n")
			res = sliceutil.Difference(res, b)

			b = []string{}

			return res, b
		}

		if operators[i] == "AND" {
			fmt.Print("RUNNING -> ", res, " INTERSECTION ", b, "\n")
			res = sliceutil.Intersection(res, b)

			b = []string{}

			return res, b
		}

		if operators[i] == "OR" && len(customOperation) > 0 && customOperation[0] == "OR" {
			fmt.Print("RUNNING -> ", res, " UNION ", b, "\n")
			res = sliceutil.Union(res, b)

			b = []string{}

			return res, b
		}

		if operators[i] == "OR" && len(customOperation) > 0 && customOperation[0] == "OR_NOT" {
			fmt.Print("RUNNING -> ", res, " OR NOT ", b, "\n")
		
			b = sliceutil.Difference(fileList, b)
			res = sliceutil.Union(res, b)

			b = []string{}

			return res, b
		}
	}

	// Not really sure about this
	return res, b
}
