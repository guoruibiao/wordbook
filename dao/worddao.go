package dao

import (
	"time"
	"words/callback"
	"words/config"
	"words/utils"

	"github.com/garyburd/redigo/redis"
)

/**
 * 针对word应用，提供底层的
 * 1. 添加新单词
 * 2. 获取每日列表，总列表，频次图等等
 **/

type WordDaoConfig struct {
	Redis_network string
	Redis_host    string
}

type WordDao struct {
	redisconfig WordDaoConfig
}

// 思考下如何加入最外层的静态配置
func New(wordconfig WordDaoConfig) (dao *WordDao, err error) {
	dao = &WordDao{
		redisconfig: wordconfig,
	}
	return dao, nil
}

// 获取key对应的单词列表，key会是一个日期相关的字符串参数，具体定义在words/config/config.go
func (dao *WordDao) GetMembers(key string) (members []string, err error) {
	client, err := redis.Dial(dao.redisconfig.Redis_network, dao.redisconfig.Redis_host)
	defer client.Close()
	if err != nil {
		return nil, err
	}
	content, err := redis.Values(client.Do("smembers", key))
	if err != nil {
		return nil, err
	}
	for _, item := range content {
		members = append(members, string(item.([]byte)))
	}
	return members, nil
}

// 添加新单词到系统
func (dao *WordDao) AddWord(origin string, translate string) (bool, error) {
	/** MD 我咋会想到用这玩意儿？？？hashcode了貌似也没啥用...
	hashcode, err := utils.GeneralHashCode(origin)
	if err != nil {
		return false, err
	}
	**/
	// 判断 是否已在词库中|不过感觉覆盖式查询更好用
	params := []interface{}{config.TAG_STORAGE, origin, translate}
	_, err := utils.RedisRun(dao.redisconfig.Redis_network, dao.redisconfig.Redis_host, callback.DoRedisNoReturnImpl, "hset", params...)
	if err != nil {
		return false, err
	}
	// 加到单词本上: timestamp 需要是一个合法的int或者float
	params = []interface{}{config.ALL_WORDS, int(time.Now().Unix()), origin}
	_, err = utils.RedisRun(dao.redisconfig.Redis_network, dao.redisconfig.Redis_host, callback.DoRedisNoReturnImpl, "zadd", params...)
	if err != nil {
		return false, err
	}
	// 加到当日单词本上
	params = []interface{}{config.DAY_KEY + utils.GetCurDate(), int(time.Now().Unix()), origin}
	_, err = utils.RedisRun(dao.redisconfig.Redis_network, dao.redisconfig.Redis_host, callback.DoRedisNoReturnImpl, "zadd", params...)
	if err != nil {
		return false, err
	}
	// 增加单词访问频率
	params = []interface{}{config.APPERENCE_RANK, 1, origin}
	_, err = utils.RedisRun(dao.redisconfig.Redis_network, dao.redisconfig.Redis_host, callback.DoRedisNoReturnImpl, "zincrby", params...)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 获取单词本词频列表,根据key来确定返回总的词频表还是具体某一天的词频表
func (dao *WordDao) GetWordsRank(key string, start, nums int, withscore string) (map[string]string, error) {
	client, err := redis.Dial(dao.redisconfig.Redis_network, dao.redisconfig.Redis_host)
	defer client.Close()
	if err != nil {
		return nil, err
	}
	content, err := redis.StringMap(client.Do("zrevrange", key, start, nums, withscore))
	if err != nil {
		return nil, err
	}

	return content, nil
}

// 导入已有系统的日记本
func Import(filepath string) bool {
	// TODO 规定好格式
	return false
}

// 导出当前系统的日记本
func Export(filepath string) bool {
	return false
}
