package redis

import (
	"strings"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) objectHandler(key string) error {
	r.Output = append(r.Output, []string{"DEBUG OBJECT " + key, OUTPUT_COMMAND})
	object, err := redis.String(r.Conn.Do("DEBUG", "OBJECT", key))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	objectArr := strings.Split(object, " ")
	for _, item := range objectArr {
		tmp := strings.Split(item, ":")
		if len(tmp) == 1 {
			continue // Value
		}
		if tmp[0] == "at" {
			r.Info = append(r.Info, []string{"Value at", tmp[1]})
		} else {
			r.Info = append(r.Info, []string{tmp[0], tmp[1]})
		}
	}
	return nil
}
