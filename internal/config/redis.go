package config

import "github.com/timfame/rusprofile-service/pkg/env"

const (
	redisHostEnv     = "REDIS_HOST"
	redisPortEnv     = "REDIS_PORT"
	redisDbNumEnv    = "REDIS_DB_NUM"
	redisPasswordEnv = "REDIS_PASSWORD"
)

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbNum    int    `json:"db_num"`
	Password string `json:"password"`
}

func (r *Redis) Init() (err error) {
	r.Host, err = env.GetString(redisHostEnv)
	if err != nil {
		return err
	}
	r.Port, err = env.GetString(redisPortEnv)
	if err != nil {
		return err
	}
	dbNum, err := env.GetInt64(redisDbNumEnv)
	if err != nil {
		return err
	}
	r.DbNum = int(dbNum)
	r.Password, _ = env.GetString(redisPasswordEnv)
	return nil
}
