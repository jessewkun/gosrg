package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) ttlHandler(key string) error {
	r.Output = append(r.Output, []string{"TTL " + key, OUTPUT_COMMAND})
	ttlres, err := redis.Int64(r.Conn.Do("TTL", key))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Info = append(r.Info, []string{"ttl", strconv.FormatInt(ttlres, 10)})
	return nil
}
