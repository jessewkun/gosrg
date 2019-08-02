package redis

import (
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

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
	r.Send(RES_OUTPUT_COMMAND, "RPUSH "+content)
	res, err := redis.Int64(r.Conn.Do("RPUSH", args...))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.CurrentKey = key
	r.CurrentKeyType = TYPE_LIST
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) llenHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "LLEN "+r.CurrentKey)
	lenres, err := redis.Int64(r.Conn.Do("LLEN", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_INFO, []string{"llen", strconv.FormatInt(lenres, 10)})
	return nil
}

func (r *Redis) lrangeHandler(contnt string) error {
	r.Send(RES_OUTPUT_COMMAND, "LRANGE "+r.CurrentKey+" 0 -1")
	res, err := redis.Strings(r.Conn.Do("LRANGE", r.CurrentKey, 0, -1))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_DETAIL, res)
	return nil
}
