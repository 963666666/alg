package main

import (
	"fmt"
	"lag/tree"
	"strconv"
)

type Info struct {
	UserIDs []int64

	Desc string
}

func main() {
	treeNode := new(tree.TreeNode)

	treeNode.Val = 1
	treeNode.Left = &tree.TreeNode{
		Val: 2,
	}
	treeNode.Right = &tree.TreeNode{
		Val: 3,
	}

	result := treeNode.InorderTraversal()

	fmt.Printf("result : %+v\n", result)

	slice := make([]int64, 10)

	for i, v := range slice {
		fmt.Println("i : ", i, "v : ", v)
	}
}

func CalcDateToDays(date string) int {
	if len(date) != 10 {
		return -1
	}
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _:= strconv.Atoi(date[9:])

	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		monthDays[1]++
	}
	if month > 11 {
		return -1
	}

	var days = day

	fmt.Println(monthDays[:month-1])

	for i := 0; i < month; i ++ {
		days += monthDays[i]
	}

	return days
}
