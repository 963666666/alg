package main

import (
	"fmt"
	"lag/api"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/user"
	"runtime/pprof"
	"runtime/trace"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/andlabs/ui/winmanifest"
)

type Info struct {
	UserIDs []int64

	Desc string
}

func main1() {
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
	day, _ := strconv.Atoi(date[9:])

	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		monthDays[1]++
	}
	if month > 11 {
		return -1
	}

	var days = day

	fmt.Println(monthDays[:month-1])

	for i := 0; i < month; i++ {
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

func main3() {
	threadProfile := pprof.Lookup("threadcreate")
	fmt.Printf("initthreadscounts:%d\n", threadProfile.Count())

	f, err := os.Create("threads_trace.out")
	if err != nil {
		panic(err)
	}

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	for i := 0; i < 20; i++ {
		go func() {
			_, err := exec.Command("bash", "-c", "sleep 3").Output()
			if err != nil {
				panic(err)
			}
		}()
	}
	time.Sleep(time.Second * 5)
	fmt.Printf("endthreadscounts:%d\n", threadProfile.Count())
}
