package link_pool

import "sync"

type LinkList struct {
	sync.RWMutex

	Id int
	coursor int
	indexs []int
	links []*Link
}
