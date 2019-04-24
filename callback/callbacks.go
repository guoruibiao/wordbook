package callback

import (
	"github.com/garyburd/redigo/redis"
)

type DoRedisNoReturn func(client redis.Conn, command string, args ...interface{}) (bool, error)

func DoRedisNoReturnImpl(client redis.Conn, command string, args ...interface{}) (bool, error) {
	_, err := client.Do(command, args...)
	if err != nil {
		return false, err
	}
	return true, nil
}
