package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	var answer string

	fmt.Println("半晚六点下班")
	b()
	go c()
	go d()
	go e()
	go f()
	go g()
	go h()
	go i()
	go j(ch)


	answer = <-ch

	if answer != "" {
		fmt.Println(answer)
	} else {
		fmt.Println("Dead ....")
	}

}

func b() {
		fmt.Println("换掉药厂的衣裳")
		time.Sleep(0.1 * 1e9)
	}
func c() {
	fmt.Println("妻子在熬粥")
	time.Sleep(0.2 * 1e9)
}

func d() {
	fmt.Println("我去喝几瓶啤酒")
	time.Sleep(0.3 * 1e9)
}
func e() {
	fmt.Println("我如此生活三十年")
	time.Sleep(0.4 * 1e9)
}
func f() {
	fmt.Println("直到大厦崩塌")
	time.Sleep(0.5 * 1e9)
}
func g() {
	fmt.Println("云层深处的黑暗呢")
	time.Sleep(0.6 * 1e9)
}
func h() {
	fmt.Println("淹没心里的景观")
	time.Sleep(0.7 * 1e9)
}
func i() {
	fmt.Println("在八角柜台")
	time.Sleep(0.8 * 1e9)
}
func j(ch chan string) {
	time.Sleep(1 * 1e9)
	ch <- "疯狂的人民商场"
}