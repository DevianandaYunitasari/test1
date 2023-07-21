package main

import "fmt"

func isBalanced(s string) string {
	stack := make([]rune, 0)

	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if closingBracket, ok := bracketPairs[char]; ok {

			if len(stack) == 0 {
				return "NO"
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top != closingBracket {
				return "NO"
			}
		} else {

			stack = append(stack, char)
		}
	}

	if len(stack) == 0 {
		return "YES"
	}

	return "NO"
}

func main() {
	// Sample input 1
	input1 := "{[()]}"
	fmt.Println("Sample 1:", isBalanced(input1))

	// Sample input 2
	input2 := "{[(])}"
	fmt.Println("Sample 2:", isBalanced(input2))

	// Sample input 3
	input3 := "{(([])[])[]}"
	fmt.Println("Sample 3:", isBalanced(input3))
}
