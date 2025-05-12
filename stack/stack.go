package stack

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(word string) {
	*s = append(*s, word)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index] // Upto , but not including the index
	return element, true
}
