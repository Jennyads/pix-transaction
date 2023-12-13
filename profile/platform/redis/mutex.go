package redis

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"profile/internal/cfg"
)

type Mutex interface {
	Lock()
	Unlock()
}

type mutex struct {
	mutex *redsync.Mutex
}

func (m *mutex) Lock() {
	m.mutex.Lock()
}

func (m *mutex) Unlock() {
	m.mutex.Unlock()
}

func NewMutex(key string, config *cfg.Config) Mutex {
	client := StartRedis(config)
	return &mutex{
		mutex: redsync.New(goredis.NewPool(client)).NewMutex(key),
	}
}
