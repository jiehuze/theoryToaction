package main

import (
	"net"
	"fmt"
	"flag"
	"sync"
)

func main() {
	input := flag.String("c", "ok", "get c")
	flag.Parse()
	waitgroup := sync.WaitGroup{}
	connectServer(input, &waitgroup)
	waitgroup.Wait()
}

func connectServer(input *string, waitgroup *sync.WaitGroup) {
	conn, err := net.Dial("tcp", "59.110.155.41:8088")
	//conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("connect failed"+err.Error())
		return
	}
	fmt.Println("connect success")
	waitgroup.Add(1)
	senddata := fmt.Sprintf("send message is 凯泽，我的好儿子, %s", *input)
	sendbyte := []byte(senddata)
	num, err := conn.Write(sendbyte)

	if err != nil {
		fmt.Printf("write data error %s\n", err.Error())
		return
	}

	fmt.Printf("send data num: %d, data: %s\n", num, senddata)
}
