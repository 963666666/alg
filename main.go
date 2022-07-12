package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lag/api"
	"net/http"
	_ "net/http/pprof"
	"os/user"
	"strconv"
)

type Info struct {
	UserIDs []int64

	Desc string
}

func main() {
	go http.ListenAndServe("0.0.0.0:6060", nil)

	server := gin.Default()

	server.GET("/index", api.Index)
	server.GET("/deepEqual", api.DeepEqual)

	server.Run(":9000")
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

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("user.Current err : %+v\n", err)

		return ""
	}

	return usr.HomeDir
}