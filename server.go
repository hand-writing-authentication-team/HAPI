package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/hand-writing-authentication-team/HAPI/controllers"
	"github.com/hand-writing-authentication-team/HAPI/queue"
)

type HAPIServerConfig struct {
	addr   string
	server *mux.Router
	QC     *queue.Queue
	RQ     *queue.ResultQueue
}

func config() HAPIServerConfig {
	var conf HAPIServerConfig
	port := strings.TrimSpace(os.Getenv("SERVER_PORT"))
	if port == "" {
		port = "9099"
	}
	conf.addr = fmt.Sprintf("0.0.0.0:%s", port)

	// mqHost := strings.TrimSpace(os.Getenv("MQ_HOST"))
	// mqPort := strings.TrimSpace(os.Getenv("MQ_PORT"))
	// mqUsername := strings.TrimSpace(os.Getenv("MQ_USER"))
	// mqPassword := strings.TrimSpace(os.Getenv("MQ_PASSWORD"))

	// redisAddr := strings.TrimSpace(os.Getenv("REDIS_ADDR"))

	// if mqHost == "" || mqPassword == "" || mqPort == "" || mqUsername == "" {
	// 	log.Fatal("one of the mq config env is not set!")
	// 	os.Exit(1)
	// }

	// if redisAddr == "" {
	// 	log.Fatal("one of the redis configuration is not set")
	// 	os.Exit(1)
	// }

	// queueClient, err := queue.NewQueueInstance(mqHost, mqPort, mqUsername, mqPassword)
	// if err != nil {
	// 	os.Exit(1)
	// }
	// conf.QC = queueClient
	// conf.RQ, err = queue.NewRedisClient(redisAddr)
	// if err != nil {
	// 	os.Exit(1)
	// }

	conf.server = controllers.NewServerControllerSet().Server
	return conf
}

func main() {
	serverConf := config()
	log.Info("start to start the server")
	log.Fatal(http.ListenAndServe(serverConf.addr, serverConf.server))
}
