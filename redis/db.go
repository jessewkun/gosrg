package redis

import (
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
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

func (r *Redis) typeHandler(content string) error {
	r.Output = append(r.Output, []string{"TYPE " + r.CurrentKey, OUTPUT_COMMAND})
	res, err := redis.String(r.Conn.Do("TYPE", r.CurrentKey))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.CurrentKeyType = res
	r.Output = append(r.Output, []string{res, OUTPUT_RES})
	r.Info = append(r.Info, []string{"type", res})
	return nil
}

func (r *Redis) dbsizeHandler(content string) error {
	r.Output = append(r.Output, []string{"DBSIZE", OUTPUT_COMMAND})
	res, err := redis.Int64(r.Conn.Do("DBSIZE"))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Info = append(r.Info, []string{"dbsize", strconv.FormatInt(res, 10)})
	return nil
}

func (r *Redis) existsHandler(content string) error {
	if content == "" {
		err := errors.New("ERR wrong number of arguments for 'exists' command")
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.CurrentKey = content
	r.Output = append(r.Output, []string{"EXISTS " + r.CurrentKey, OUTPUT_COMMAND})
	res, err := redis.Int64(r.Conn.Do("EXISTS", r.CurrentKey))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return nil
}

func (r *Redis) renameHandler(content string) error {
	t := strings.Split(content, " ")
	if len(t) != 2 {
		err := errors.New("ERR wrong number of arguments for 'rename' command")
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.CurrentKey = t[0]
	r.Output = append(r.Output, []string{"RENAME " + r.CurrentKey + " " + t[1], OUTPUT_COMMAND})
	res, err := redis.String(r.Conn.Do("RENAME", r.CurrentKey, t[1]))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{res, OUTPUT_RES})
	return nil
}

func (r *Redis) renamenxHandler(content string) error {
	t := strings.Split(content, " ")
	if len(t) != 2 {
		err := errors.New("ERR wrong number of arguments for 'renamenx' command")
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.CurrentKey = t[0]
	r.Output = append(r.Output, []string{"RENAMENX " + r.CurrentKey + " " + t[1], OUTPUT_COMMAND})
	res, err := redis.Int64(r.Conn.Do("RENAMENX", r.CurrentKey, t[1]))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return nil
}

func (r *Redis) moveHandler(content string) error {
	t := strings.Split(content, " ")
	if len(t) != 2 {
		err := errors.New("ERR wrong number of arguments for 'move' command")
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.CurrentKey = t[0]
	r.Output = append(r.Output, []string{"MOVE " + r.CurrentKey + " " + t[1], OUTPUT_COMMAND})
	res, err := redis.Int64(r.Conn.Do("MOVE", r.CurrentKey, t[1]))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{strconv.FormatInt(res, 10), OUTPUT_RES})
	return nil
}

func (r *Redis) randomkeyHandler(content string) error {
	r.Output = append(r.Output, []string{"RANDOMKEY", OUTPUT_COMMAND})
	res, err := redis.String(r.Conn.Do("RANDOMKEY"))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{res, OUTPUT_RES})
	return nil
}

func (r *Redis) flushdbHandler(content string) error {
	r.Output = append(r.Output, []string{"FLUSHDB", OUTPUT_COMMAND})
	res, err := redis.String(r.Conn.Do("FLUSHDB"))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{res, OUTPUT_RES})
	return nil
}

func (r *Redis) flushallHandler(content string) error {
	r.Output = append(r.Output, []string{"FLUSHALL", OUTPUT_COMMAND})
	res, err := redis.String(r.Conn.Do("FLUSHALL"))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{res, OUTPUT_RES})
	return nil
}

func (r *Redis) swapdbHandler(content string) error {
	t := strings.Split(content, " ")
	if len(t) != 2 {
		err := errors.New("ERR wrong number of arguments for 'swapdb' command")
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{"SWAPDB " + t[0] + " " + t[1], OUTPUT_COMMAND})
	res, err := redis.String(r.Conn.Do("SWAPDB", t[0], t[1]))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	r.Output = append(r.Output, []string{res, OUTPUT_RES})
	return nil
}
