package stack

import "fmt"

type Stack struct {
	elements []interface{}
	len int64
}

// 判断字符串括号是否合法
func judgeStrIsValid(s string) bool {
	stack := new(Stack)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			stack.Push(string(s[i]))
			continue
		}

		if string(s[i]) == ")" {
			len, _ := stack.Pop()
			if len < 0 {
				fmt.Printf("字符串：%s, 第 %d 个元素\n", string(s[i]), i)

				return false
			}
		}
	}

	return true
}

// 查找数组中左边比当前小的值
func searchLeftMin(arr []int) int {
	var index int
	stack := NewStack()

	for k, v := range arr {
		if k > 0 {
			_, ele := stack.Pop()
			fmt.Println(ele)
		}

		stack.Push(v)
	}

	return index
}

func (s *Stack) MemberCount() int64 {
	return s.len
}

func NewStack() *Stack {
	stack := new(Stack)
	elements := make([]interface{}, 0)
	stack.elements = elements

	return stack
}

func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
	s.len ++
}

func (s *Stack) Pop() (int, interface{}) {
	if len(s.elements) == 0 {
		return -1, nil
	}

	if len(s.elements) == 1 {
		element := s.elements[len(s.elements)-1]
		s.elements = make([]interface{}, 0)
		s.len = 0

		return 0, element
	}

	if len(s.elements) > 1 {
		element := s.elements[len(s.elements)-1]
		s.elements = s.elements[:len(s.elements)-1]
		s.len --

		return len(s.elements), element
	}

	return -1, nil
}


