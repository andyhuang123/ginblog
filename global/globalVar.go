package global

import (
	"gin-blog/news/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Settings config.ServerConfig
)

var (
	Lg *zap.Logger
)

var Trans ut.Translator

var DB *gorm.DB

var Redis *redis.Client
