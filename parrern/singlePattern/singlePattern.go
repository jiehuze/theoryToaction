package main

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type SinglePattern struct {
}

func (this *SinglePattern) get() string {
	return "hello"
}

var singleInstace *SinglePattern
var mu sync.Mutex
var once sync.Once

func GetSingleInstance() *SinglePattern {
	//if singleInstace == nil {
	//	mu.Lock()
	//	defer mu.Unlock() //入栈，函数执行完成后，才执行该条代码
	//	singleInstace = new(SinglePattern)
	//}

	//加锁的一次判断
	once.Do(func() {
		singleInstace = new(SinglePattern)
	})

	return singleInstace
}

func main() {
	single := GetSingleInstance()

	logrus.Info("get: ", single.get())
}
