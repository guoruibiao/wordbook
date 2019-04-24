package dao

import (
	"fmt"
	"testing"
)

func Test_GetMembers(t *testing.T) {
	redisconfig := WordDaoConfig{
		Redis_network: "tcp",
		Redis_host:    "localhost:6379",
	}
	dao, _ := New(redisconfig)
	members, err := dao.GetMembers("set")
	fmt.Println(members)
	if err != nil {
		t.Error("GetMembers failed")
	} else {
		fmt.Println(members)
		t.Log(members)
	}
}

func Test_AddWord(t *testing.T) {
	origin := "constraint"
	translate := "限制"
	redisconfig := WordDaoConfig{
		Redis_network: "tcp",
		Redis_host:    "localhost:6379",
	}
	dao, _ := New(redisconfig)
	success, err := dao.AddWord(origin, translate)
	if success != true {
		fmt.Println(err)
		t.Error("Test_AddWord failed")
	} else {
		fmt.Println("Test_AddWord passed")
	}
}

func Test_GetWordsRank(t *testing.T) {
	redisconfig := WordDaoConfig{
		Redis_network: "tcp",
		Redis_host:    "localhost:6379",
	}
	dao, _ := New(redisconfig)
	ranks, err := dao.GetWordsRank("rank", 0, -1, "withscores")
	if err != nil {
		t.Error("Test_GetWordsRank failed")
	} else {
		t.Log("Test_GetWordsRank passed")
	}
	fmt.Println(ranks)
	ranks, err = dao.GetWordsRank("2019-04-24", 0, -1, "withscores")
	if err != nil {
		t.Error("Test_GetWordsRank failed")
	} else {
		t.Log("Test_GetWordsRank passed")
	}
	fmt.Println(ranks)
}
