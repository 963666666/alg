package queue

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	list := new(Queue)
	list.elements = make([]interface{}, 0)

	return list
}

func (l *Queue) Push(element interface{}) {
	l.elements = append(l.elements, element)
}
func (l *Queue) Pop() (int, interface{}) {
	if len(l.elements) > 0 {
		element := l.elements[0]
		l.elements = l.elements[1:]

		return len(l.elements), element
	}

	return -1, nil
}

type CircularQueue struct {
	Used     int
	Front    int
	Rear     int
	Capacity int
	Members  []interface{}
}

func NewCircularQueue(cap int) *CircularQueue {
	circalarQueue := new(CircularQueue)
	circalarQueue.Capacity = cap
	circalarQueue.Members = make([]interface{}, 0)

	return circalarQueue
}

func (cq *CircularQueue) Push(member interface{}) bool {
	if cq.IsFull() {
		return false
	}

	cq.Used += 1

	cq.Members = append(cq.Members, member)

	return true
}

func (cq *CircularQueue) Pop() interface{} {
	if cq.IsEmpty() {
		return nil
	}

	cq.Used -= 1
	member := cq.Members[0]
	cq.Members = cq.Members[1:]

	return member
}

func (cq *CircularQueue) IsFull() bool {
	return cq.Used == cq.Capacity
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.Used == 0
}

func (cq *CircularQueue) Next(i int) interface{} {
	return cq.Members[(i+1)%cq.Capacity]
}

func (cq *CircularQueue) Last(i int) interface{} {
	return cq.Members[(i-1+cq.Capacity)%cq.Capacity]
}
