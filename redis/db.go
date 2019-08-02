package redis

import (
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func (r *Redis) selectHandler(db string) error {
	r.Send(RES_OUTPUT_COMMAND, "SELECT "+db)
	_, err := redis.String(r.Conn.Do("SELECT", db))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Db, _ = strconv.Atoi(db)
	return nil
}

func (r *Redis) keysHandler(content string) error {
	var err error
	var res []string
	if r.Pattern == "*" {
		r.Send(RES_OUTPUT_COMMAND, "KEYS *")
		res, err = redis.Strings(r.Conn.Do("KEYS", "*"))
	} else {
		r.Send(RES_OUTPUT_COMMAND, "KEYS *"+r.Pattern+"*")
		res, err = redis.Strings(r.Conn.Do("KEYS", "*"+r.Pattern+"*"))
	}
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_KEYS, res)
	return nil
}

func (r *Redis) delHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "DEL "+r.CurrentKey)
	res, err := redis.Int64(r.Conn.Do("DEL", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) typeHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "TYPE "+r.CurrentKey)
	res, err := redis.String(r.Conn.Do("TYPE", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.CurrentKeyType = res
	r.Send(RES_OUTPUT_RES, res)
	r.Send(RES_INFO, []string{"type", res})
	return nil
}

func (r *Redis) dbsizeHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "DBSIZE")
	res, err := redis.Int64(r.Conn.Do("DBSIZE"))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_INFO, []string{"dbsize", strconv.FormatInt(res, 10)})
	return nil
}

func (r *Redis) existsHandler(content string) error {
	if content == "" {
		err := errors.New("ERR wrong number of arguments for 'exists' command")
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.CurrentKey = content
	r.Send(RES_OUTPUT_COMMAND, "EXISTS "+r.CurrentKey)
	res, err := redis.Int64(r.Conn.Do("EXISTS", r.CurrentKey))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) renameHandler(content string) error {
	t := strings.Split(content, " ")
	if len(t) != 2 {
		err := errors.New("ERR wrong number of arguments for 'rename' command")
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.CurrentKey = t[0]
	r.Send(RES_OUTPUT_COMMAND, "RENAME "+r.CurrentKey+" "+t[1])
	res, err := redis.String(r.Conn.Do("RENAME", r.CurrentKey, t[1]))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) renamenxHandler(content string) error {
	t := strings.Split(content, " ")
	if len(t) != 2 {
		err := errors.New("ERR wrong number of arguments for 'renamenx' command")
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.CurrentKey = t[0]
	r.Send(RES_OUTPUT_COMMAND, "RENAMENX "+r.CurrentKey+" "+t[1])
	res, err := redis.Int64(r.Conn.Do("RENAMENX", r.CurrentKey, t[1]))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) moveHandler(content string) error {
	t := strings.Split(content, " ")
	if len(t) != 2 {
		err := errors.New("ERR wrong number of arguments for 'move' command")
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.CurrentKey = t[0]
	r.Send(RES_OUTPUT_COMMAND, "MOVE "+r.CurrentKey+" "+t[1])
	res, err := redis.Int64(r.Conn.Do("MOVE", r.CurrentKey, t[1]))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) randomkeyHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "RANDOMKEY")
	res, err := redis.String(r.Conn.Do("RANDOMKEY"))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) flushdbHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "FLUSHDB")
	res, err := redis.String(r.Conn.Do("FLUSHDB"))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) flushallHandler(content string) error {
	r.Send(RES_OUTPUT_COMMAND, "FLUSHALL")
	res, err := redis.String(r.Conn.Do("FLUSHALL"))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_RES, res)
	return nil
}

func (r *Redis) swapdbHandler(content string) error {
	t := strings.Split(content, " ")
	if len(t) != 2 {
		err := errors.New("ERR wrong number of arguments for 'swapdb' command")
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_COMMAND, "SWAPDB "+t[0]+" "+t[1])
	res, err := redis.String(r.Conn.Do("SWAPDB", t[0], t[1]))
	if err != nil {
		r.Send(RES_OUTPUT_ERROR, err.Error())
		return err
	}
	r.Send(RES_OUTPUT_RES, res)
	return nil
}
