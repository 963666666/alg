package arr

import (
	"fmt"
	"sync"
	"time"
)

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] != target {
			return -1
		}else {
			return 0
		}
	}

	middleIndex := len(nums) / 2
	if nums[middleIndex] == target {
		return middleIndex
	} else if nums[middleIndex] > target {
		return search(nums[0:middleIndex], target)
	}else {
		return search(nums[middleIndex:len(nums)-1], target)
	}
}

func strOverReturn(str string) string {
	strBytes := []byte(str)

	perStart := 0
	perEnd := len(str)

	for perStart < perEnd {
		strBytes[perStart], strBytes[perEnd-1] = strBytes[perEnd-1], strBytes[perStart]
		perStart += 1
		perEnd -= 1
		fmt.Printf("perStart : %d, perEnd : %d\n", perStart, perEnd)
	}

	return string(strBytes)
}

//给定两个字符串 s 和 t，编写一个函数来判断 t 是否是 s 的字母异位词。
func MemberEctopic(s, t string) bool {
	return false
}

func GoProcess() int32 {
	queue := make(chan bool, 50)
	wg := &sync.WaitGroup{}
	var result int32

	cityIDs := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for _, cityID := range cityIDs {
		queue <- true
		wg.Add(1)

		go func(cityID int32) {
			if cityID == 10 {
				time.Sleep(1 * time.Second)
				result = cityID
			}
			fmt.Println("-------", result)

			<-queue
			wg.Done()
		}(cityID)
	}
	wg.Wait()

	return result
}
