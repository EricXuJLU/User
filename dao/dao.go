package dao

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
)

type Dao struct {
	DB *sql.DB
	R  *redis.Client
}

func New(opts ...Option) (*Dao, error) {
	d := &Dao{}
	for _, opt := range opts {
		if err := opt.Apply(d); err != nil {
			return nil, err
		}
	}
	return d, nil
}
