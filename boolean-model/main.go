package main

import (
	"boolean-model/indexer"
	"boolean-model/interpreter"
	"bufio"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var queryStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ffffff"))

func main() {
	const DOC_PATH = "./docs/"
	const TEST_DOC = "doc1.txt"
	fileList := readFilesFromDir(DOC_PATH)
	index := make(map[string]string)

	indexer.CreateIndex(fileList, DOC_PATH, index)
	// fmt.Println(index)

	fmt.Print(queryStyle.Render("Query:"), " ")
	reader := bufio.NewReader(os.Stdin)
	query, _ := reader.ReadString('\n')

	result := interpreter.ParseQuery(query, index, fileList)
	printResult(result)
}

func printResult(result []string) {
	var style = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
    PaddingRight(2)

	var resultStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#3C3C3C"))

	fmt.Print("\n")
	fmt.Println(style.Render(" QUERY RESULT:"), resultStyle.Render(result...))
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
