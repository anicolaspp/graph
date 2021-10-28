package counter

type Linear struct {
	value int
}

func (l *Linear) Next() int {
	l.value++
	return l.value
}
