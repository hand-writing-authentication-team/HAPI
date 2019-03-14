package queue

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/hand-writing-authentication-team/HAPI/models"
	"github.com/hand-writing-authentication-team/HAPI/utils"

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
	timeout := time.Second * 5
	t1 := time.Now()
	for true {
		str := rq.redisDB.Get(jobID)
		if strings.TrimSpace(str.Val()) == "" {
			log.Debug("retrieved nothing")
		} else {
			var resultResp models.ResultResp
			err := json.Unmarshal(([]byte)(str.Val()), &resultResp)
			if err != nil {
				log.WithError(err).Errorf("failed to unmarshal result for job %s", jobID)
				return nil, errors.New(utils.ErrorMsgInternalServerError)
			}
			return &resultResp, nil
		}
		if time.Since(t1) > timeout {
			break
		}
	}
	log.Errorf("response timeout for jobID %s", jobID)
	return nil, errors.New(utils.ErrorMsgTimeout)
}
