package main

import "fmt"

func main() {
	str := "abcdefghijklmnopqrstuvwxyz"
	result := strOverReturn(str)

	fmt.Println(result)
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
