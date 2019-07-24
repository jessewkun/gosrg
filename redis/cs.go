package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) infoHandler(content string) error {
	var err error
	r.Output = append(r.Output, []string{"INFO", OUTPUT_COMMAND})
	r.Detail, err = redis.String(r.Conn.Do("INFO"))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{"DBSIZE", OUTPUT_COMMAND})
	dbsize, err := redis.Int64(r.Conn.Do("DBSIZE"))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Info = append(r.Info, []string{"dbsize", strconv.FormatInt(dbsize, 10)})
	return nil
}
