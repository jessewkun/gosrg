package redis

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) hgetallHandler(key string) error {
	var err error
	r.Output = append(r.Output, []string{"HGETALl " + key, OUTPUT_COMMAND})
	r.Detail, err = redis.StringMap(r.Conn.Do("HGETALL", key))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.hlenHandler(key)
	return nil
}

func (r *Redis) hlenHandler(key string) error {
	r.Output = append(r.Output, []string{"HLEN " + key, OUTPUT_COMMAND})
	lenres, err := redis.Int64(r.Conn.Do("HLEN", key))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Info = append(r.Info, []string{"hlen", strconv.FormatInt(lenres, 10)})
	return nil
}

func (r *Redis) hmsetHandler(content string) error {
	key := r.CurrentKey
	tmpArr := strings.Split(content, "\n")
	var args []interface{}
	temp := key
	args = append(args, key)
	for k, v := range tmpArr {
		t := strings.Split(v, SEPARATOR)
		if len(t) != 2 {
			err := errors.New("Line " + strconv.Itoa(k+1) + " include incorrect format data")
			r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
			return err
		}
		temp += " " + t[0] + " " + t[1]
		args = append(args, t[0], t[1])
	}

	if err := r.delHandler(""); err != nil {
		return err
	}
	r.Output = append(r.Output, []string{"HMSET " + temp, OUTPUT_COMMAND})
	res, err := redis.String(r.Conn.Do("HMSET", args...))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.CurrentKey = key
	r.CurrentKeyType = "hash"
	r.Output = append(r.Output, []string{res, OUTPUT_RES})
	return nil
}
