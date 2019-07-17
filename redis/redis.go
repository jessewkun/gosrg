package redis

import (
	"gosrg/utils"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

var R *Redis

const (
	REDIS_NETWORK  = "tcp"
	OUTPUT_COMMAND = "c"
	OUTPUT_INFO    = "i"
	OUTPUT_ERROR   = "e"
)

type Redis struct {
	Host           string
	Port           string
	Pwd            string
	Redis          redis.Conn
	Db             int
	CurrentKey     string
	CurrentKeyType string
}

func InitRedis() {
	R = &Redis{
		Host: "127.0.0.1",
		Port: "6379",
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
	output = append(output, []string{"select " + strconv.Itoa(db), OUTPUT_COMMAND})
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

func (R *Redis) Keys(pattern string) (output [][]string, keys []string) {
	output = append(output, []string{"keys " + pattern, OUTPUT_COMMAND})
	keys, err := redis.Strings(R.Redis.Do("keys", pattern))
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
	output = append(output, []string{"del " + R.CurrentKey, OUTPUT_COMMAND})
	res, err := redis.Int64(R.Redis.Do("del", R.CurrentKey))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	R.CurrentKey = ""
	R.CurrentKeyType = ""
	output = append(output, []string{strconv.FormatInt(res, 10), OUTPUT_INFO})
	return
}

func (R *Redis) Info() (output [][]string, info string) {
	output = append(output, []string{"info", OUTPUT_COMMAND})
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

func (R *Redis) KeyDetail(key string) (output [][]string, res interface{}) {
	output = append(output, []string{"type " + key, OUTPUT_COMMAND})
	keyType, err := redis.String(R.Redis.Do("type", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	R.CurrentKey = key
	R.CurrentKeyType = keyType
	output = append(output, []string{keyType, OUTPUT_INFO})
	switch keyType {
	case "string":
		o, detail := getString(key)
		output = append(output, o...)
		res = detail
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

func getString(key string) (output [][]string, res string) {
	output = append(output, []string{"get " + key, OUTPUT_COMMAND})
	res, err := redis.String(R.Redis.Do("get", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
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
	output = append(output, []string{"set " + R.CurrentKey + " " + content, OUTPUT_COMMAND})
	res, err := redis.String(R.Redis.Do("set", R.CurrentKey, content))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		utils.Logger.Fatalln(err)
		return
	}
	output = append(output, []string{res, OUTPUT_INFO})
	return
}

func setHash(content string) {}
func setSet(content string)  {}
func setZset(content string) {}
func setList(content string) {}
