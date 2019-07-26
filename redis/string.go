package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) getHandler(content string) error {
	var err error
	r.Output = append(r.Output, []string{"GET " + r.CurrentKey, OUTPUT_COMMAND})
	r.Detail, err = redis.String(r.Conn.Do("GET", r.CurrentKey))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	return nil
}

func (r *Redis) strlenHandler(content string) error {
	r.Output = append(r.Output, []string{"STRLEN " + r.CurrentKey, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Conn.Do("STRLEN", r.CurrentKey))
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
