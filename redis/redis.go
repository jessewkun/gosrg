package redis

import (
	"gosrg/utils"
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
	SEPARATOR      = " ==> "
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

func InitRedis(host string, port string, pwd string, pattern string) {
	R = &Redis{
		Host:    host,
		Port:    port,
		Pwd:     pwd,
		Pattern: pattern,
	}

	conn, err := redis.Dial(REDIS_NETWORK, R.Host+":"+R.Port)
	if err != nil {
		utils.Exit(err)
	}

	if R.Pwd != "" {
		if _, err := conn.Do("AUTH", R.Pwd); err != nil {
			conn.Close()
			utils.Exit(err)
		}
	}
	utils.Info.Println("Redis conn ok")

	R.Redis = conn
}

func (r *Redis) SelectDb(db int) (output [][]string) {
	output = append(output, []string{"SELECT " + strconv.Itoa(db), OUTPUT_COMMAND})
	_, err := redis.String(r.Redis.Do("SELECT", db))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	r.CurrentKey = ""
	r.CurrentKeyType = ""
	return
}

func (r *Redis) Keys() (output [][]string, keys []string) {
	var err error
	if r.Pattern == "*" {
		output = append(output, []string{"KEYS *", OUTPUT_COMMAND})
		keys, err = redis.Strings(r.Redis.Do("KEYS", "*"))
	} else {
		output = append(output, []string{"KEYS *" + r.Pattern + "*", OUTPUT_COMMAND})
		keys, err = redis.Strings(r.Redis.Do("KEYS", "*"+r.Pattern+"*"))
	}
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	r.CurrentKey = ""
	r.CurrentKeyType = ""
	return
}

func (r *Redis) Del() (output [][]string, err error) {
	output = append(output, []string{"DEL " + r.CurrentKey, OUTPUT_COMMAND})
	res, err := redis.Int64(r.Redis.Do("DEL", r.CurrentKey))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	r.CurrentKey = ""
	r.CurrentKeyType = ""
	output = append(output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return
}

func (r *Redis) Info() (output [][]string, res string, info [][]string) {
	output = append(output, []string{"INFO", OUTPUT_COMMAND})
	res, err := redis.String(r.Redis.Do("INFO"))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	output = append(output, []string{"DBSIZE", OUTPUT_COMMAND})
	dbsize, err := redis.Int64(r.Redis.Do("DBSIZE"))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	info = append(info, []string{"dbsize", strconv.FormatInt(dbsize, 10)})
	r.CurrentKey = ""
	r.CurrentKeyType = ""
	return
}

func (r *Redis) Send(ommandName string, args ...interface{}) {
	// keyType, err := redis.String(r.Redis.Do(ommandName, args...))
}

func (r *Redis) KeyDetail(key string) (output [][]string, res interface{}, info [][]string) {
	output = append(output, []string{"TYPE " + key, OUTPUT_COMMAND})

	keyType, err := redis.String(r.Redis.Do("TYPE", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	r.CurrentKey = key
	r.CurrentKeyType = keyType
	info = append(info, []string{"type", keyType})

	output = append(output, []string{"DEBUG OBJECT " + key, OUTPUT_COMMAND})
	object, err := redis.String(r.Redis.Do("DEBUG", "OBJECT", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
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
	ttlres, err := redis.Int64(r.Redis.Do("TTL", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	info = append(info, []string{"ttl", strconv.FormatInt(ttlres, 10)})

	switch keyType {
	case "string":
		o, detail, stringinfo := r.getString(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	case "hash":
		o, detail, stringinfo := r.getHash(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	case "set":
		o, detail, stringinfo := r.getSet(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	case "zset":
		o, detail, stringinfo := r.getZset(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	case "list":
		o, detail, stringinfo := r.getList(key)
		output = append(output, o...)
		res = detail
		info = append(info, stringinfo...)
	}

	return
}

func (r *Redis) getString(key string) (output [][]string, res string, info [][]string) {
	output = append(output, []string{"GET " + key, OUTPUT_COMMAND})
	res, err := redis.String(r.Redis.Do("GET", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}

	output = append(output, []string{"STRLEN " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Redis.Do("STRLEN", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	info = append(info, []string{"len", strconv.FormatInt(lenres, 10)})
	return
}

func (r *Redis) getHash(key string) (output [][]string, res map[string]string, info [][]string) {
	output = append(output, []string{"HGETALl " + key, OUTPUT_COMMAND})
	res, err := redis.StringMap(r.Redis.Do("HGETALL", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}

	output = append(output, []string{"HLEN " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Redis.Do("HLEN", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	info = append(info, []string{"hlen", strconv.FormatInt(lenres, 10)})
	return
}

func (r *Redis) getSet(key string) (output [][]string, res []string, info [][]string) {
	output = append(output, []string{"SMEMBERS " + key, OUTPUT_COMMAND})
	res, err := redis.Strings(r.Redis.Do("SMEMBERS", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}

	output = append(output, []string{"SCARD " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Redis.Do("SCARD", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	info = append(info, []string{"scard", strconv.FormatInt(lenres, 10)})
	return
}

func (r *Redis) getZset(key string) (output [][]string, res map[string]string, info [][]string) {
	output = append(output, []string{"ZRANGE " + key + " 0 -1 WITHSCORES", OUTPUT_COMMAND})
	res, err := redis.StringMap(r.Redis.Do("ZRANGE", key, 0, -1, "WITHSCORES"))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}

	output = append(output, []string{"ZCARD " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Redis.Do("ZCARD", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	info = append(info, []string{"zcard", strconv.FormatInt(lenres, 10)})
	return
}

func (r *Redis) getList(key string) (output [][]string, res []string, info [][]string) {
	output = append(output, []string{"LRANGE " + key + " 0 -1", OUTPUT_COMMAND})
	res, err := redis.Strings(r.Redis.Do("LRANGE", key, 0, -1))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}

	output = append(output, []string{"LLEN " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Redis.Do("LLEN", key))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	info = append(info, []string{"llen", strconv.FormatInt(lenres, 10)})
	return
}

func (r *Redis) SetKeyDetail(content string) (output [][]string) {
	if r.CurrentKey == "" || r.CurrentKeyType == "" {
		return
	}
	switch r.CurrentKeyType {
	case "string":
		output, _ = r.setString(content)
	case "hash":
		output, _ = r.setHash(content)
	case "set":
		output, _ = r.setSet(content)
	case "zset":
		output, _ = r.setZset(content)
	case "list":
		output, _ = r.setList(content)
	}
	return
}

func (r *Redis) setString(content string) (output [][]string, err error) {
	output = append(output, []string{"SET " + r.CurrentKey + " " + content, OUTPUT_COMMAND})
	res, err := redis.String(r.Redis.Do("SET", r.CurrentKey, content))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	output = append(output, []string{res, OUTPUT_RES})
	return
}

func (r *Redis) setHash(content string) (output [][]string, err error) {
	key := r.CurrentKey
	output, err = r.Del()
	if err != nil {
		return output, err
	}
	tmpArr := strings.Split(content, "\n")
	var args []interface{}
	temp := key
	args = append(args, key)
	for _, v := range tmpArr {
		t := strings.Split(v, SEPARATOR)
		temp += " " + t[0] + " " + t[1]
		args = append(args, t[0], t[1])
	}
	output = append(output, []string{"HMSET " + temp, OUTPUT_COMMAND})
	res, err := redis.String(r.Redis.Do("HMSET", args...))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	r.CurrentKey = key
	r.CurrentKeyType = "hash"
	output = append(output, []string{res, OUTPUT_RES})
	return
}

func (r *Redis) setSet(content string) (output [][]string, err error) {
	key := r.CurrentKey
	output, err = r.Del()
	if err != nil {
		return output, err
	}
	tmpArr := strings.Split(content, "\n")
	content = key + " " + strings.Join(tmpArr, " ")
	var args []interface{}
	args = append(args, key)
	for _, v := range tmpArr {
		args = append(args, v)
	}
	output = append(output, []string{"SADD " + content, OUTPUT_COMMAND})
	res, err := redis.Int64(r.Redis.Do("SADD", args...))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	r.CurrentKey = key
	r.CurrentKeyType = "set"
	output = append(output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return
}

func (r *Redis) setZset(content string) (output [][]string, err error) {
	key := r.CurrentKey
	output, err = r.Del()
	if err != nil {
		return output, err
	}
	tmpArr := strings.Split(content, "\n")
	var args []interface{}
	temp := key
	args = append(args, key)
	for _, v := range tmpArr {
		t := strings.Split(v, SEPARATOR)
		temp += " " + t[1] + " " + t[0]
		args = append(args, t[1], t[0])
	}
	output = append(output, []string{"ZADD " + temp, OUTPUT_COMMAND})
	res, err := redis.Int64(r.Redis.Do("ZADD", args...))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	r.CurrentKey = key
	r.CurrentKeyType = "zset"
	output = append(output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return
}

func (r *Redis) setList(content string) (output [][]string, err error) {
	key := r.CurrentKey
	output, err = r.Del()
	if err != nil {
		return output, err
	}
	tmpArr := strings.Split(content, "\n")
	content = key + " " + strings.Join(tmpArr, " ")
	var args []interface{}
	args = append(args, key)
	for _, v := range tmpArr {
		args = append(args, v)
	}
	output = append(output, []string{"RPUSH " + content, OUTPUT_COMMAND})
	res, err := redis.Int64(r.Redis.Do("RPUSH", args...))
	if err != nil {
		output = append(output, []string{err.Error(), OUTPUT_ERROR})
		return
	}
	r.CurrentKey = key
	r.CurrentKeyType = "list"
	output = append(output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return
}
