package queue

type List struct {
	elements []interface{}
}

func NewList() *List {
	list := new(List)
	list.elements = make([]interface{}, 0)

	return list
}

func (l *List) Push(element interface{}) {
	l.elements = append(l.elements, element)
}
func (l *List) Pop() (int, interface{}) {
	if len(l.elements) > 0 {
		element := l.elements[0]
		l.elements = l.elements[1:]

		return len(l.elements), element
	}

	return -1, nil
}

type CircularList struct {
	Use int
	Front int
	Rear int
	Capacity int
}