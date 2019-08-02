package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) ttlHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "TTL "+r.CurrentKey)
	ttlres, err := redis.Int64(r.Conn.Do("TTL", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_INFO, []string{"ttl", strconv.FormatInt(ttlres, 10)})
	return nil
}
