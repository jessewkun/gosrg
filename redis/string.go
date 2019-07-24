package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) getHandler(key string) error {
	var err error
	r.Output = append(r.Output, []string{"GET " + key, OUTPUT_COMMAND})
	r.Detail, err = redis.String(r.Conn.Do("GET", key))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.strlenHandler(key)
	return nil
}

func (r *Redis) strlenHandler(key string) error {
	r.Output = append(r.Output, []string{"STRLEN " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Conn.Do("STRLEN", key))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Info = append(r.Info, []string{"len", strconv.FormatInt(lenres, 10)})
	return nil
}

func (r *Redis) setHandler(content string) error {
	r.Output = append(r.Output, []string{"SET " + r.CurrentKey + " " + content, OUTPUT_COMMAND})
	res, err := redis.String(r.Conn.Do("SET", r.CurrentKey, content))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{res, OUTPUT_RES})
	return nil
}
