package main

import (
	"boolean-model/sliceutil"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const DOC_PATH = "./docs/"
	const TEST_DOC = "doc1.txt"
	fileList := readFilesFromDir(DOC_PATH)
	index := make(map[string]string)
	createIndex(fileList, DOC_PATH, index)
	fmt.Println(index)

	// tokens := tokenizer(readFile(TEST_DOC, DOC_PATH))

	// fmt.Println("TOKENS ->", tokens)
	// fmt.Println(tokenizeByDelimiter("doct1.txt,doc2.txt,", ','))

	fmt.Print("Query: ")
	reader := bufio.NewReader(os.Stdin)
	query, _ := reader.ReadString('\n')

	// queryIndex(query, index)
	fmt.Println("<---------->")
	queryIndexV2(query, index, fileList)
}

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
			// fmt.Println("Curr Token " + tokenizedQuery[j])
			if tokenizedQuery[j] == operators[i] {
				if j == 0 && tokenizedQuery[j] == "NOT" {
					// fmt.Println("UNIMPLEMENTED: Handle the first NOT")
					docStr := index[tokenizedQuery[j+1]]
					tokenizedQuery[min(j+1, len(tokenizedQuery))] = EVAL_FLAG
					documents := tokenizeByDelimiter(docStr, ',')

					res = documents

					res, b = executeOperations(operators, res, b, fileList, i)
					continue // Probably needs to execute operations before this
				}

				if j != 0 && tokenizedQuery[j] == "NOT" {
					fmt.Println("UNIMPLEMENTED: Handle NOT found in the middle")
					continue
				}

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
				}
				documents := tokenizeByDelimiter(docStr, ',')
				b = documents
				res, b = executeOperations(operators, res, b, fileList, i)
			}
		}
	}

	fmt.Println("<-----END----->")
	fmt.Println(res)
}

// Slices are being passed by value
func executeOperations(operators, res, b, fileList []string, i int) ([]string, []string) {
	if len(res) > 0 && len(b) == 0 {
		fmt.Println(res, " //  ", b, " //curr-op: ", operators[i])
		if operators[i] == "NOT" {
			fmt.Print("RUNNING -> ", fileList, " DIFFERENCE ", res, "\n")
			res = sliceutil.Difference(fileList, res)

			return res, b
		}
	}

	// Not sure if this condition is required now that i changed the parsing structure a little bit
	if len(res) > 0 && len(b) > 0 || true {
		if operators[i] == "NOT" {
			fmt.Println("Should run NOT")
		}

		if operators[i] == "AND" {
			fmt.Print("RUNNING -> ", res, " INTERSECTION ", b, "\n")
			res = sliceutil.Intersection(res, b)

			b = []string{}

			return res, b
		}

		if operators[i] == "OR" {
			fmt.Print("RUNNING -> ", res, " UNION ", b, "\n")
			res = sliceutil.Union(res, b)

			b = []string{}

			return res, b
		}
	}

	// Not really sure about this
	return res, b
}

func createIndex(fileList []string, basePath string, index map[string]string) {
	for _, filename := range fileList {
		fileStr := readFile(filename, basePath)
		fileStr = strings.ToLower(fileStr)

		tokens := tokenizer(fileStr)

		addTokensToIndex(tokens, filename, index)
	}
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

func readFilesFromDir(path string) []string {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var fileList []string
	for _, file := range dirEntries {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}

	return fileList
}

func tokenizer(str string) []string {
	var tokens []string
	var currChars []byte

	for i := 0; i < len(str); i++ {
		// 32 is ASCII for " " & 10 is for \n
		if str[i] == ' ' || str[i] == '\n' {
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

// Pass the delimiter as a byte
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

func readFile(filename, path string) string {
	data, err := os.ReadFile(path + filename)
	if err != nil {
		panic(err)
	}

	fileData := string(data)
	fileData = fileData + "\n" // Adding our own terminator to know if the string has ended

	return fileData
}
