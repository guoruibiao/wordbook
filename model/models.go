package model

/**
 *  Redis-Server的暂存信息
 **/
type RedisConfigs struct {
	Network string `json:"redis_network"`
	Host    string `json:"redis_host"`
}
