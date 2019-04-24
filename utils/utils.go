package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"words/callback"

	"github.com/garyburd/redigo/redis"
)

func Show() {
	fmt.Println("hello world.")
}

// 配置文件格式按照key=value格式存储，不进行分级配置了
func LoadConfigurations(filepath string) (map[string]string, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		// return nil, err
		log.Fatal(err)
	}
	// fmt.Println(string(bytes))
	// var redisConfigs model.RedisConfigs
	var configurations map[string]string
	json.Unmarshal(bytes, &configurations)
	return configurations, nil
}

// 通用型哈希码计算
func GeneralHashCode(origin string) (string, error) {
	bytes := []byte(origin)
	ret := make([]byte, 1000)
	for index, _ := range bytes {
		ret = append(ret, bytes[index]>>1+bytes[index]>>2+bytes[index]>>3)
	}
	return string(ret), nil
}

// redis-client在函数终止时会defer close链接对象，因此加入callback策略进行具体的处理
func RedisRun(network string, host string, callback callback.DoRedisNoReturn, command string, params ...interface{}) (bool, error) {
	client, err := redis.Dial(network, host)
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	_, err = callback(client, command, params...)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 获取今天格式化的日期串
func GetCurDate() string {
	return time.Now().Format("2006-01-02")
}
