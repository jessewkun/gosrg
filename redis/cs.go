package redis

import (
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
	if err := r.dbsizeHandler(""); err != nil {
		return err
	}
	return nil
}
