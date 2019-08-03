package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) setHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "SET "+r.CurrentKey+" "+content)
	res, err := redis.String(r.Conn.Do("SET", r.CurrentKey, content))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) getHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "GET "+r.CurrentKey)
	res, err := redis.String(r.Conn.Do("GET", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_DETAIL, res)
	return nil
}

func (r *Redis) strlenHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "STRLEN "+r.CurrentKey)
	lenres, err := redis.Int64(r.Conn.Do("STRLEN", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_INFO, []string{"strlen", strconv.FormatInt(lenres, 10)})
	return nil
}
