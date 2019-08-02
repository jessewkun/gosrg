package redis

import (
	"strings"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) debugObjectHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "DEBUG OBJECT "+r.CurrentKey)
	object, err := redis.String(r.Conn.Do("DEBUG", "OBJECT", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	objectArr := strings.Split(object, " ")
	for _, item := range objectArr {
		tmp := strings.Split(item, ":")
		if len(tmp) == 1 {
			continue // Value
		}
		if tmp[0] == "at" {
			r.Send(RES_INFO, []string{"Value at", tmp[1]})
		} else {
			r.Send(RES_INFO, []string{tmp[0], tmp[1]})
		}
	}
	return nil
}
