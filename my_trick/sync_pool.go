package my_trick

import (
	"fmt"
	"sync"
)

//var bufferpool=sync.Pool{
//	New: func() interface{} {
//		return 0
//	},
//}
func TestSyncPool() {
	var count int32 = 0
	var bufferpool = sync.Pool{
		New: func() interface{} {
			count = count + 1
			//atomic.AddInt32(&count,1)
			return count
		},
	}
	var wg sync.WaitGroup
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			i := bufferpool.Get().(int32)
			fmt.Println("i: ", i)
			//atomic.AddInt32(&i,1)
			defer bufferpool.Put(i)
		}()
	}
	wg.Wait()

}
