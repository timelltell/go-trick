package ProConsum

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func product(factor int,out chan<- int){
	for i:=0;;i++{
		out<-i*factor
	}
}
func consumer(in <-chan int){
	for v:=range in{
		fmt.Println(v)
	}
}
func Product_Consumer(){
	ch:=make(chan int ,64)
	go product(3,ch)
	go product(5,ch)
	go consumer(ch)
	//运行一定时间退出
	//time.Sleep(5*time.Second)
	//按control c 退出
	sig:=make(chan os.Signal,1)
	signal.Notify(sig,syscall.SIGINT,syscall.SIGTERM)
	fmt.Printf("quit  %v",<-sig)

}