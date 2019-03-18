package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hand-writing-authentication-team/HAPI/queue"
)

type ControllerConf struct {
	Server *mux.Router
	RQ     *queue.ResultQueue
	QC     *queue.Queue
}

func NewServerControllerSet() *ControllerConf {
	// defined the controller mapping here
	server := mux.NewRouter()
	controllerConf := &ControllerConf{
		Server: server,
	}
	controllerConf.Server.HandleFunc("/version", controllerConf.VersionGet).Methods("GET")
	controllerConf.Server.HandleFunc("/create_account", controllerConf.CreateAccoundHandler).Methods("POST")
	controllerConf.Server.HandleFunc("/login", controllerConf.AuthAccountHandler).Methods("POST")
	controllerConf.Server.HandleFunc("/collect_handwriting", controllerConf.CreateAccoundHandler).Methods("POST")

	return controllerConf
}

func (c *ControllerConf) VersionGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "version: 0.0.1")
}
