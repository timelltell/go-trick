package Simple

import (
	"fmt"
	"sync"
)

func Run(){
	var mu sync.Mutex
	mu.Lock()
	go func(){
		fmt.Println("hello word")
		mu.Unlock()
	}()
	mu.Lock()
}