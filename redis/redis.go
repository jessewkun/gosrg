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
		"set":               R.setHandler,
		"setnx":             R.setnxHandler,
		"setex":             R.setexHandler,
		"psetex":            R.psetexHandler,
		"get":               R.getHandler,
		"getset":            R.getsetHandler,
		"strlen":            R.strlenHandler,
		"append":            R.appendHandler,
		"setrange":          R.setrangeHandler,
		"getrange":          R.getrangeHandler,
		"incr":              R.incrHandler,
		"incrby":            R.incrbyHandler,
		"incrbyfloat":       R.incrbyfloatHandler,
		"decr":              R.decrHandler,
		"decrby":            R.decrbyHandler,
		"mset":              R.msetHandler,
		"msetnx":            R.msetnxHandler,
		"mget":              R.mgetHandler,
		"hset":              R.hsetHandler,
		"hsetnx":            R.hsetnxHandler,
		"hget":              R.hgetHandler,
		"hexists":           R.hexistsHandler,
		"hdel":              R.hdelHandler,
		"hlen":              R.hlenHandler,
		"hstrlen":           R.hstrlenHandler,
		"hincrby":           R.hincrbyHandler,
		"hincrbyfloat":      R.hincrbyfloatHandler,
		"hmset":             R.hmsetHandler,
		"hmget":             R.hmgetHandler,
		"hkeys":             R.hkeysHandler,
		"hvals":             R.hvalsHandler,
		"hgetall":           R.hgetallHandler,
		"hscan":             R.hscanHandler,
		"lpush":             R.lpushHandler,
		"lpushx":            R.lpushxHandler,
		"rpush":             R.rpushHandler,
		"rpushx":            R.rpushxHandler,
		"lpop":              R.lpopHandler,
		"rpop":              R.rpopHandler,
		"rpoplpush":         R.rpoplpushHandler,
		"lrem":              R.lremHandler,
		"llen":              R.llenHandler,
		"lindex":            R.lindexHandler,
		"linsert":           R.linsertHandler,
		"lset":              R.lsetHandler,
		"lrange":            R.lrangeHandler,
		"ltrim":             R.ltrimHandler,
		"blpop":             R.blpopHandler,
		"brpop":             R.brpopHandler,
		"brpoplpush":        R.brpoplpushHandler,
		"sadd":              R.saddHandler,
		"sismember":         R.sismemberHandler,
		"spop":              R.spopHandler,
		"srandmember":       R.srandmemberHandler,
		"srem":              R.sremHandler,
		"smove":             R.smoveHandler,
		"scard":             R.scardHandler,
		"smembers":          R.smembersHandler,
		"sscan":             R.sscanHandler,
		"sinter":            R.sinterHandler,
		"sinterstore":       R.sinterstoreHandler,
		"sunion":            R.sunionHandler,
		"sunionstore":       R.sunionstoreHandler,
		"sdiff":             R.sdiffHandler,
		"sdiffstore":        R.sdiffstoreHandler,
		"zadd":              R.zaddHandler,
		"zscore":            R.zscoreHandler,
		"zincrby":           R.zincrbyHandler,
		"zcard":             R.zcardHandler,
		"zcount":            R.zcountHandler,
		"zrange":            R.zrangeHandler,
		"zrevrange":         R.zrevrangeHandler,
		"zrangebyscore":     R.zrangebyscoreHandler,
		"zrevrangebyscore":  R.zrevrangebyscoreHandler,
		"zrank":             R.zrankHandler,
		"zrevrank":          R.zrevrankHandler,
		"zrem":              R.zremHandler,
		"zremrangebyrank":   R.zremrangebyrankHandler,
		"zremrangebyscore":  R.zremrangebyscoreHandler,
		"zrangebylex":       R.zrangebylexHandler,
		"zlexcount":         R.zlexcountHandler,
		"zremrangebylex":    R.zremrangebylexHandler,
		"zscan":             R.zscanHandler,
		"zunionstore":       R.zunionstoreHandler,
		"zinterstore":       R.zinterstoreHandler,
		"pfadd":             R.pfaddHandler,
		"pfcount":           R.pfcountHandler,
		"pfmerge":           R.pfmergeHandler,
		"geoadd":            R.geoaddHandler,
		"geopos":            R.geoposHandler,
		"geodist":           R.geodistHandler,
		"georadius":         R.georadiusHandler,
		"georadiusbymember": R.georadiusbymemberHandler,
		"geohash":           R.geohashHandler,
		"setbit":            R.setbitHandler,
		"getbit":            R.getbitHandler,
		"bitcount":          R.bitcountHandler,
		"bitpos":            R.bitposHandler,
		"bitop":             R.bitopHandler,
		"bitfield":          R.bitfieldHandler,
		"exists":            R.existsHandler,
		"type":              R.typeHandler,
		"rename":            R.renameHandler,
		"renamenx":          R.renamenxHandler,
		"move":              R.moveHandler,
		"del":               R.delHandler,
		"randomkey":         R.randomkeyHandler,
		"dbsize":            R.dbsizeHandler,
		"keys":              R.keysHandler,
		"scan":              R.scanHandler,
		"sort":              R.sortHandler,
		"flushdb":           R.flushdbHandler,
		"flushall":          R.flushallHandler,
		"select":            R.selectHandler,
		"swapdb":            R.swapdbHandler,
		"expire":            R.expireHandler,
		"expireat":          R.expireatHandler,
		"ttl":               R.ttlHandler,
		"persist":           R.persistHandler,
		"pexpire":           R.pexpireHandler,
		"pexpireat":         R.pexpireatHandler,
		"pttl":              R.pttlHandler,
		"multi":             R.multiHandler,
		"exec":              R.execHandler,
		"discard":           R.discardHandler,
		"watch":             R.watchHandler,
		"unwatch":           R.unwatchHandler,
		"eval":              R.evalHandler,
		"evalsha":           R.evalshaHandler,
		"script_load":       R.scriptLoadHandler,
		"script_exists":     R.scriptExistsHandler,
		"script_flush":      R.scriptFlushHandler,
		"script_kill":       R.scriptKillHandler,
		"save":              R.saveHandler,
		"bgsave":            R.bgsaveHandler,
		"bgrewriteaof":      R.bgrewriteaofHandler,
		"lastsave":          R.lastsaveHandler,
		"publish":           R.publishHandler,
		"subscribe":         R.subscribeHandler,
		"psubscribe":        R.psubscribeHandler,
		"unsubscribe":       R.unsubscribeHandler,
		"punsubscribe":      R.punsubscribeHandler,
		"pubsub":            R.pubsubHandler,
		"slaveof":           R.slaveofHandler,
		"role":              R.roleHandler,
		"auth":              R.authHandler,
		"quit":              R.quitHandler,
		"info":              R.infoHandler,
		"shutdown":          R.shutdownHandler,
		"time":              R.timeHandler,
		"client_getname":    R.clientGetnameHandler,
		"client_kill":       R.clientKillHandler,
		"client_list":       R.clientListHandler,
		"client_setname":    R.clientSetnameHandler,
		"config_set":        R.configSetHandler,
		"config_get":        R.configGetHandler,
		"config_resetstat":  R.configResetstatHandler,
		"config_rewrite":    R.configRewriteHandler,
		"ping":              R.pingHandler,
		"echo":              R.echoHandler,
		"object":            R.objectHandler,
		"slowlog":           R.slowlogHandler,
		"monitor":           R.monitorHandler,
		"debug_object":      R.debugObjectHandler,
		"debug_segfault":    R.debugSegfaultHandler,
		"migrate":           R.migrateHandler,
		"dump":              R.dumpHandler,
		"restore":           R.restoreHandler,
		"sync":              R.syncHandler,
		"psync":             R.psyncHandler,
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
	r.Clear()
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
