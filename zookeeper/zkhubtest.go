/**
 * @Author: jiehu
 * @Description:
 * @Project: test
 * @File:  zkhubtest
 * @CopyRight: fuxi
 * @Date: 2020/7/7 2:19 下午
 */
package main

import (
	log "fuGin/utils"
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"time"

	"theoryToaction/zookeeper/zkhub"
)

func readW(zk *zkhub.ZookeeperHub) {

	for {
		_, _, event, err := zk.Conn.ExistsW("/mytest4")
		//_, _, event, err := zk.Conn.GetW("/mytest4")
		if err != nil {
			log.Info("-----------err:", err.Error())
			time.Sleep(time.Duration(1))
			continue
		}
		log.Info("---------- readw ---------------")
		select {
		case evt := <-event:
			log.Info("======readw evet path:  ", evt.Path)
			log.Info("======readw evet type: ", evt.Type, "===stat: ", evt.State)
		}
	}
}

func childrenW(zk *zkhub.ZookeeperHub) {

	for {
		_, _, event, err := zk.Conn.ChildrenW("/mytest4")
		//_, _, event, err := zk.Conn.GetW("/mytest4")
		if err != nil {
			log.Info("-----------err:", err.Error())
			time.Sleep(time.Duration(1))
			continue
		}
		log.Info("---------- readw ---------------")
		select {
		case evt := <-event:
			log.Info("======childrenW evet path:  ", evt.Path)
			log.Info("======childrenW evet type: ", evt.Type, "===stat: ", evt.State)
		}
	}
}

func main() {

	wg := &sync.WaitGroup{}

	zkInstance := zkhub.GetZKHubInstance()
	err := zkInstance.Init("127.0.0.1:2181", time.Duration(5))
	if err != nil {
		log.Infof("connect error !")
		return
	}

	go zkInstance.Start()

	go readW(zkInstance)
	go childrenW(zkInstance)

	//参数 4：为znode的类型
	str, err := zkInstance.Conn.Create("/mytest4", []byte("nihao"), int32(1), zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Info(err.Error())
	}
	log.Info("=========str: ", str)

	//exist, stat, err := zkInstance.Conn.Exists("/mytest")
	//
	//log.Info("=========exist: ", exist, "=====stat:", stat.Aversion, "======err:", err)
	////var event <-chan zk.Event
	//st, stat, event, err := zkInstance.Conn.GetW("/mytest")
	//
	////st, stat, err := zkInstance.Conn.Get("/mytest")
	//log.Info("=========st: ", string(st), "=====stat:", stat.Aversion)
	//
	//evt := <-event
	//log.Info("type: ", evt.Type, "---path", evt.Path)

	wg.Add(1)

	log.Info("=========wg wait ")
	wg.Wait()
	log.Info("=========wg close ")
}
