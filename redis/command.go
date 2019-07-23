package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func (r *Redis) selectHandler(db string) error {
	r.Output = append(r.Output, []string{"SELECT " + db, OUTPUT_COMMAND})
	_, err := redis.String(r.Redis.Do("SELECT", db))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Db, _ = strconv.Atoi(db)
	return nil
}

func (r *Redis) KeysHandler(content string) error {
	// var err error
	// var keys []string
	// if r.Pattern == "*" {
	// 	r.Output = append(r.Output, []string{"KEYS *", OUTPUT_COMMAND})
	// 	keys, err = redis.Strings(r.Redis.Do("KEYS", "*"))
	// } else {
	// 	r.Output = append(r.Output, []string{"KEYS *" + r.Pattern + "*", OUTPUT_COMMAND})
	// 	keys, err = redis.Strings(r.Redis.Do("KEYS", "*"+r.Pattern+"*"))
	// }
	// if err != nil {
	// 	r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
	// 	return err
	// }
	// r.Output = append(r.Output, keys)
	return nil
}
