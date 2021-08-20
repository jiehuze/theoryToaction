package main

import (
	"log"
	"theoryToaction/hbase/client"
	"theoryToaction/hbase/options"
)

func main() {
	client.InitHBase()
	options.ScanRowKey()
	log.Print("hello world!")
}
