package interpreter

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func checkSyntax(tokenizedQuery []string) bool{
	length := len(tokenizedQuery)
	flags := make([]string, length)

	var foundError bool = false
	var errorMsg string

	copy(flags, tokenizedQuery)

	for i := 0; i < length; i++ {
		if tokenizedQuery[i] == "AND" {
			prevTerm := tokenizedQuery[max(i-1, 0)]
			nextTerm := tokenizedQuery[min(i+1, length-1)]

			if prevTerm == "AND" || prevTerm == "OR" || prevTerm == "NOT" {
				flags[max(i-1, 0)] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Unexpected/absent token " + prevTerm + " \nExpected [TERM] AND [TERM]" 
				break
			}

			if nextTerm == "AND" || nextTerm == "OR" {
				flags[min(i+1, length-1)] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Unexpected/absent token " + nextTerm + " \nExpected [TERM] AND [TERM]" 
				break
			}
		}

		if tokenizedQuery[i] == "OR" {
			prevTerm := tokenizedQuery[max(i-1, 0)]
			nextTerm := tokenizedQuery[min(i+1, length-1)]

			if prevTerm == "AND" || prevTerm == "OR" || prevTerm == "NOT" {
				flags[max(i-1, 0)] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Unexpected/absent token " + prevTerm + " \nExpected [TERM] OR [TERM]" 
				break
			}

			if nextTerm == "AND" || nextTerm == "OR" {
				flags[min(i+1, length-1)] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Unexpected/absent token " + nextTerm + " \nExpected [TERM] OR [TERM]" 
				break
			}
		}

		if tokenizedQuery[i] == "NOT" {
			if i == 0  && i+1 == len(tokenizedQuery) {
				flags[i] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Expected NOT [TERM] , got nothing" 
				break
			}

			prevTerm := tokenizedQuery[max(i-1, 0)]
			nextTerm := tokenizedQuery[min(i+1, length-1)]

			if  prevTerm == "NOT" && i != 0 {
				flags[max(i-1, 0)] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Unexpected token " + prevTerm + " \nExpected [OPERATOR] NOT [TERM]" 
				break
			}

			if nextTerm == "AND" || nextTerm == "OR" || nextTerm == "NOT" {
				flags[min(i+1, length-1)] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Unexpected token " + nextTerm + " \nExpected NOT [TERM]" 
				break
			}
		}

		// Problem area, might need to look into it more
		currTerm := tokenizedQuery[i]
		if length > 1 && i != 0  && currTerm != "AND" && currTerm != "OR" && currTerm != "NOT" {
			prevTerm := tokenizedQuery[max(i-1, 0)]
			nextTerm := tokenizedQuery[min(i+1, length-1)]

			if prevTerm != "AND" && prevTerm != "OR" && prevTerm != "NOT" {
				flags[max(i-1, 0)] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Unexpected token " + prevTerm + " \nExpected [OPERATOR] [TERM]" 
				break
			}

			if nextTerm != "AND" && nextTerm != "OR" && prevTerm != "NOT" {
				flags[min(i+1, length-1)] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Unexpected token " + nextTerm + " \nExpected [TERM] [OPERATOR]" 
				break
			}
		}
	}

	if foundError {
		printErrorMessage(errorMsg, tokenizedQuery, flags)
	}

	return foundError
}

func printErrorMessage(errorMsg string, tokenizedQuery, flags []string) {
	var errorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF2C2C"))
	var bold = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#e1e1e1"))
	var arrow = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffb914"))
	
	fmt.Println(errorStyle.Render(errorMsg))

	for i := 0; i < len(tokenizedQuery); i++ {
		fmt.Print(bold.Render(tokenizedQuery[i]), " ")
	}
	fmt.Print("\n")
	for i := 0; i < len(flags); i++ {
		var spaces string = flags[i]

		if flags[i] != "^" {
			spaces = strings.Repeat(" ", len(flags[i]))
		}

		fmt.Print(arrow.Render(spaces), " ")
	}
	fmt.Print("\n")
}