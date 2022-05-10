package linked_list

type OneWayList struct {
	Next *OneWayList
	Val  interface{}

	length int
}

func NewOneWayList() *OneWayList {
	return &OneWayList{
		Val: -1,
	}
}

func (l *OneWayList) AddAtTail(val interface{}) {
	l.Next = &OneWayList{
		Val: val,
	}

	l.length++
}

func (l *OneWayList) AddAtHead(val interface{}) {
	l.Next = &OneWayList{
		Val:  val,
		Next: l.Next.Next,
	}

	l.length++
}

func (l *OneWayList) GetByIndex(index int) interface{} {
	return nil
}

func (l *OneWayList) AddAtIndex(index int, val interface{}) {

}

func (l *OneWayList) DeleteAtIndex(index int) {

}

func (l *OneWayList) getPrevNode(index int) {
	var indexFront *OneWayList = l.Next
	var indexBack *OneWayList = l

	for i := 0; i < index; i++ {
		indexBack = indexFront
		indexFront = indexBack.Next
	}
}

func (l *OneWayList) Reverse() *OneWayList {
	newList := NewOneWayList()
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
