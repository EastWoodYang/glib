package redis

import (
	"fmt"
	"strings"
	"time"
)

import (
	redis_go "github.com/garyburd/redigo/redis"
)

import (
	"github.com/sanxia/glib"
)

/* ================================================================================
 * Redis操作模块
 * author: mliu
 * ================================================================================ */
type (
	redisClient struct {
		prefixKey string
		pool      *redis_go.Pool
	}
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取RedisClient实例
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewRedisClient(host string, port int, password string, timeout uint64) *redisClient {
	client := new(redisClient)
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	client.pool = newRedisPool(hostAddr, password, timeout)

	return client
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Run Command
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *redisClient) SetPrefixKey(prefixKey string) {
	s.prefixKey = prefixKey
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Run Command
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *redisClient) Do(commandName string, args ...interface{}) (interface{}, error) {
	redisPool := s.pool.Get()
	defer redisPool.Close()

	return redisPool.Do(commandName, args...)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *redisClient) SetData(structData interface{}, args ...interface{}) error {
	key := s.getModelCacheKey(structData)
	var time int = 5 * 60

	if len(args) == 1 {
		key = string(args[0])
	}

	if len(args) == 2 {
		if timeValue, err := strconv.Atoi(str); err == nil {
			time = timeValue
		}
	}

	if jsonString, err := glib.ToJson(structData); err != nil {
		return err
	} else if len(jsonString) > 0 {
		if err := s.StringSet(key, jsonString, time); err != nil {
			return err
		}
	}

	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *redisClient) GetData(structData interface{}, args ...interface{}) error {
	key := s.getModelCacheKey(structData)

	if len(args) == 1 {
		key = string(args[0])
	}

	if jsonString, err := s.StringGet(key); err != nil {
		return err
	} else {
		if err := glib.FromJson(jsonString, &structData); err != nil {
			return err
		}
	}
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * String SET 包装
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *redisClient) StringSet(args ...interface{}) error {
	if _, err := s.Do("SET", args...); err != nil {
		return err
	}
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * String GET 包装
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *redisClient) StringGet(key string) (string, error) {
	resultString := ""
	if value, err := s.String(s.Do("GET", key)); err != nil {
		return "", err
	} else {
		resultString = value
	}

	return resultString, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Hash HMSET包装
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *redisClient) HashSet(structData interface{}, args ...interface{}) error {
	key := s.getModelCacheKey(structData)
	var time int = 5 * 60

	if len(args) == 1 {
		key = string(args[0])
	}

	if len(args) == 2 {
		if timeValue, err := strconv.Atoi(str); err == nil {
			time = timeValue
		}
	}

	if _, err := s.Do("HMSET", redis_go.Args{}.Add(key).AddFlat(structData)...); err != nil {
		return err
	}
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Hash HGETALL包装
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *redisClient) HashGet(structObject interface{}, args ...interface{}) error {
	key := s.getModelCacheKey(structData)
	if len(args) == 1 {
		key = string(args[0])
	}

	data, err := redis_go.Values(s.Do("HGETALL", key))
	if err != nil {
		return err
	}

	if err := redis_go.ScanStruct(data, structObject); err != nil {
		return err
	}

	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * redigo帮助方法包装
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *redisClient) Bool(reply interface{}, err error) (bool, error) {
	return redis_go.Bool(reply, err)
}

func (s *redisClient) ByteSlices(reply interface{}, err error) ([][]byte, error) {
	return redis_go.ByteSlices(reply, err)
}

func (s *redisClient) Bytes(reply interface{}, err error) ([]byte, error) {
	return redis_go.Bytes(reply, err)
}

func (s *redisClient) Float64(reply interface{}, err error) (float64, error) {
	return redis_go.Float64(reply, err)
}

func (s *redisClient) Int(reply interface{}, err error) (int, error) {
	return redis_go.Int(reply, err)
}

func (s *redisClient) Int64(reply interface{}, err error) (int64, error) {
	return redis_go.Int64(reply, err)
}

func (s *redisClient) Int64Map(result interface{}, err error) (map[string]int64, error) {
	return redis_go.Int64Map(result, err)
}

func (s *redisClient) IntMap(result interface{}, err error) (map[string]int, error) {
	return redis_go.IntMap(result, err)
}

func (s *redisClient) Ints(reply interface{}, err error) ([]int, error) {
	return redis_go.Ints(reply, err)
}

func (s *redisClient) MultiBulk(reply interface{}, err error) ([]interface{}, error) {
	return redis_go.MultiBulk(reply, err)
}

func (s *redisClient) Scan(src []interface{}, dest ...interface{}) ([]interface{}, error) {
	return redis_go.Scan(src, dest...)
}

func (s *redisClient) ScanSlice(src []interface{}, dest interface{}, fieldNames ...string) error {
	return redis_go.ScanSlice(src, dest)
}

func (s *redisClient) ScanStruct(src []interface{}, dest interface{}) error {
	return redis_go.ScanStruct(src, dest)
}

func (s *redisClient) String(reply interface{}, err error) (string, error) {
	return redis_go.String(reply, err)
}

func (s *redisClient) StringMap(result interface{}, err error) (map[string]string, error) {
	return redis_go.StringMap(result, err)
}

func (s *redisClient) Strings(reply interface{}, err error) ([]string, error) {
	return redis_go.Strings(reply, err)
}

func (s *redisClient) Uint64(reply interface{}, err error) (uint64, error) {
	return redis_go.Uint64(reply, err)
}

func (s *redisClient) Values(reply interface{}, err error) ([]interface{}, error) {
	return redis_go.Values(reply, err)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 链接 Redis 服务器
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *redisClient) getModelCacheKey(model interface{}) string {
	return glib.GetModelKey(model, s.prefixKey, "Id")
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取 RedisPool
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func newRedisPool(address, password string, timeout uint64) *redis_go.Pool {
	return &redis_go.Pool{
		MaxIdle:     8,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis_go.Conn, error) {
			return dial(address, password, timeout)
		},
		TestOnBorrow: func(conn redis_go.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 链接 Redis 服务器
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func dial(address, password string, timeout uint64) (redis.Conn, error) {
	conn, err := redis_go.DialTimeout(
		"tcp",
		server,
		time.Duration(timeout)*time.Second,
		time.Duration(timeout)*time.Second,
		time.Duration(timeout)*time.Second,
	)
	if err != nil {
		return nil, err
	}
	if len(password) > 0 {
		if _, err := conn.Do("AUTH", password); err != nil {
			conn.Close()
			return nil, err
		}
	}

	return conn, err
}
