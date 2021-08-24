package main

import (
	"fmt"
)

func main() {
	fmt.Println(isBalanced("{[]}"))
	// print true
	fmt.Println(isBalanced("{[}]"))
	// print false

}
func isBalanced(str string) (balanced bool) {
	balanced = true
	rArray, rBrackets := make([]rune, 0, len(str)), map[rune]rune{
		'[': ']',
		'{': '}',
	}

	for _, strChar := range str {
		if _, ok := rBrackets[strChar]; ok == true {
			rArray = append(rArray, strChar)
		} else if len(rArray) == 0 || rBrackets[rArray[len(rArray)-1]] != strChar {
			balanced = false
			break
		} else {
			rArray = rArray[:len(rArray)-1]
		}
	}

	if len(rArray) != 0 {
		balanced = false
	}

	return
}
