package oneeight

import (
	"log"
	"strconv"
	"unicode"
)

const (
	lparen = '('
	rparen = ')'
	num    = '0'
	add    = '+'
	mul    = '*'
)

type token struct {
	typ rune
	val int
}

func parse(tokens chan token) int {
	var n, res int
	o := '?'

	for i := 0; i < 2; i++ {
		for tok := range tokens {
			switch tok.typ {
			case add, mul:
				o = tok.typ
			case num:
				n = tok.val
			case lparen:
				n = parse(tokens)
			case rparen:
				return res
			}

			if tok.typ == num || tok.typ == lparen {
				if res == 0 {
					res = n
				} else {
					switch o {
					case '*':
						res *= n
					case '+':
						res += n
					}
				}
			}
		}
	}

	return res
}

func parseV2(ch chan token) int {
	var val1, val2 int
	var op rune
	values := make([]int, 0)
	ops := make([]rune, 0)
	i := 0
	tokens := make([]token, 0)
	for tok := range ch {
		tokens = append(tokens, tok)
	}

	for i < len(tokens) {
		switch tokens[i].typ {
		case lparen:
			ops = append(ops, lparen)
		case num:
			values = append(values, tokens[i].val)
		case rparen:
			for len(ops) != 0 && ops[len(ops)-1] != lparen {
				values, val2 = pop(values)
				values, val1 = pop(values)
				ops, op = popR(ops)
				values = append(values, apply(val1, val2, op))
			}
			ops, _ = popR(ops)
		case add, mul:
			for len(ops) != 0 && precedence(ops[len(ops)-1]) >= precedence(tokens[i].typ) {
				values, val2 = pop(values)
				values, val1 = pop(values)
				ops, op = popR(ops)
				values = append(values, apply(val1, val2, op))
			}
			ops = append(ops, tokens[i].typ)
		}
		i++
	}

	for len(ops) != 0 {
		values, val2 = pop(values)
		values, val1 = pop(values)
		ops, op = popR(ops)
		values = append(values, apply(val1, val2, op))
	}

	return values[len(values)-1]
}

func precedence(r rune) int {
	switch r {
	case '+':
		return 2
	case '*':
		return 1
	}
	return 0
}

func apply(a, b int, op rune) int {
	switch op {
	case '*':
		return a * b
	case '+':
		return a + b
	default:
		log.Fatalf("invalid op: %c", op)
	}
	return -1
}

func pop(sl []int) ([]int, int) {
	el := sl[len(sl)-1]
	return sl[:len(sl)-1], el
}

func popR(sl []rune) ([]rune, rune) {
	el := sl[len(sl)-1]
	return sl[:len(sl)-1], el
}

func lex(expr string) chan token {
	ch := make(chan token)
	go func() {
		for _, r := range expr {
			switch {
			case r == '(':
				ch <- token{typ: lparen}
			case r == ')':
				ch <- token{typ: rparen}
			case r == '*':
				ch <- token{typ: mul}
			case r == '+':
				ch <- token{typ: add}
			case unicode.IsDigit(r):
				n, _ := strconv.Atoi(string(r))
				ch <- token{typ: num, val: n}
			}
		}
		close(ch)
	}()
	return ch
}
