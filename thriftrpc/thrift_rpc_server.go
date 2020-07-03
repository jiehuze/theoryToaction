package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"time"
	"test/thriftrpc/gen-go/batu/demo"
	"context"
)

type batuThrift struct {

}

func (this *batuThrift) CallBack(ctx context.Context, callTime int64, name string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->from client Call:", time.Unix(callTime, 0).Format("2006-01-02 15:04:05"), name, paramMap)
	r = append(r, "key:"+paramMap["a"]+"    value:"+paramMap["b"])
	return
}

func (this *batuThrift)Put(ctx context.Context, s *demo.Article)  (err error){
	fmt.Printf("Article--->id: %d\tTitle:%s\tContent:%t\tAuthor:%d\n", s.ID, s.Title, s.Content, s.Author)
	return nil
}

func main() {
	transforsocket, err := thrift.NewTServerSocket("127.0.0.1:19909")
	if err != nil{
		fmt.Println("thrift newTSocket failed")
		return
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	//handler := &batuThrift{}
	//processor := demo.NewBatuThriftProcessor(handler)

	handler := &batuThrift{}
	processor := demo.NewBatuThriftProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, transforsocket, transportFactory, protocolFactory)
	fmt.Printf("start server!")
	server.Serve()
}
