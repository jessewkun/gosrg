package redis

import (
	"errors"
	"gosrg/utils"
	"strings"

	"github.com/gomodule/redigo/redis"
)

var R *Redis

const (
	RES_EXIT = iota
	RES_KEYS
	RES_DETAIL
	RES_INFO
	RES_OUTPUT_COMMAND
	RES_OUTPUT_INFO
	RES_OUTPUT_ERROR
	RES_OUTPUT_RES
)

const (
	REDIS_NETWORK = "tcp"

	SEPARATOR = " ==> "

	TYPE_STRING = "string"
	TYPE_HASH   = "hash"
	TYPE_LIST   = "list"
	TYPE_SET    = "set"
	TYPE_ZSET   = "zset"
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
	ResultChan     chan map[int]interface{}
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
	IS_BOOT = true
	if R == nil {
		R = &Redis{
			Host:       host,
			Port:       port,
			Pwd:        pwd,
			Conn:       conn,
			ResultChan: make(chan map[int]interface{}),
			Db:         0, // new connection reset db to 0
		}
	} else {
		R.Host = host
		R.Port = port
		R.Pwd = pwd
		R.Conn = conn
		R.Db = 0
	}
	R.Pattern = pattern
	if len(pattern) == 0 {
		R.Pattern = "*"
	}
	R.Send(RES_OUTPUT_INFO, "connect to "+host+":"+port+" success")
	registerHandler()
	return nil
}

func (r *Redis) Send(rtype int, data interface{}) {
	go func() {
		t := map[int]interface{}{rtype: data}
		r.ResultChan <- t
		return
	}()
}

func registerHandler() {
	commandMap = map[string]CommandHandler{
		"set":          R.setHandler,
		"get":          R.getHandler,
		"strlen":       R.strlenHandler,
		"hlen":         R.hlenHandler,
		"hmset":        R.hmsetHandler,
		"hgetall":      R.hgetallHandler,
		"rpush":        R.rpushHandler,
		"llen":         R.llenHandler,
		"lrange":       R.lrangeHandler,
		"sadd":         R.saddHandler,
		"scard":        R.scardHandler,
		"smembers":     R.smembersHandler,
		"zadd":         R.zaddHandler,
		"zcard":        R.zcardHandler,
		"zrange":       R.zrangeHandler,
		"exists":       R.existsHandler,
		"type":         R.typeHandler,
		"rename":       R.renameHandler,
		"renamenx":     R.renamenxHandler,
		"move":         R.moveHandler,
		"del":          R.delHandler,
		"randomkey":    R.randomkeyHandler,
		"dbsize":       R.dbsizeHandler,
		"keys":         R.keysHandler,
		"flushdb":      R.flushdbHandler,
		"flushall":     R.flushallHandler,
		"select":       R.selectHandler,
		"swapdb":       R.swapdbHandler,
		"ttl":          R.ttlHandler,
		"info":         R.infoHandler,
		"debug_object": R.debugObjectHandler,
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

func (r *Redis) CommandIsExisted(cmd string) (CommandHandler, error) {
	cmd = strings.ToLower(cmd)
	fun, ok := commandMap[cmd]
	if !ok {
		err := errors.New("unknown command `" + cmd + "`")
		r.Send(RES_OUTPUT_ERROR, err.Error)
		return nil, err
	}
	return fun, nil
}

func (r *Redis) Exec(cmd string, content string) error {
	fun, err := r.CommandIsExisted(cmd)
	if err != nil {
		return err
	}
	if mc, ok := multCommand[cmd]; ok {
		return r.mult(cmd, mc, content)
	}
	return fun(content)
}

func (r *Redis) mult(cmd string, mc []string, content string) error {
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

func (r *Redis) GetKey(key string) error {
	r.CurrentKey = key
	if err := r.typeHandler(""); err != nil {
		return err
	}
	if err := r.debugObjectHandler(""); err != nil {
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
		if err := r.smembersHandler(""); err != nil {
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
