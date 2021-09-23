package main

import (
	"fmt"
)

type heap struct {
	members []interface{}
}

func New() *heap {
	return &heap{
		members: make([]interface{}, 0),
	}
}

func (h *heap) Push(member interface{}) {

}

func (h *heap) Pop(member interface{}) {

}

func (h *heap) swap(i, j int) {
	h.members[i], h.members[j] = h.members[j], h.members[i]
}

func main() {
	m := make([]string, 0)

	for i := 10; i < 20; i++ {
		m = append(m, fmt.Sprintf("%d", i))
	}

	fmt.Println(m[:2], m[2:])

	fmt.Printf("format: %s", "hello world")
}
