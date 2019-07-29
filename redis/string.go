package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

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

func (r *Redis) setnxHandler(content string) error {
	return nil
}
func (r *Redis) setexHandler(content string) error {
	return nil
}
func (r *Redis) psetexHandler(content string) error {
	return nil
}

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
func (r *Redis) getsetHandler(content string) error {
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
func (r *Redis) appendHandler(content string) error {
	return nil
}
func (r *Redis) setrangeHandler(content string) error {
	return nil
}
func (r *Redis) getrangeHandler(content string) error {
	return nil
}
func (r *Redis) incrHandler(content string) error {
	return nil
}
func (r *Redis) incrbyHandler(content string) error {
	return nil
}
func (r *Redis) incrbyfloatHandler(content string) error {
	return nil
}
func (r *Redis) decrHandler(content string) error {
	return nil
}
func (r *Redis) decrbyHandler(content string) error {
	return nil
}
func (r *Redis) msetHandler(content string) error {
	return nil
}
func (r *Redis) msetnxHandler(content string) error {
	return nil
}
func (r *Redis) mgetHandler(content string) error {
	return nil
}
