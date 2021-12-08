package mytime

import (
	"fmt"
	"time"
)

func TestTime()  {
	//fmt.Println(time.Now())
	//fmt.Println(time.Now().Year())
	//fmt.Println(time.Now().Month())
	//fmt.Println(time.Now().Day())
	//fmt.Println(time.Now().Hour())
	//fmt.Println(time.Now().Second())
	//fmt.Println(time.Now().Nanosecond())

	currentTime:=time.Now()     //获取当前时间，类型是Go的时间类型Time
	t1:=time.Now().Year()        //年

	t2:=time.Now().Month()       //月

	t3:=time.Now().Day()         //日

	t4:=time.Now().Hour()        //小时

	t5:=time.Now().Minute()      //分钟

	t6:=time.Now().Second()      //秒

	t7:=time.Now().Nanosecond()  //纳秒
	currentTimeData:=time.Date(t1,t2,t3,t4,t5,t6,t7,time.Local) //获取当前时间，返回当前时间Time

	fmt.Println(currentTime)       //打印结果：2017-04-11 12:52:52.794351777 +0800 CST
	fmt.Println(t1,t2,t3,t4,t5,t6)     //打印结果：2017 April 11 12 52 52

	fmt.Println(currentTimeData)    //打印结果：2017-04-11 12:52:52.794411287 +0800 CST

	timeStr:=time.Now().Format("2006-01-02 15:04:05")  //当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法

	fmt.Println(timeStr)    //打印结果：2017-04-11 13:24:04

	timeUnix:=time.Now().Unix()   //已知的时间戳

	timeUnix=timeUnix - 3600*24*3

	formatTimeStr:=time.Unix(timeUnix,0).Format("2006-01-02 15:04:05")

	fmt.Println(formatTimeStr)   //打印结果：2017-04-11 13:30:39
}
