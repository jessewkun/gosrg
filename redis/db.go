package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) selectHandler(db string) error {
	r.Output = append(r.Output, []string{"SELECT " + db, OUTPUT_COMMAND})
	_, err := redis.String(r.Conn.Do("SELECT", db))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Db, _ = strconv.Atoi(db)
	return nil
}

func (r *Redis) keysHandler(content string) error {
	var err error
	if r.Pattern == "*" {
		r.Output = append(r.Output, []string{"KEYS *", OUTPUT_COMMAND})
		r.Keys, err = redis.Strings(r.Conn.Do("KEYS", "*"))
	} else {
		r.Output = append(r.Output, []string{"KEYS *" + r.Pattern + "*", OUTPUT_COMMAND})
		r.Keys, err = redis.Strings(r.Conn.Do("KEYS", "*"+r.Pattern+"*"))
	}
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	return nil
}

func (r *Redis) delHandler(content string) error {
	r.Output = append(r.Output, []string{"DEL " + r.CurrentKey, OUTPUT_COMMAND})
	res, err := redis.Int64(r.Conn.Do("DEL", r.CurrentKey))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return nil
}

func (r *Redis) typeHandler(key string) error {
	r.Output = append(r.Output, []string{"TYPE " + key, OUTPUT_COMMAND})
	keyType, err := redis.String(r.Conn.Do("TYPE", key))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.CurrentKey = key
	r.CurrentKeyType = keyType
	r.Output = append(r.Output, []string{keyType, OUTPUT_RES})
	r.Info = append(r.Info, []string{"type", keyType})
	return nil
}
