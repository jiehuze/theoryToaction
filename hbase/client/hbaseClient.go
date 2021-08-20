package client

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tsuna/gohbase"
)

var (
	// Version information. Passed from "go build -ldflags"
	HbaseQBMessageTable = "signalMirror"

	HBaseClient gohbase.Client

	online = "cloudtable-15c2-zk1-8kbOvYrv.cloudtable.com:2181,cloudtable-15c2-zk3-tGh7vMsM.cloudtable.com:2181,cloudtable-15c2-zk2-d9pLJqCF.cloudtable.com:2181"
)

func InitHBase() {
	HBaseClient = gohbase.NewClient(online)
	//getRequest, _ := hrpc.NewGetStr(context.Background(), utils.HBASE_TABLE, "LTWA35K17KS000420_1603272719000_0x09A")
	//getRsp, _ := HBaseClient.Get(getRequest)
	//for _, cell := range getRsp.Cells{
	//
	//	fmt.Println(string((*pb.Cell)(cell).GetFamily()))
	//
	//	fmt.Println(string((*pb.Cell)(cell).GetQualifier()))
	//
	//	fmt.Println(string((*pb.Cell)(cell).GetValue()))
	//
	//	fmt.Println((*pb.Cell)(cell).GetTimestamp())
	//
	//}
}
