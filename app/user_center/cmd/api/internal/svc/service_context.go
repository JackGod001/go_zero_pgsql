package svc

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_zero_pgsql/app/user_center/cmd/api/internal/config"
	genModel "go_zero_pgsql/app/user_center/model"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel genModel.UsersModel
	Redis      *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	var postgresql = sqlx.NewSqlConn("pgx", c.Database.PGDataSource)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       0,
	})
	return &ServiceContext{
		Config:     c,
		UsersModel: genModel.NewUsersModel(postgresql, c.Cache),
		Redis:      redisClient,
	}
}
