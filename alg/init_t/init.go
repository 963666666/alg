package init_t

import (
	"fmt"
	_ "lag/alg/init_t/init_second"
)

func init() {
	fmt.Println("hello world init_t.init_t")
}

func Hello() string {
	return "world"
}