package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"openscrm/common/log"
)

var Store sessions.Store

func Setup(redisHost, redisPassword, aesKey string) {
	var err error
	Store, err = redis.NewStore(10, "tcp", redisHost, redisPassword, []byte(aesKey))
	if err != nil {
		log.Sugar.Fatalw("setup session failed", "err", err)
		return
	}
}
