package core

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"github.com/genzai-io/sliced/app/api"
)

type NodeTransport interface {
	Send(command api.Command)
}

type localNodeTransport struct {
}

func (t *localNodeTransport) Send(command api.Command) {
}

type remoteNodeTransport struct {
	pool *redis.Pool
}

func newNodeTransport(target string) *remoteNodeTransport {
	return &remoteNodeTransport{
		pool: &redis.Pool{
			MaxIdle: 5, // figure 5 should suffice most clusters.
			//MaxActive:   25,
			IdleTimeout: time.Minute, //
			Wait:        false,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", target)
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		},
	}
}

func (t *remoteNodeTransport) Get() redis.Conn {
	return t.pool.Get()
}

func (t *remoteNodeTransport) Send(command api.Command) {
	conn := t.pool.Get()
	if conn == nil {
		return
	}

}