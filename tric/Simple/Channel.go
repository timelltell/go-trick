package Simple

import (
	"fmt"
	"sync"
)

func Channel_run(){
	done:=make(chan int)
	go func(){
		fmt.Println("channel hello world")
		<-done
	}()
	done<-1
}
//根据Go语言内存模型规范，对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。因此，后台线程<-done接收操作完成之后，main线程的done <- 1发送操作才可能完成（从而退出main、退出程序），而此时打印工作已经完成了。
//上面的代码虽然可以正确同步，但是对管道的缓存大小太敏感：如果管道有缓存的话，就无法保证main退出之前后台线程能正常打印了。更好的做法是将管道的发送和接收方向调换一下，这样可以避免同步事件受管道缓存大小的影响：
func Channel_run2(){
	done:=make(chan int ,1)
	go func(){
		fmt.Println("channel2 hello world")
		done<-1
	}()
	<-done
}
//对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是Channel的缓存大小。虽然管道是带缓存的，main线程接收完成是在后台线程发送开始但还未完成的时刻，此时打印工作也是已经完成的。
//基于带缓存的管道，我们可以很容易将打印线程扩展到N个。下面的例子是开启10个后台线程分别打印：
func Channnel_run3(){
	done:=make(chan int ,10)
	for i:=0;i<10;i++{
		go func(){
			fmt.Println("channel3 hello world")
			done<-1
		}()
	}
	for i:=0;i<10;i++{
		<-done
	}
}
//对于这种要等待N个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用sync.WaitGroup来等待一组事件：
func Wait_group(){
	var wai sync.WaitGroup
	for i:=0;i<10;i++{
		wai.Add(1)
		go func(){
			fmt.Println("wait hello world")
			wai.Done()
		}()
	}
	wai.Wait()
}