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
	// This has to be tokenized as well

	queryIndex(query, index)

	var testSet1 = []string{"cat", "hat", "bat", "kit"}
	var testSet2 = []string{"sat", "umaiza"}

	res := sliceutil.Union(testSet2, testSet1)
	fmt.Println(res)
}

func queryIndex(query string, index map[string]string) {
	// TODOS:
	// 1. Tokenize the query string (Not doing operators now)
	// 2. Query the index, get the document string back
	// 3. Tokenize the document string, add to an array
	// 4. Have an answer array, that adds or removes documents

	// fmt.Println(tokenizeByDelimiter(query, ' '))

	tokenizedQuery := tokenizeByDelimiter(query, ' ')

	for _, token := range tokenizedQuery {
		if token != "AND" {

		}
	}

	docStr := index[query]

	resultingDocs := tokenizeByDelimiter(docStr, ',')

	fmt.Println(resultingDocs)
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
