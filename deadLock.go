package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var lock sync.Mutex

func init() {
	go PeriodicTask()
	http.HandleFunc("/deadlock/", deadlockHandler)
}

func deadlockHandler(w http.ResponseWriter, _ *http.Request) {
	lock.Lock()
	if _, err := w.Write([]byte("deadlock")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PeriodicTask() {
	for {
		lock.Lock()
		fmt.Println("Hello world")
		time.Sleep(3 * time.Second)
		lock.Unlock()
	}
}
