package main

import (
	"fmt"
	"github.com/ndjordjevic/challenges/Stacks/stack"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("a+b*c+d infix to postfix: %s", infixToPostFix("a+b*c+d"))
	fmt.Printf("a+b*(c^d-e)^(f+g*h)-i infix to postfix: %s", infixToPostFix("a+b*(c^d-e)^(f+g*h)-i"))
}

func infixToPostFix(s string) string {
	var sb strings.Builder

	stk := &stack.Stack{}

	for _, value := range s {
		if unicode.IsLetter(value) || unicode.IsDigit(value) {
			sb.WriteRune(value)
		} else if value == '(' {
			stk.Push(value)
		} else if value == ')' {
			for !stk.IsEmpty() && stk.Peek() != '(' {
				v, err := stk.Pop()
				if err != nil {
					panic(err)
				}
				sb.WriteRune(v.(rune))
			}

			if !stk.IsEmpty() && stk.Peek() != '(' {
				return "Invalid expression"
			} else {
				_, _ = stk.Pop()
			}
		} else {
			for !stk.IsEmpty() && prec(value) <= prec(stk.Peek().(rune)) {
				if stk.Peek() == '(' {
					return "Invalid Expression"
				}

				v, err := stk.Pop()
				if err != nil {
					panic(err)
				}
				sb.WriteRune(v.(rune))
			}

			stk.Push(value)
		}
	}

	for !stk.IsEmpty() {
		if stk.Peek() == '(' {
			return "Invalid expression"
		}

		v, err := stk.Pop()
		if err != nil {
			panic(err)
		}
		sb.WriteRune(v.(rune))
	}

	return sb.String()
}

func prec(r rune) int {
	switch r {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	default:
		return -1
	}
}
