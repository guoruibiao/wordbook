# 记录生僻的单词，以备后续复习

一个做了一半的单词本 terminal version

`/tmp/wordsrc` runtime configuration 
```
{
  "key": "value",
  "redis_network": "tcp",
  "redis_host": "localhost:6379"
}
```

感觉这里用到的一个经典的概念就是`callback`，原因还是在与Redis链接对象释放时机的问题。

```golang
// 定义一个callback对象
type DoRedisNoReturn func(client redis.Conn, command string, args ...interface{}) (bool, error)

func DoRedisNoReturnImpl(client redis.Conn, command string, args ...interface{}) (bool, error) {
	_, err := client.Do(command, args...)
	if err != nil {
		return false, err
	}
	return true, nil
}
```

```golang
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


```

```golang
// 应用案例
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
```

golang大法好，尤其是单元测试`_test`的引入，回归起来简直太好用了。