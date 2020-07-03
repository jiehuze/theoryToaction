package main

import (
	"time"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"os"
	"test/thriftrpc/gen-go/batu/demo"
	"strconv"
	"net"
)

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

func main() {
	startTime := currentTimeMillis()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "19909"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport, _ := transportFactory.GetTransport(transport)


	//两种方式，1）
	//iprot := protocolFactory.GetProtocol(useTransport)
	//oprot := protocolFactory.GetProtocol(useTransport)
	//client := demo.NewBatuThriftClient(thrift.NewTStandardClient(iprot, oprot))
	//第2中方式 2）
	client := demo.NewBatuThriftClientFactory(useTransport, protocolFactory)

	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:19909", err)
		os.Exit(1)
	}

	defer transport.Close()

	for i := 0; i < 10; i++ {
		paramMap := make(map[string]string)
		paramMap["a"] = "batu.demo"
		paramMap["b"] = "theoryToaction" + strconv.Itoa(i+1)
		r1, _ := client.CallBack(nil, startTime, "go client", paramMap)
		fmt.Println("GOClient Call->", r1)
	}

	article := &demo.Article{
		ID: 100,
		Title: "测试title",
		Content: "测试content",
		Author: "theoryToaction author",
	}

	client.Put(nil, article)
}
