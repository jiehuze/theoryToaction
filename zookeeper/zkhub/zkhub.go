/**
 * @Author: jiehu
 * @Description:
 * @Project: test
 * @File:  zk
 * @CopyRight: fuxi
 * @Date: 2020/7/7 2:18 下午
 */
package zkhub

import (
	"github.com/golang/glog"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
)

const (
	ChildrenW = 0
	ExistW    = 1
	GetW      = 2
)

type WatchHandler interface {
	Do(zk.Event)
}

var instanceZKHub *ZookeeperHub

func GetZKHubInstance() *ZookeeperHub {
	if instanceZKHub == nil {
		instanceZKHub = newZKHub()
	}
	return instanceZKHub
}

type ZookeeperHub struct {
	ZaAddress string
	Done      chan bool
	Conn      *zk.Conn
	ConnEvent <-chan zk.Event
}

func newZKHub() *ZookeeperHub {
	r := &ZookeeperHub{}
	r.init()
	return r
}

func (zkh *ZookeeperHub) init() {
	zkh.ConnEvent = make(<-chan zk.Event)
	zkh.Done = make(chan bool)
}

func (zkh *ZookeeperHub) Init(address string, sessionTimeoutSeconds time.Duration) error {
	zkh.ZaAddress = address
	return zkh.initZKHubConn(sessionTimeoutSeconds)
}

func (zkh *ZookeeperHub) Start() {
	zkh.run()
}

func (zkh *ZookeeperHub) Stop() {
	zkh.Done <- false
}

func (zkh *ZookeeperHub) dial(sessionTimeoutSeconds time.Duration) (*zk.Conn, <-chan zk.Event, error) {
	zksStr := zkh.ZaAddress
	zks := strings.Split(zksStr, ",")
	return zk.Connect(zks, sessionTimeoutSeconds*time.Second) //async connect to zookeeper
}

func (zkh *ZookeeperHub) initZKHubConn(sessionTimeoutSeconds time.Duration) error {
	glog.V(0).Infoln("ZKHub initZKHubConn() zkAddress = ", zkh.ZaAddress)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		conn, event, err := zkh.dial(sessionTimeoutSeconds)
		if err != nil {
			glog.Errorln("ZKHub err:%s initZKHubConn() trying to reconnect in 5 seconds...", err)
			continue
		}
		zkh.Conn = conn
		zkh.ConnEvent = event
		glog.V(0).Infoln("ZKHub init zk conn over")
		return nil
	}
}

func (zkh *ZookeeperHub) run() {
	for {
		select {
		case event := <-zkh.ConnEvent:
			glog.V(0).Infoln("ZkHub receive connection event = ", event)

			if event.State == zk.StateExpired { //TODO it
				//glog.Info("ZkHub session has expired  kill self bye bye.")
				glog.Fatal("ZkHub session has expired  kill self bye bye.")
			}
		case <-zkh.Done:
			zkh.Conn.Close()
		}
	}
}

func (zkh *ZookeeperHub) WatchChildren(path string, handler WatchHandler) error {
	return zkh.watchIt(path, ChildrenW, handler)
}
func (zkh *ZookeeperHub) WatchExist(path string, handler WatchHandler) error {
	return zkh.watchIt(path, ExistW, handler)
}
func (zkh *ZookeeperHub) WatchGet(path string, handler WatchHandler) error {
	return zkh.watchIt(path, GetW, handler)
}

func (zkh *ZookeeperHub) watchIt(path string, wtype uint8, handler WatchHandler) error {
	watchErr := make(chan error)
	first := true
	go func() {
		var event <-chan zk.Event
		var err error
		for {
			glog.V(0).Infoln("ZKHub Watch  wtype = ", wtype, ", path = ", path)
			switch wtype {
			case ChildrenW:
				_, _, event, err = zkh.Conn.ChildrenW(path)
			case ExistW:
				_, _, event, err = zkh.Conn.ExistsW(path)
			case GetW:
				_, _, event, err = zkh.Conn.GetW(path)
			}
			if first == true {
				watchErr <- err
				if err != nil {
					return
				}
				first = false
			}

			//			if err != nil {//TODO ?
			//				glog.Errorln("ZKHub WatchIt() Event Error,Ignore it  , retry again to wait reconnection with zk")
			//				time.Sleep(2 * time.Second)
			//				continue
			//			}

			evt := <-event
			if evt.Err != nil { //TODO?
				glog.Fatal("ZKHub WatchIt() Event Error, Exit Now") //TODO
				//				time.Sleep(2 * time.Second)
				//				fmt.Println("Debug : ZKHub WatchIt() Event Error,Ignore it  , retry again to wait reconnection with zk")
				//continue
			}
			go handler.Do(evt)
		}
	}()
	err := <-watchErr
	return err
}
