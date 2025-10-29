package main

import "fmt"

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false // 移除分号
	}
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	} // 改进 map 的格式化
	stack := []rune{}

	for _, char := range s {
		if closing, exist := pairs[char]; exist {
			stack = append(stack, closing)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
}
