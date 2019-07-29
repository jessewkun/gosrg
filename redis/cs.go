package redis

import (
	"github.com/gomodule/redigo/redis"
)

func (r *Redis) authHandler(content string) error {
	return nil
}
func (r *Redis) quitHandler(content string) error {
	return nil
}
func (r *Redis) infoHandler(content string) error {
	var err error
	r.Output = append(r.Output, []string{"INFO", OUTPUT_COMMAND})
	r.Detail, err = redis.String(r.Conn.Do("INFO"))
	if err != nil {
		r.Output = append(r.Output, []string{err.Error(), OUTPUT_ERROR})
		return err
	}
	return nil
}
func (r *Redis) shutdownHandler(content string) error {
	return nil
}
func (r *Redis) timeHandler(content string) error {
	return nil
}
func (r *Redis) clientGetnameHandler(content string) error {
	return nil
}
func (r *Redis) clientKillHandler(content string) error {
	return nil
}
func (r *Redis) clientListHandler(content string) error {
	return nil
}
func (r *Redis) clientSetnameHandler(content string) error {
	return nil
}
