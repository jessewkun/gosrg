package redis

import (
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) lrangeHandler(key string) error {
	var err error
	r.Output = append(r.Output, []string{"LRANGE " + key + " 0 -1", OUTPUT_COMMAND})
	r.Detail, err = redis.Strings(r.Conn.Do("LRANGE", key, 0, -1))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.llenHandler(key)
	return nil
}

func (r *Redis) llenHandler(key string) error {
	r.Output = append(r.Output, []string{"LLEN " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Conn.Do("LLEN", key))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Info = append(r.Info, []string{"llen", strconv.FormatInt(lenres, 10)})
	return nil
}

func (r *Redis) rpushHandler(content string) error {
	key := r.CurrentKey
	tmpArr := strings.Split(content, "\n")
	content = key + " " + strings.Join(tmpArr, " ")
	var args []interface{}
	args = append(args, key)
	for _, v := range tmpArr {
		args = append(args, v)
	}
	if err := r.delHandler(""); err != nil {
		return err
	}
	r.Output = append(r.Output, []string{"RPUSH " + content, OUTPUT_COMMAND})
	res, err := redis.Int64(r.Conn.Do("RPUSH", args...))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.CurrentKey = key
	r.CurrentKeyType = TYPE_LIST
	r.Output = append(r.Output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return nil
}
