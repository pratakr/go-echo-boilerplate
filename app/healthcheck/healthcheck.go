package healthcheck

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var ctx = context.Background()

type DbResource struct{
	*gorm.DB
}

func (d DbResource) HealthCheck() error{
	sql,err := d.DB.DB()
	if err!=nil{
		return err
	}
	return sql.Ping()
}

type RedisResource struct {
	*redis.Client
}

func (r RedisResource) HealthCheck() error {
	return r.Client.Ping(ctx).Err()
}
