package Context

import (
	"context"
	"fmt"
	"time"
)

func Get(ctx context.Context,k string){
	if v,ok:=ctx.Value(k).(string);ok{
		fmt.Println(v)
	}
}
func Run1(){
	ctx:=context.WithValue(context.Background(),string("asong"),"hello")
	Get(ctx,string("asong"))
	Get(ctx,string("song"))
}

func Speak(ctx context.Context){
	for range time.Tick(time.Second){
		select{
		case <-ctx.Done():
			return
		default:
			fmt.Println("hello")
		}
	}
}
func Run2(){
	ctx,cancel:=context.WithCancel(context.Background())
	defer cancel()
	go Speak(ctx)
	time.Sleep(3*time.Second)
}
func Monitor(ctx context.Context){
	select{
	case<-ctx.Done():
		fmt.Println("hello")
	case <-time.After(5*time.Second):
		fmt.Println("done")
	}
}
func Run3(){
	ctx,cancel:=context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()
	go Monitor(ctx)
	time.Sleep(5*time.Second)
}