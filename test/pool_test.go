package test

import (
	"fmt"
	"lag/pool"
	"math/rand"
	"os"
	"runtime"
	"testing"

	"time"
)

func TestPool(t *testing.T) {
	fmt.Println(os.Hostname())

	rand.Seed(time.Now().Unix())
	// 这里限制5个并发
	p := pool.New(6)
	fmt.Println("the NumGoroutine begin is:", runtime.NumGoroutine())
	for i := 0; i < 20; i++ {
		p.Add()
		go func(i int) {
			randNum := rand.Intn(5)
			time.Sleep(time.Duration(randNum) * time.Second)
			fmt.Println("the NumGoroutine continue is:", runtime.NumGoroutine(), "time : ", randNum, "I : ", i)
			p.Down()
		}(i)
	}
	p.Wait()
	fmt.Println("the NumGoroutine done is:", runtime.NumGoroutine())
}
