package redis

import (
	"gosrg/utils"
	"os"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

var R *Redis

const (
	REDIS_NETWORK  = "tcp"
	OUTPUT_COMMAND = "c"
	OUTPUT_INFO    = "i"
	OUTPUT_ERROR   = "e"
	OUTPUT_RES     = "r"
	OUTPUT_DEBUG   = "d"
)

type Redis struct {
	Host           string
	Port           string
	Pwd            string
	Redis          redis.Conn
	Db             int
	CurrentKey     string
	CurrentKeyType string
	Pattern        string
}

func InitRedis() {
	R = &Redis{
		Host:    "127.0.0.1",
		Port:    "6379",
		Pattern: "*",
	}
	conn, err := redis.Dial(REDIS_NETWORK, R.Host+":"+R.Port)
	if err != nil {
		println("redis connect fail")
		utils.Logger.Println(err.Error())
		os.Exit(1)
	}
	R.Redis = conn
}

func (R *Redis) SelectDb(db int) (output [][]string) {
	output = append(output, []string{"SELECT " + strconv.Itoa(db), OUTPUT_COMMAND})
	_, err := redis.String(R.Redis.Do("SELECT", db))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	R.CurrentKey = ""
	R.CurrentKeyType = ""
	return
}

func (R *Redis) Keys() (output [][]string, keys []string) {
	output = append(output, []string{"KEYS " + R.Pattern, OUTPUT_COMMAND})
	keys, err := redis.Strings(R.Redis.Do("KEYS", R.Pattern))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	R.CurrentKey = ""
	R.CurrentKeyType = ""
	return
}

func (R *Redis) Del() (output [][]string) {
	output = append(output, []string{"DEL " + R.CurrentKey, OUTPUT_COMMAND})
	res, err := redis.Int64(R.Redis.Do("DEL", R.CurrentKey))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	R.CurrentKey = ""
	R.CurrentKeyType = ""
	output = append(output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return
}

func (R *Redis) Info() (output [][]string, info string) {
	output = append(output, []string{"INFO", OUTPUT_COMMAND})
	info, err := redis.String(R.Redis.Do("INFO"))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	R.CurrentKey = ""
	R.CurrentKeyType = ""
	return
}

func (R *Redis) Send(ommandName string, args ...interface{}) {
	// keyType, err := redis.String(R.Redis.Do(ommandName, args...))
}

func (R *Redis) KeyDetail(key string) (output [][]string, res interface{}, info [][]string) {
	output = append(output, []string{"TYPE " + key, OUTPUT_COMMAND})

	keyType, err := redis.String(R.Redis.Do("TYPE", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Println(err)
		return
	}
	R.CurrentKey = key
	R.CurrentKeyType = keyType
	info = append(info, []string{"type", keyType})

	output = append(output, []string{"DEBUG OBJECT " + key, OUTPUT_COMMAND})
	object, err := redis.String(R.Redis.Do("DEBUG", "OBJECT", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Println(err)
		return
	}
	objectArr := strings.Split(object, " ")
	for _, item := range objectArr {
		tmp := strings.Split(item, ":")
		if len(tmp) == 1 {
			continue // Value
		}
		if tmp[0] == "at" {
			info = append(info, []string{"Value at", tmp[1]})
		} else {
			info = append(info, []string{tmp[0], tmp[1]})
		}
	}

	output = append(output, []string{"TTL " + key, OUTPUT_COMMAND})
	ttlres, err := redis.Int64(R.Redis.Do("TTL", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	info = append(info, []string{"ttl", strconv.FormatInt(ttlres, 10)})

	switch keyType {
	case "string":
		o, detail, stringinfo := getString(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	case "hash":
		o, detail, stringinfo := getHash(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	case "set":
		o, detail, stringinfo := getSet(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	case "zset":
		o, detail, stringinfo := getZset(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	case "list":
		o, detail, stringinfo := getList(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	}

	return
}

func getString(key string) (output [][]string, res string, info [][]string) {
	output = append(output, []string{"GET " + key, OUTPUT_COMMAND})
	res, err := redis.String(R.Redis.Do("GET", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}

	output = append(output, []string{"STRLEN " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(R.Redis.Do("STRLEN", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	info = append(info, []string{"len", strconv.FormatInt(lenres, 10)})
	return
}

func getHash(key string) (output [][]string, res interface{}, info [][]string) {
	output = append(output, []string{"HGETALl " + key, OUTPUT_COMMAND})
	res, err := redis.StringMap(R.Redis.Do("HGETALL", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}

	output = append(output, []string{"HLEN " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(R.Redis.Do("HLEN", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	info = append(info, []string{"hlen", strconv.FormatInt(lenres, 10)})
	return
}

func getSet(key string) (output [][]string, res []string, info [][]string) {
	output = append(output, []string{"SMEMBERS " + key, OUTPUT_COMMAND})
	res, err := redis.Strings(R.Redis.Do("SMEMBERS", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}

	output = append(output, []string{"SCARD " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(R.Redis.Do("SCARD", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	info = append(info, []string{"scard", strconv.FormatInt(lenres, 10)})
	return
}

func getZset(key string) (output [][]string, res interface{}, info [][]string) {
	output = append(output, []string{"ZRANGE " + key + " 0 -1 WITHSCORES", OUTPUT_COMMAND})
	res, err := redis.StringMap(R.Redis.Do("ZRANGE", key, 0, -1, "WITHSCORES"))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}

	output = append(output, []string{"ZCARD " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(R.Redis.Do("ZCARD", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	info = append(info, []string{"zcard", strconv.FormatInt(lenres, 10)})
	return
}

func getList(key string) (output [][]string, res []string, info [][]string) {
	output = append(output, []string{"LRANGE " + key + " 0 -1", OUTPUT_COMMAND})
	res, err := redis.Strings(R.Redis.Do("LRANGE", key, 0, -1))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}

	output = append(output, []string{"LLEN " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(R.Redis.Do("LLEN", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	info = append(info, []string{"llen", strconv.FormatInt(lenres, 10)})
	return
}

func (R *Redis) SetKeyDetail(content string) (output [][]string) {
	if R.CurrentKey == "" || R.CurrentKeyType == "" {
		return
	}
	switch R.CurrentKeyType {
	case "string":
		output = setString(content)
	case "hash":
		setHash(content)
	case "set":
		setSet(content)
	case "zset":
		setZset(content)
	case "list":
		setList(content)
	}
	return
}

func setString(content string) (output [][]string) {
	content = utils.Trim(content)
	output = append(output, []string{"SET " + R.CurrentKey + " " + content, OUTPUT_COMMAND})
	res, err := redis.String(R.Redis.Do("set", R.CurrentKey, content))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	output = append(output, []string{res, OUTPUT_RES})
	return
}

func setHash(content string) {}
func setSet(content string)  {}
func setZset(content string) {}
func setList(content string) {}
