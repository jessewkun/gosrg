package redis

import (
	"gosrg/config"
	"gosrg/utils"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func InitConn() (redis.Conn, error) {
	conn, err := redis.Dial(config.REDIS_NETWORK, config.Srg.Host+":"+config.Srg.Port)
	if err != nil {
		utils.Logger.Panicln(err)
		return nil, err
	}
	return conn, nil
}

func Db() {
	_, err := redis.String(config.Srg.Redis.Do("select", config.Srg.Db))
	if err != nil {
		utils.Logger.Fatalln(err)
		utils.OErrorOuput(err.Error())
	} else {
		config.Srg.CurrentKey = ""
		config.Srg.CurrentKeyType = ""
		utils.OCommandOuput("select " + strconv.Itoa(config.Srg.Db))
	}
}

func Keys() {
	keys, err := redis.Strings(config.Srg.Redis.Do("keys", "*"))
	if err != nil {
		utils.Logger.Fatalln(err)
	} else {
		config.Srg.CurrentKey = ""
		config.Srg.CurrentKeyType = ""
		utils.OCommandOuput("keys *")
		for _, key := range keys {
			utils.Kouput(key)
		}
	}
}

func Info() {
	command := "info"
	info, err := redis.String(config.Srg.Redis.Do(command))
	if err != nil {
		utils.Logger.Fatalln(err)
	} else {
		config.Srg.CurrentKey = ""
		config.Srg.CurrentKeyType = ""
		utils.OCommandOuput(command)
		utils.Douput(info)
	}
}

func KeyDetail(key string) {
	keyType, err := redis.String(config.Srg.Redis.Do("type", key))
	if err != nil {
		utils.Logger.Fatalln(err)
	} else {
		config.Srg.CurrentKey = key
		config.Srg.CurrentKeyType = keyType
		utils.OCommandOuput("type " + key)
		utils.OInfoOuput(keyType)
		switch keyType {
		case "string":
			getString(key)
		case "hash":
			getHash(key)
		case "set":
			getSet(key)
		case "zset":
			getZset(key)
		case "list":
			getList(key)
		}
	}
}

func getString(key string) {
	res, err := redis.String(config.Srg.Redis.Do("get", key))
	if err != nil {
		utils.Logger.Fatalln(err)
	} else {
		utils.OCommandOuput("get " + key)
		utils.Douput(res)
	}
}

func getHash(key string) {
}

func getSet(key string) {
}

func getZset(key string) {
}

func getList(key string) {
}

func SetKeyDetail(content string) error {
	if config.Srg.CurrentKey == "" || config.Srg.CurrentKeyType == "" {
		return nil
	}
	switch config.Srg.CurrentKeyType {
	case "string":
		setString(content)
	case "hash":
		setHash(content)
	case "set":
		setSet(content)
	case "zset":
		setZset(content)
	case "list":
		setList(content)
	}
	return nil
}

func setString(content string) {
	content = strings.Trim(content, " ")
	content = strings.Trim(content, "\n")
	res, err := redis.String(config.Srg.Redis.Do("set", config.Srg.CurrentKey, content))
	if err != nil {
		utils.OErrorOuput(err.Error())
		utils.Logger.Fatalln(err)
	} else {
		utils.OCommandOuput("set " + config.Srg.CurrentKey + " " + content)
		utils.OInfoOuput(res)
	}
}
func setHash(content string) {}
func setSet(content string)  {}
func setZset(content string) {}
func setList(content string) {}
