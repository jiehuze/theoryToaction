package main

import (
	"net"
	"fmt"
	"sync"
)

func main() {
	startServer()
}

func startServer() {
	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Println("listen tcp localhost error")
		return
	}

	waitgroup := &sync.WaitGroup{}
	for {
		conn, err := listen.Accept()
		if err != nil{
			fmt.Printf("accept is error ")
			break
		}
		println("%v", conn)
		waitgroup.Add(1)
		doServer(conn, waitgroup)
	}

	waitgroup.Wait()
}

func doServer(conn net.Conn, waitgroup *sync.WaitGroup)  {
	dataInfo := make([]byte, 512)

	num , err := conn.Read(dataInfo)
	if err != nil{
		fmt.Printf("read soket data failed")
		return
	}

	waitgroup.Done()

	fmt.Printf("get num: %d, data : %s\n", num, string(dataInfo))
	fmt.Printf("get remote addrs: %s\n", conn.RemoteAddr().String())
}
