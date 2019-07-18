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
	_, err := redis.String(R.Redis.Do("select", db))
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
	keys, err := redis.Strings(R.Redis.Do("keys", R.Pattern))
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
	res, err := redis.Int64(R.Redis.Do("del", R.CurrentKey))
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
	info, err := redis.String(R.Redis.Do("info"))
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

	keyType, err := redis.String(R.Redis.Do("type", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Println(err)
		return
	}
	R.CurrentKey = key
	R.CurrentKeyType = keyType
	info = append(info, []string{"type", keyType})

	output = append(output, []string{"DEBUG OBJECT " + key, OUTPUT_COMMAND})
	object, err := redis.String(R.Redis.Do("debug", "object", key))
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
	ttlres, err := redis.Int64(R.Redis.Do("ttl", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	info = append(info, []string{"ttl", strconv.FormatInt(ttlres, 10)})

	// output = append(output, []string{keyType, OUTPUT_RES})
	switch keyType {
	case "string":
		o, detail, stringinfo := getString(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	case "hash":
		getHash(key)
	case "set":
		getSet(key)
	case "zset":
		getZset(key)
	case "list":
		getList(key)
	}

	return
}

func getString(key string) (output [][]string, res string, info [][]string) {
	output = append(output, []string{"GET " + key, OUTPUT_COMMAND})
	res, err := redis.String(R.Redis.Do("get", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}

	output = append(output, []string{"STRLEN " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(R.Redis.Do("strlen", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	info = append(info, []string{"len", strconv.FormatInt(lenres, 10)})
	return
}

func getHash(key string) {
}

func getSet(key string) {
}

func getZset(key string) {
}

func getList(key string) {
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
