package redis

import (
	"fmt"
	"reflect"
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
var redisPool *redis_go.Pool

func init() {
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置配置参数
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func SetConfig(host string, port int, password string) {
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	redisPool = newRedisPool(hostAddr, password, 15)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Run Command
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Do(commandName string, args ...interface{}) (interface{}, error) {
	redisClient := redisPool.Get()
	defer redisClient.Close()

	return redisClient.Do(commandName, args...)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 设置数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func SetData(structData interface{}, args ...interface{}) error {
	key := GetModelCacheKey(structData)
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
		if err := StringSet(key, jsonString, time); err != nil {
			return err
		}
	}

	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetData(structData interface{}, args ...interface{}) error {
	key := GetModelCacheKey(structData)

	if len(args) == 1 {
		key = string(args[0])
	}

	if jsonString, err := StringGet(key); err != nil {
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
func StringSet(args ...interface{}) error {
	if _, err := Do("SET", args...); err != nil {
		return err
	}
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * String GET 包装
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func StringGet(key string) (string, error) {
	resultString := ""
	if value, err := String(Do("GET", key)); err != nil {
		return "", err
	} else {
		resultString = value
	}

	return resultString, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Hash HMSET包装
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HashSet(structData interface{}, args ...interface{}) error {
	key := GetModelCacheKey(structData)
	var time int = 5 * 60

	if len(args) == 1 {
		key = string(args[0])
	}

	if len(args) == 2 {
		if timeValue, err := strconv.Atoi(str); err == nil {
			time = timeValue
		}
	}

	if _, err := Do("HMSET", redis_go.Args{}.Add(key).AddFlat(structData)...); err != nil {
		return err
	}
	return nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Hash HGETALL包装
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HashGet(structObject interface{}, args ...interface{}) error {
	key := GetModelCacheKey(structData)
	if len(args) == 1 {
		key = string(args[0])
	}

	data, err := redis_go.Values(Do("HGETALL", key))
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
func Bool(reply interface{}, err error) (bool, error) {
	return redis_go.Bool(reply, err)
}

func ByteSlices(reply interface{}, err error) ([][]byte, error) {
	return redis_go.ByteSlices(reply, err)
}

func Bytes(reply interface{}, err error) ([]byte, error) {
	return redis_go.Bytes(reply, err)
}

func Float64(reply interface{}, err error) (float64, error) {
	return redis_go.Float64(reply, err)
}

func Int(reply interface{}, err error) (int, error) {
	return redis_go.Int(reply, err)
}

func Int64(reply interface{}, err error) (int64, error) {
	return redis_go.Int64(reply, err)
}

func Int64Map(result interface{}, err error) (map[string]int64, error) {
	return redis_go.Int64Map(result, err)
}

func IntMap(result interface{}, err error) (map[string]int, error) {
	return redis_go.IntMap(result, err)
}

func Ints(reply interface{}, err error) ([]int, error) {
	return redis_go.Ints(reply, err)
}

func MultiBulk(reply interface{}, err error) ([]interface{}, error) {
	return redis_go.MultiBulk(reply, err)
}

func Scan(src []interface{}, dest ...interface{}) ([]interface{}, error) {
	return redis_go.Scan(src, dest...)
}

func ScanSlice(src []interface{}, dest interface{}, fieldNames ...string) error {
	return redis_go.ScanSlice(src, dest)
}

func ScanStruct(src []interface{}, dest interface{}) error {
	return redis_go.ScanStruct(src, dest)
}

func String(reply interface{}, err error) (string, error) {
	return redis_go.String(reply, err)
}

func StringMap(result interface{}, err error) (map[string]string, error) {
	return redis_go.StringMap(result, err)
}

func Strings(reply interface{}, err error) ([]string, error) {
	return redis_go.Strings(reply, err)
}

func Uint64(reply interface{}, err error) (uint64, error) {
	return redis_go.Uint64(reply, err)
}

func Values(reply interface{}, err error) ([]interface{}, error) {
	return redis_go.Values(reply, err)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取数据模型的缓存key
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetModelCacheKey(model interface{}) string {
	typeOf := reflect.TypeOf(model)
	valueOf := reflect.ValueOf(model)
	valueElem := valueOf.Elem()

	if kind := typeOf.Kind(); kind != reflect.Ptr {
		panic("Model is not a pointer type")
	}

	cacheKey := ""
	pkgName := strings.Split(valueElem.String(), " ")[0][1:]

	if _, ok := valueElem.Type().FieldByName("Id"); !ok {
		panic("Model does not contain Id field")
	} else {
		idValue := valueElem.FieldByName("Id").Uint()
		cacheKey = fmt.Sprintf("%s||%s||%d", common.Settings.Redis.PrefixKey, pkgName, idValue)
		cacheKey = strings.ToLower(cacheKey)
	}

	return cacheKey
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取 RedisPool
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func newRedisPool(address, password string, timeoutSeconds uint64) *redis_go.Pool {
	return &redis_go.Pool{
		MaxIdle:     8,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis_go.Conn, error) {
			return dial(address, password)
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
func dial(address, password string) (redis.Conn, error) {
	conn, err := redis_go.DialTimeout(
		"tcp",
		server,
		time.Duration(timeoutSeconds)*time.Second,
		time.Duration(timeoutSeconds)*time.Second,
		time.Duration(timeoutSeconds)*time.Second,
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
