package list

type List struct {
	Next *List
	Val  interface{}

	length int
}

func NewDummy() *List {
	return &List{
		Val: -1,
	}
}

func (l *List) AddAtTail(val interface{}) {
	l.Next = &List{
		Val: val,
	}

	l.length++
}

func (l *List) AddAtHead(val interface{}) {
	l.Next = &List{
		Val:  val,
		Next: l.Next.Next,
	}

	l.length++
}

func (l *List) GetByIndex(index int) interface{} {
	return nil
}

func (l *List) AddAtIndex(index int, val interface{}) {

}

func (l *List) DeleteAtIndex(index int) {

}

func (l *List) getPrevNode(index int) {
	var indexFront *List = l.Next
	var indexBack *List = l

	for i := 0; i < index; i++ {
		indexBack = indexFront
		indexFront = indexBack.Next
	}
}

func (l *List) Reverse() *List {
	newList := NewDummy()
	var prevNode = l.Next
	for i := 0; i < l.length; i++ {
		if prevNode == nil {
			return newList
		}

		newList.AddAtHead(prevNode.Val)
		prevNode = l.Next
	}

	return newList
}
