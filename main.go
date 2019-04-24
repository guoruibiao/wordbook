package main

import (
	"fmt"
	"words/config"
	"words/dao"
	"words/utils"
)

func main() {
	configurations, err := utils.LoadConfigurations("/tmp/wordsrc")
	if err != nil {
		fmt.Println("加载配置文件失败" + err.Error())
	}
	fmt.Println(configurations)
	wordDaoConfigs := dao.WordDaoConfig{
		Redis_network: configurations["redis_network"],
		Redis_host:    configurations["redis_host"],
	}
	worddao, err := dao.New(wordDaoConfigs)
	if err != nil {
		fmt.Println("初始化worddao失败" + err.Error())
	}
	origin := "demonstration"
	translate := "示例"
	success, err := worddao.AddWord(origin, translate)
	if success != true {
		fmt.Println("添加单词失败" + err.Error())
	}
	// 显示总词频
	rank, err := worddao.GetWordsRank(config.APPERENCE_RANK, 0, -1, "withscores")
	if err != nil {
		fmt.Println("获取总词频失败" + err.Error())
	} else {
		fmt.Println("获取词频总结果")
		fmt.Println(rank)
	}
	// 显示总词库
	//words = worddao.GetWordsRank
	// 显示当日词库
	// TODO 命令行下的菜单形式
	fmt.Println("hello world.")
}
