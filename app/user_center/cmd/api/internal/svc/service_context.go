package svc

import (
	"go_zero_pgsql/app/user_center/cmd/api/internal/config"
	i18n2 "go_zero_pgsql/app/user_center/cmd/api/internal/i18n"
	"go_zero_pgsql/app/user_center/cmd/api/internal/middleware"
	genModel "go_zero_pgsql/app/user_center/model"

	"go_zero_pgsql/common/i18n"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel genModel.UsersModel
	Redis      *redis.Client
	I18n       rest.Middleware
	Trans      *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {
	var postgresql = sqlx.NewSqlConn("pgx", c.Database.PGDataSource)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       0,
	})
	trans := i18n.NewTranslator(c.I18nConf, i18n2.LocaleFS)
	return &ServiceContext{
		Config:     c,
		I18n:       middleware.NewI18nMiddleware().Handle,
		UsersModel: genModel.NewUsersModel(postgresql, c.Cache),
		Redis:      redisClient,
		Trans:      trans,
	}
}
