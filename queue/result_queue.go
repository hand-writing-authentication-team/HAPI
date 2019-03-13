package queue

import (
	"time"

	"github.com/hand-writing-authentication-team/HAPI/models"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type ResultQueue struct {
	redisDB *redis.Client
}

func NewRedisClient(addr string) (*ResultQueue, error) {
	rq := &ResultQueue{}
	rq.redisDB = redis.NewClient(&redis.Options{
		Addr:         addr,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	_, err := rq.redisDB.Ping().Result()
	if err != nil {
		log.WithError(err).Error("error when pinging redis, will retry in 5")
		var counter int
		for err != nil {
			counter++
			time.Sleep(5 * time.Second)
			log.Infof("retrying for the %s th time", counter)
			_, err = rq.redisDB.Ping().Result()
		}
	}
	log.Info("successfully connected to redis!")
	return rq, nil
}

func (rq *ResultQueue) Listen(jobID string) (*models.ResultResp, error) {
	str := rq.redisDB.Get(jobID)
	str.String()
	// need to do a poc on this
	return nil, nil
}
