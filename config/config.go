package config

const (
	DAY_KEY   = "daykey:zset"             // 每日单词统计 hashcode => added|updated timestamp
	ALL_WORDS = "allwords:zset"           // 单词本 hashcode => added|updated timestamp
	TAG_STORAGE = "tagstorage:hash"       // 每一个单词的含义 hashcode => origin@translate
	APPERENCE_RANK = "apperencerank:zset" // 出现频率频次榜 hashcode => numbers
)

