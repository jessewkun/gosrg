package redis

import (
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) lpushHandler(content string) error {
	return nil
}
func (r *Redis) lpushxHandler(content string) error {
	return nil
}
func (r *Redis) rpushHandler(content string) error {
	key := r.CurrentKey
	tmpArr := strings.Split(content, "\n")
	content = key + " " + strings.Join(tmpArr, " ")
	var args []interface{}
	args = append(args, key)
	for _, v := range tmpArr {
		args = append(args, v)
	}
	if err := r.delHandler(""); err != nil {
		return err
	}
	r.Output = append(r.Output, []string{"RPUSH " + content, OUTPUT_COMMAND})
	res, err := redis.Int64(r.Conn.Do("RPUSH", args...))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.CurrentKey = key
	r.CurrentKeyType = TYPE_LIST
	r.Output = append(r.Output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return nil
}
func (r *Redis) rpushxHandler(content string) error {
	return nil
}
func (r *Redis) lpopHandler(content string) error {
	return nil
}
func (r *Redis) rpopHandler(content string) error {
	return nil
}
func (r *Redis) rpoplpushHandler(content string) error {
	return nil
}
func (r *Redis) lremHandler(content string) error {
	return nil
}
func (r *Redis) llenHandler(content string) error {
	r.Output = append(r.Output, []string{"LLEN " + r.CurrentKey, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Conn.Do("LLEN", r.CurrentKey))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Info = append(r.Info, []string{"llen", strconv.FormatInt(lenres, 10)})
	return nil
}
func (r *Redis) lindexHandler(content string) error {
	return nil
}
func (r *Redis) linsertHandler(content string) error {
	return nil
}
func (r *Redis) lsetHandler(content string) error {
	return nil
}
func (r *Redis) lrangeHandler(contnt string) error {
	var err error
	r.Output = append(r.Output, []string{"LRANGE " + r.CurrentKey + " 0 -1", OUTPUT_COMMAND})
	r.Detail, err = redis.Strings(r.Conn.Do("LRANGE", r.CurrentKey, 0, -1))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}

	return nil
}
func (r *Redis) ltrimHandler(content string) error {
	return nil
}
func (r *Redis) blpopHandler(content string) error {
	return nil
}
func (r *Redis) brpopHandler(content string) error {
	return nil
}
func (r *Redis) brpoplpushHandler(content string) error {
	return nil
}
