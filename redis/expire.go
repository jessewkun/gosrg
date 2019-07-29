package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) ttlHandler(content string) error {
	r.Output = append(r.Output, []string{"TTL " + r.CurrentKey, OUTPUT_COMMAND})
	ttlres, err := redis.Int64(r.Conn.Do("TTL", r.CurrentKey))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Info = append(r.Info, []string{"ttl", strconv.FormatInt(ttlres, 10)})
	return nil
}
