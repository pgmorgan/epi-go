package ch09q03a

func CheckBrackets(s string) bool {
	// stack := make([]rune, len(s)) // This is wrong
	var stack []rune

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			if char == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if char == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if char == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
