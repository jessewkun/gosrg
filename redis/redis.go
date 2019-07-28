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
var multCommand map[string][]string

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
		"select":    R.selectHandler,
		"keys":      R.keysHandler,
		"del":       R.delHandler,
		"info":      R.infoHandler,
		"type":      R.typeHandler,
		"dbsize":    R.dbsizeHandler,
		"object":    R.objectHandler,
		"ttl":       R.ttlHandler,
		"get":       R.getHandler,
		"set":       R.setHandler,
		"strlen":    R.strlenHandler,
		"hgetall":   R.hgetallHandler,
		"hmset":     R.hmsetHandler,
		"hlen":      R.hlenHandler,
		"smember":   R.smemberHandler,
		"scard":     R.scardHandler,
		"sadd":      R.saddHandler,
		"zcard":     R.zcardHandler,
		"zrange":    R.zrangeHandler,
		"zadd":      R.zaddHandler,
		"lrange":    R.lrangeHandler,
		"llen":      R.llenHandler,
		"rpush":     R.rpushHandler,
		"exists":    R.existsHandler,
		"rename":    R.renameHandler,
		"renamenx":  R.renamenxHandler,
		"move":      R.moveHandler,
		"randomkey": R.randomkeyHandler,
		"flushdb":   R.flushdbHandler,
		"flushall":  R.flushallHandler,
		"swapdb":    R.swapdbHandler,
	}
	multCommand = map[string][]string{
		"info":     []string{"dbsize"},
		"rename":   []string{"keys"},
		"renamenx": []string{"keys"},
		"move":     []string{"keys"},
		"flushdb":  []string{"keys"},
		"flushall": []string{"keys"},
		"swapdb":   []string{"keys"},
	}
}

func (r *Redis) Clear() {
	r.Output = [][]string{}
	r.Keys = []string{}
	r.Detail = nil
	r.Info = [][]string{}
	return
}

func (r *Redis) CommandIsExisted(cmd string) (CommandHandler, error) {
	cmd = strings.ToLower(cmd)
	fun, ok := commandMap[cmd]
	if !ok {
		err := errors.New("unknown command `" + cmd + "`")
		utils.Error.Println(err)
		return nil, err
	}
	return fun, nil
}

func (r *Redis) Exec(cmd string, content string) error {
	fun, err := r.CommandIsExisted(cmd)
	if err != nil {
		return err
	}
	r.Clear()
	if mc, ok := multCommand[cmd]; ok {
		if err := commandMap[cmd](content); err != nil {
			return err
		}
		for _, item := range mc {
			if err := commandMap[item](content); err != nil {
				return err
			}
		}
		return nil
	}
	return fun(content)
}

func (r *Redis) multInfo() error {
	r.Clear()
	if err := r.infoHandler(""); err != nil {
		return err
	}
	if err := r.dbsizeHandler(""); err != nil {
		return err
	}
	return nil
}

func (r *Redis) multRename(content string) error {
	r.Clear()
	if err := r.renameHandler(content); err != nil {
		return err
	}
	if err := r.keysHandler(""); err != nil {
		return err
	}
	return nil
}

func (r *Redis) GetKey(key string) error {
	r.Clear()
	r.CurrentKey = key
	if err := r.typeHandler(""); err != nil {
		return err
	}
	if err := r.objectHandler(""); err != nil {
		return err
	}
	if err := r.ttlHandler(""); err != nil {
		return err
	}

	switch r.CurrentKeyType {
	case TYPE_STRING:
		if err := r.getHandler(""); err != nil {
			return err
		}
		if err := r.strlenHandler(""); err != nil {
			return err
		}
	case TYPE_HASH:
		if err := r.hgetallHandler(""); err != nil {
			return err
		}
		if err := r.hlenHandler(""); err != nil {
			return err
		}
	case TYPE_SET:
		if err := r.smemberHandler(""); err != nil {
			return err
		}
		if err := r.scardHandler(""); err != nil {
			return err
		}
	case TYPE_ZSET:
		if err := r.zrangeHandler(""); err != nil {
			return err
		}
		if err := r.zcardHandler(""); err != nil {
			return err
		}
	case TYPE_LIST:
		if err := r.lrangeHandler(""); err != nil {
			return err
		}
		if err := r.llenHandler(""); err != nil {
			return err
		}
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
