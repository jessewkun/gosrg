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
	r.Send(RES_OUTPUT_COMMAND, "SADD "+content)
	res, err := redis.Int64(r.Conn.Do("SADD", args...))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.CurrentKey = key
	r.CurrentKeyType = TYPE_SET
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) scardHandler(key string) error {
	r.Send(RES_OUTPUT_COMMAND, "SCARD "+r.CurrentKey)
	lenres, err := redis.Int64(r.Conn.Do("SCARD", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_INFO, []string{"scard", strconv.FormatInt(lenres, 10)})
	return nil
}

func (r *Redis) smembersHandler(key string) error {
	r.Send(RES_OUTPUT_COMMAND, "SMEMBERS "+r.CurrentKey)
	res, err := redis.Strings(r.Conn.Do("SMEMBERS", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_DETAIL, res)
	return nil
}
