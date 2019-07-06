package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func init() {
	runtime.SetBlockProfileRate(1)
}

func main() {
	log.Fatal(http.ListenAndServe(":6060", nil))
}

// 测试死锁
//var quit chan int
//var c chan int
//
//func init() {
//	c, quit := make(chan int), make(chan int)
//	go func() {
//		c <- 1
//		quit <- 0
//	}()
//}
//
//func main() {
//	<-quit
//	<-c
//}
