package interpreter

import "fmt"

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

			if  prevTerm == "NOT" {
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

		currTerm := tokenizedQuery[i]
		if length > 1 && i != 0 && i != length-1 && currTerm != "AND" && currTerm != "OR" && currTerm != "NOT" {
			prevTerm := tokenizedQuery[max(i-1, 0)]
			nextTerm := tokenizedQuery[min(i+1, length-1)]

			if prevTerm != "AND" && prevTerm != "OR" && prevTerm != "NOT" {
				flags[max(i-1, 0)] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Unexpected token " + prevTerm + " \nExpected [OPERATOR] [TERM]" 
				break
			}

			if nextTerm != "AND" && nextTerm != "OR" && nextTerm != "NOT" {
				flags[min(i+1, length-1)] = "^"
				foundError = true
				errorMsg = "SYNTAX ERROR: Unexpected token " + nextTerm + " \nExpected [TERM] [OPERATOR]" 
				break
			}
		}
	}

	fmt.Println(errorMsg)
	fmt.Println(tokenizedQuery)
	fmt.Println(flags)

	return foundError
}