package list

type List struct {
	Next *List
	Val  interface{}

	length
}

func NewSingle() *List {

}

func (l *List) AddAtTail(val interface{}) {
	l.Next = &List{
		Val: val,
	}

	l.length++
}

func (l *List) AddAtHead(val interface{}) {

}

func (l *List) GetByIndex(index int) interface{} {
	return nil
}

func (l *List) AddAtIndex(index int, val interface{}) {

}

func (l *List) DeleteAtIndex(index int) {

}
