package correctparenthesis

type stack []rune

func createStack(l int) *stack {
	s := make(stack, 0, l)
	return &s
}

func (s *stack) size() int {
	return len(*s)
}

func (s *stack) isEmpty() bool {
	return s.size() == 0
}

func (s *stack) push(c rune) {
	*s = append(*s, c)
}

func (s *stack) pop() (rune, bool) {
	if s.isEmpty() {
		return 0, false
	}
	i := s.size() - 1
	el := (*s)[i]
	*s = (*s)[:i]
	return el, true
}

func CorrectParenthesis(str string) bool {
	s := createStack(len(str))
	for _, c := range str {
		switch c {
		case '(':
			s.push(')')
		case '[':
			s.push(']')
		case '{':
			s.push('}')
		default:
			if f, ok := s.pop(); !ok || f != c {
				return false
			}

		}
	}
	return s.isEmpty()
}
