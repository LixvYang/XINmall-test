// Package models provides 
package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

type cacheDB struct{}

var (
	CacheDB = &cacheDB{}

	redisClient cache.Cache
	YzmClient cache.Cache
    enableRedis, _ = beego.AppConfig.Bool("enableRedis")
	redisTime, _ = beego.AppConfig.Int("redisTime")
)



func init(){
  if enableRedis {
	  config := map[string]string {
		  	"key":		beego.AppConfig.String("redisKey"),
		  	"conn":		beego.AppConfig.String("redisConn"),
		  	"dbNum":    beego.AppConfig.String("redisDbNum"),
		  	"password": beego.AppConfig.String("redisPwd"),
	  }

	  redisconfig, _ := json.Marshal(config)
	  
	  redisClient, err = cache.NewCache("redis",string(redisconfig))
	  if err != nil {
		  beego.Error("connect redis failed")
	  } else {
		  beego.Info("connect redis succeed")
	  }
  }
}


//set data
func (c *cacheDB) Set(key string, value interface{})  {
	if enableRedis {
		bytes, err := json.Marshal(value)
		if err != nil {
			beego.Error(err)
		}
		redisClient.Put(key,string(bytes),time.Second * time.Duration(redisTime))
	}
}

//get data
func (c *cacheDB) Get(key string, obj interface{}) bool {
	if enableRedis {
		if redisStr := redisClient.Get(key); redisStr != nil {
			fmt.Println("在redis里面读取数据...")
			redisValue, ok := redisStr.([]uint8)
			if !ok {
				fmt.Println("获取redis数据失败")
				return false
			}
			json.Unmarshal([]byte(redisValue), obj)
			return true
		}
		return false
	}
	return false
}

