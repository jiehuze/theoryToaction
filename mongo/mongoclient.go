package main

import (
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
	"math/rand"
	"gopkg.in/mgo.v2/bson"
)

func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}

type GameReport struct {
	// id          bson.ObjectId `bson:"_id"`
	Game_id     int64
	Game_length int64
	Game_map_id string
}

type GameReport2 struct {
	// id          bson.ObjectId `bson:"_id"`
	Game_id     int64
	Game_length int64
	Game_map_id string
	Game_user   string
}

func main() {
	fmt.Printf("start mongo theoryToaction\n")
	//初始化连接mongo info
	dailInfo := &mgo.DialInfo{
		Addrs:     []string{"10.110.93.157"},
		Direct:    false,
		Timeout:   time.Second * 1,
		Database:  "Call",
		Source:    "admin",
		Username:  "",
		Password:  "",
		PoolLimit: 1024,
	}

	//连接mongo服务器
	session, err := mgo.DialWithInfo(dailInfo)
	if err != nil {
		fmt.Printf("mongo DialWithInfo error")
		err_handler(err)
	}
	//异常处理连接失败
	defer session.Close()

	//获取设计库中的table/collect，对表进行操作
	c := session.DB("gorun").C("te")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	report := GameReport{
		// id:          bson.NewObjectId(),
		Game_id:     100,
		Game_length: r.Int63() % 3600,
		Game_map_id: "hello",
	}

	//插入操作
	err = c.Insert(report)
	if err != nil {
		fmt.Printf("try insert record error[%s]\n", err.Error())
		err_handler(err)
	}

	report2 := GameReport2{
		Game_id:     200,
		Game_length: r.Int63() % 3600,
		Game_map_id: "hello2",
		Game_user:   "testuser",
	}

	err = c.Insert(report2)
	if err != nil {
		fmt.Printf("try insert record error[%s]\n", err.Error())
		err_handler(err)
	}

	fmt.Printf("end mongo theoryToaction\n")

	result := GameReport{}
	var to_find_game_id int64 = 100
	c.Find(bson.M{"game_id": to_find_game_id}).One(&result)
	if err != nil {
		fmt.Printf("try find record error[%s]\n", err.Error())
		err_handler(err)
	}
	fmt.Printf("res, game_id[%d] length[%d] game_map_id[%s]\n",
		to_find_game_id, result.Game_length, result.Game_map_id)

	//err = c.Update(bson.M{"game_length": 2539}, bson.M{"$set": bson.M{"game_length": 577, "game_map_id": "hel"}})
	//if err != nil {
	//	fmt.Printf("collect update error: %s\n", err)
	//}

	// update 中的updateinterface必须增加"$set"，否则会更新所有参数
	_, err = c.UpdateAll(bson.M{"game_id": 100}, bson.M{"$set": bson.M{"game_map_id": "hello world"}})
	if err != nil {
		fmt.Printf("collect update error: %s\n", err)
	}
	var results []GameReport
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		fmt.Printf("try game all record of game_detail_report error[%s]\n",
			err.Error())
		err_handler(err)
	}

	//result_count, err := c.Find(bson.M{}).Count()

	//result_count := len(results)
	//fmt.Printf("result count: %d\n", result_count)
	//for i, report := range results {
	//	fmt.Printf("index: %d, report{ game_id: %d, game_length: %d, game_map_id: %s}\n",
	//		i, report.Game_id, report.Game_length, report.Game_map_id)
	//}

	c.Remove(bson.M{"game_length": 577})
	//c.RemoveAll(bson.M{"game_id": 100})

	c.Find(bson.M{}).Limit(10).All(&results)
	result_count := len(results)
	fmt.Printf("result count: %d\n", result_count)
	for i, report := range results {
		fmt.Printf("index: %d, report{ game_id: %d, game_length: %d, game_map_id: %s}\n",
			i, report.Game_id, report.Game_length, report.Game_map_id)
	}
	time.Sleep(time.Second * 3)
}
