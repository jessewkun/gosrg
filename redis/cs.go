package redis

import (
	"github.com/gomodule/redigo/redis"
)

func (r *Redis) infoHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "INFO")
	res, err := redis.String(r.Conn.Do("INFO"))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_DETAIL, res)
	return nil
}
