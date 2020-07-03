package main

import (
	//"github.com/go-redis/redis"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
)

type RawApp struct {
	Pkg    string `json:"pkg"`
	UserId string `json:"userid"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
	Desc   string `json:"desc"`
	AppKey string `json:"appkey,omitempty"`
	AppSec string `json:"appsec,omitempty"`
	Quota  int    `json:"quota,omitempty"`
}

func main() {
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "10.112.32.41:16379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//defer client.Close()
	//
	//ss := client.HGetAll("db_apps")
	//
	//for _, v := range ss.Val() {
	//	fmt.Printf("%s ", v.([]byte))
	//}

	c, err := redis.Dial("tcp", "10.112.32.41:16379")
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}
	defer c.Close()

	//result, err := redis.Values(c.Do("hgetall", "db_apps"))
	map1, err := redis.StringMap(c.Do("hgetall", "db_apps"))
	var s [150]string
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("all keys and values are:")
		var i int
		i = 0
		//for _, v := range result {
		//	i++
		//	fmt.Printf("[%d] %s \n", i,  v.(map[string]))
		//}

		for key, val := range map1 {

			fmt.Printf("[%d] key: %s, value: %s\n", i, key, val)
			var rawapp RawApp
			err = json.Unmarshal([]byte(val), &rawapp)
			if err != nil {
				fmt.Println("json Unmarshal failed:", err)
			} else {
				s[i] = rawapp.Pkg
				i++
				fmt.Printf("[%d] key: %s, value: %s\n", i, key, rawapp.Pkg)
				_, err := c.Do("hset", "db_packages", rawapp.Pkg, key)
				if err != nil {
					fmt.Printf("redis hset pkg: %s error: n", rawapp.Pkg)
				} else {
					//fmt.Println("hset ok\n")
				}
			}
		}
	}

	fmt.Println("-------------------")
	for i := 0; i < 150; i++ {
		s1 := s[i]
		//fmt.Printf("==========%s\n", s1)
		if strings.Compare(s1, "") == 0 {
			fmt.Println("81 error")
			break
		}
		for j := i + 1; j < 150; j++ {
			if strings.Compare(s[j], "") == 0 {
				fmt.Println("89 error")
				break
			} else {
				//fmt.Println("++++++++++++++++++\n")
				if strings.Compare(s1, s[j]) == 0 {
					fmt.Printf("compara key %s is same", s1)
				}
			}

		}
	}
}
