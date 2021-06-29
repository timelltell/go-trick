package Middle

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func RedisPrac(){
	var (
		RedisIp ="9.135.91.46"
		RedisPort="16379"
		expiretime = 300
		rdb *redis.Client
	)
	rdb=redis.NewClient(&redis.Options{
		Addr: RedisIp+":"+RedisPort,
		Password: "cluster",

	})


	_,err:=rdb.Ping().Result()
	if err!=nil{
		fmt.Println("redis conn failed")
	}else {
		fmt.Println("conn success")
	}


	a,err:=rdb.Exists("chenbin").Result()
	if err!=nil{
		fmt.Println("not exists")
	}else {
		fmt.Println("a: ",a)
	}


	err=rdb.Set("chenbin","chenbin",time.Duration(expiretime)*time.Second).Err()
	if err!=nil{
		fmt.Println("set key error")
	}else{
		fmt.Println("set success")
	}


	v,err:=rdb.Get("chenbin").Result()
	if err!=nil{
		fmt.Println("get key error")
	}else{
		fmt.Println("v: ",v)
	}

	err=rdb.Expire("chenbin",time.Duration(100)*time.Second).Err()
	if err!=nil{
		fmt.Println("Expire key error")
	}else{
		fmt.Println("Expire success")
	}


	key,err:=rdb.HSet("chenbin1","id","12313").Result()
	if err!=nil{
		fmt.Println("HSet key failed")
	}else{
		fmt.Println("HSet :",key)
	}


	key1,err:=rdb.HGet("chenbin1","id").Result()
	if err!=nil{
		fmt.Println("HGet error")
	}else{
		fmt.Println("HGet :",key1)
	}


	status,err:=rdb.HExists("chenbin1","id").Result()
	if err!=nil{
		fmt.Println("HExists failed")
	}else{
		fmt.Println("HExists :",status)
	}


	statusDel,err:=rdb.HDel("chenbin1","id").Result()
	if err!=nil{
		fmt.Println("HDel failed")
	}else{
		fmt.Println("HDel :",statusDel)
	}
	if 1 == statusDel {
		fmt.Println("删除hash值：id成功")
	}


	statusDel, err = rdb.Del("ming").Result()
	if 1 == statusDel {
		fmt.Println("删除值成功")
	}
}

