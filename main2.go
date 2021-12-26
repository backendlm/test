package main

import "fmt"

func main() {
ch :=make(chan string)
	go func() {
		ch <-"下山的路又堵起了"
	}()
   msg := <-ch
	fmt.Println(msg)
}