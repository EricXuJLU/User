package dao

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Option interface {
	Apply(*Dao) error
}

// WithDB .
func WithDB(db *gorm.DB) Option {
	return withDB{db}
}

type withDB struct {
	db *gorm.DB
}

func (w withDB) Apply(d *Dao) error {
	d.DB = w.db
	return nil
}

// WithRedis .
func WithRedis(r *redis.Client) Option {
	return withRedis{r}
}

type withRedis struct {
	r *redis.Client
}

func (w withRedis) Apply(d *Dao) error {
	d.R = w.r
	return nil
}
