package linked_list

type TwoWayList struct {
	Front  *TwoWayList
	Next *TwoWayList
	
	val interface{}
	len int
}

func NewTwoWayList() *TwoWayList {
	return &TwoWayList{
		Front: nil,
	}
}