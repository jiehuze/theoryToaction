package options

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
	"github.com/tsuna/gohbase/pb"
	"strings"
	"theoryToaction/hbase/client"
)

func ScanRowKey() error {
	var scan *hrpc.Scan
	var err error

	prefixFilter := filter.NewColumnPrefixFilter([]byte("basicTotalMileage"))

	scan, err = hrpc.NewScanStr(context.Background(), "signalMirror", hrpc.Filters(prefixFilter))
	if err != nil {
		log.Fatal(err)
	}
	scanner := client.HBaseClient.Scan(scan)
	alist := make([]string, 0)
	for {
		scan, err := scanner.Next()
		if err != nil {
			break
		}

		for k, cell := range scan.Cells {
			log.Infof("%d row: %s, %s, %s, %s", k, string((*pb.Cell)(cell).GetRow()), string((*pb.Cell)(cell).GetValue()), string((*pb.Cell)(cell).GetFamily()), string((*pb.Cell)(cell).GetQualifier()))
			if !strings.Contains(string((*pb.Cell)(cell).GetRow()), "LTW") {
				alist = append(alist, string((*pb.Cell)(cell).GetRow()))
			}
		}
	}

	for k, v := range alist {
		log.Infof("** -------- num: %v, rowKey: %v", k, v)
		//delStr, _ := hrpc.NewDelStr(context.Background(), "signalMirror", v, nil)
		//result, _ := HBaseClient.Delete(delStr)
		//log.Info("######---- result: ", result)
	}

	return nil
}
