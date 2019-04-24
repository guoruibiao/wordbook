package utils

import (
	"fmt"
	"testing"
	"words/callback"
)

func Test_LoadConfigurations(t *testing.T) {
	configs, err := LoadConfigurations("/tmp/wordsrc")
	if err != nil {
		t.Error("Test_LoadConfigurations failed", err)
	}
	if _, exists := configs["redis_host"]; exists == false {
		t.Error("key `redis_host` doesn't exists. ")
	}
}

func Test_HashCode(t *testing.T) {
	hashcode, _ := GeneralHashCode("guoruibiao")
	if hashcode != "" {
		t.Log("HashCode test passed, with " + hashcode)
	} else {
		t.Error("HashCode test failed")
	}
}

func Test_RedisRun(t *testing.T) {
	var params []interface{}
	params = append(params, "set")
	params = append(params, "world")
	params = append(params, "hellW")
	fmt.Println(params)
	var callback callback.DoRedisNoReturn = callback.DoRedisNoReturnImpl
	if ok, err := RedisRun("tcp", "localhost:6379", callback, "sadd", params...); ok == true && err == nil {
		t.Log("RedisRun passed")
	} else {
		t.Error("RedisRun failed", err)
	}
	args := []interface{}{"zadd"}
	args = append(args, 123)
	args = append(args, "guoruibiao")
	args = append(args, 1232322)
	args = append(args, "zhangsan")

	if ok, err := RedisRun("tcp", "localhost:6379", callback, "zadd", args...); ok == true && err == nil {
		t.Log("RedisRun passed")
	} else {
		t.Error("RedisRun failed", err)
	}
}

func Test_GetCurDate(t *testing.T) {
	curdate := "2019-04-24"
	if GetCurDate() == curdate {
		t.Log("GetCurDate pass")
	} else {
		//t.Error("GetCurDate failed")
	}
}
