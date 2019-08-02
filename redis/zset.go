package redis

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) zaddHandler(content string) error {
	key := r.CurrentKey
	tmpArr := strings.Split(content, "\n")
	var args []interface{}
	temp := key
	args = append(args, key)
	for k, v := range tmpArr {
		t := strings.Split(v, SEPARATOR)
		if len(t) != 2 {
			err := errors.New("Line " + strconv.Itoa(k+1) + " include incorrect format data")
			r.Send(RES_OUTPUT_ERROR, err.Error())
			return err
		}
		temp += " " + t[1] + " " + t[0]
		args = append(args, t[1], t[0])
	}
	if err := r.delHandler(""); err != nil {
		return err
	}
	r.Send(RES_OUTPUT_COMMAND, "ZADD "+temp)
	res, err := redis.Int64(r.Conn.Do("ZADD", args...))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.CurrentKey = key
	r.CurrentKeyType = TYPE_ZSET
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) zcardHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "ZCARD "+r.CurrentKey)
	lenres, err := redis.Int64(r.Conn.Do("ZCARD", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_INFO, []string{"zcard", strconv.FormatInt(lenres, 10)})
	return nil
}

func (r *Redis) zrangeHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "ZRANGE "+r.CurrentKey+" 0 -1 WITHSCORES")
	res, err := redis.StringMap(r.Conn.Do("ZRANGE", r.CurrentKey, 0, -1, "WITHSCORES"))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_DETAIL, res)
	return nil
}
