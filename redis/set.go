package redis

import (
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) saddHandler(content string) error {
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
	r.Output = append(r.Output, []string{"SADD " + content, OUTPUT_COMMAND})
	res, err := redis.Int64(r.Conn.Do("SADD", args...))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.CurrentKey = key
	r.CurrentKeyType = TYPE_SET
	r.Output = append(r.Output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return nil
}

func (r *Redis) scardHandler(key string) error {
	r.Output = append(r.Output, []string{"SCARD " + r.CurrentKey, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Conn.Do("SCARD", r.CurrentKey))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Info = append(r.Info, []string{"scard", strconv.FormatInt(lenres, 10)})
	return nil
}

func (r *Redis) smembersHandler(key string) error {
	var err error
	r.Output = append(r.Output, []string{"SMEMBERS " + r.CurrentKey, OUTPUT_COMMAND})
	r.Detail, err = redis.Strings(r.Conn.Do("SMEMBERS", r.CurrentKey))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	return nil
}
