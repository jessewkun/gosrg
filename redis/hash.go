package redis

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) hlenHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "HELN "+r.CurrentKey)
	lenres, err := redis.Int64(r.Conn.Do("HLEN", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_INFO, []string{"hlen", strconv.FormatInt(lenres, 10)})
	return nil
}

func (r *Redis) hmsetHandler(content string) error {
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
		temp += " " + t[0] + " " + t[1]
		args = append(args, t[0], t[1])
	}

	if err := r.delHandler(""); err != nil {
		return err
	}
	r.Send(RES_OUTPUT_COMMAND, "HMSET "+temp)
	res, err := redis.String(r.Conn.Do("HMSET", args...))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.CurrentKey = key
	r.CurrentKeyType = TYPE_HASH
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) hgetallHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "HGETALl "+r.CurrentKey)
	res, err := redis.StringMap(r.Conn.Do("HGETALL", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_DETAIL, res)
	return nil
}
