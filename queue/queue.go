package queue

type Queue []string

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(word string) {
	*q = append(*q, word)
}

func (q *Queue) Dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}

	element := (*q)[0]
	*q = (*q)[1:] // Keep everything , except the first one
	return element, true
}
