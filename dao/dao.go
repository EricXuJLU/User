package dao

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Dao struct {
	DB *gorm.DB
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
