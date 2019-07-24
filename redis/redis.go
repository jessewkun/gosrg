package redis

import (
	"errors"
	"gosrg/utils"
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
	SEPARATOR      = " ==> "
	TYPE_STRING    = "string"
	TYPE_HASH      = "hash"
	TYPE_LIST      = "list"
	TYPE_SET       = "set"
	TYPE_ZSET      = "zset"
)

var IS_BOOT = false

type Redis struct {
	Host           string
	Port           string
	Pwd            string
	Conn           redis.Conn
	Db             int
	CurrentKey     string
	CurrentKeyType string
	Pattern        string
	Output         [][]string
	Keys           []string
	Detail         interface{}
	Info           [][]string
}

type CommandHandler func(content string) error

var commandMap map[string]CommandHandler

func InitRedis(host string, port string, pwd string, pattern string) error {
	conn, err := redis.Dial(REDIS_NETWORK, host+":"+port)
	if err != nil {
		if IS_BOOT {
			return err
		}
		utils.Exit(err)
	}

	if pwd != "" {
		if _, err := conn.Do("AUTH", pwd); err != nil {
			conn.Close()
			if IS_BOOT {
				return err
			}
			utils.Exit(err)
		}
	}
	utils.Info.Println("Redis conn ok")
	IS_BOOT = true
	R = &Redis{
		Host: host,
		Port: port,
		Pwd:  pwd,
		Conn: conn,
	}
	R.Pattern = pattern
	if len(pattern) == 0 {
		R.Pattern = "*"
	}
	R.Db = 0
	registerHandler()
	return nil
}

func registerHandler() {
	commandMap = map[string]CommandHandler{
		"select":  R.selectHandler,
		"keys":    R.keysHandler,
		"del":     R.delHandler,
		"info":    R.infoHandler,
		"type":    R.typeHandler,
		"object":  R.objectHandler,
		"ttl":     R.ttlHandler,
		"get":     R.getHandler,
		"set":     R.setHandler,
		"strlen":  R.strlenHandler,
		"hgetall": R.hgetallHandler,
		"hmset":   R.hmsetHandler,
		"hlen":    R.hlenHandler,
		"smember": R.smemberHandler,
		"scard":   R.scardHandler,
		"sadd":    R.saddHandler,
		"zcard":   R.zcardHandler,
		"zrange":  R.zrangeHandler,
		"zadd":    R.zaddHandler,
		"lrange":  R.lrangeHandler,
		"llen":    R.llenHandler,
		"rpush":   R.rpushHandler,
	}
}

func (r *Redis) Clear() {
	r.Output = [][]string{}
	r.Keys = []string{}
	r.Detail = nil
	r.Info = [][]string{}
	return
}

func (r *Redis) Exec(cmd string, content string) error {
	cmd = strings.ToLower(cmd)
	fun, ok := commandMap[cmd]
	if !ok {
		err := errors.New("redis cmd " + cmd + " handler is not existed")
		utils.Error.Println(err)
		return err
	}
	r.Clear()
	return fun(content)
}

func (r *Redis) Send(commandName string, args ...interface{}) {
	// keyType, err := redis.String(r.Conn.Do(commandName, args...))
}

func (r *Redis) GetKey(key string) error {
	r.Clear()
	if err := r.typeHandler(key); err != nil {
		return err
	}
	if err := r.objectHandler(key); err != nil {
		return err
	}
	if err := r.ttlHandler(key); err != nil {
		return err
	}

	switch r.CurrentKeyType {
	case TYPE_STRING:
		return r.getHandler(key)
	case TYPE_HASH:
		return r.hgetallHandler(key)
	case TYPE_SET:
		return r.smemberHandler(key)
	case TYPE_ZSET:
		return r.zrangeHandler(key)
	case TYPE_LIST:
		return r.lrangeHandler(key)
	}

	return nil
}

func (r *Redis) SetKey(content string) error {
	if r.CurrentKey == "" || r.CurrentKeyType == "" {
		return nil
	}
	r.Clear()
	content = utils.Trim(content)
	switch r.CurrentKeyType {
	case TYPE_STRING:
		return r.setHandler(content)
	case TYPE_HASH:
		return r.hmsetHandler(content)
	case TYPE_SET:
		return r.saddHandler(content)
	case TYPE_ZSET:
		return r.zaddHandler(content)
	case TYPE_LIST:
		return r.rpushHandler(content)
	}
	return nil
}
