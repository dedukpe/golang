package piscine

func ListClear(l *List) {
	current := l.Head
	for current != nil {
		next := current.Next
		current.Next = nil
		current.Data = nil
		current = next
	}
	l.Head = nil
	l.Tail = nil
}
