
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
	}()

	mu.Unlock()
}
//lock与unlock并行了，要先lock再unlock
