package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hand-writing-authentication-team/HAPI/models"
	"github.com/hand-writing-authentication-team/HAPI/utils"
	log "github.com/sirupsen/logrus"
)

func (c *ControllerConf) CreateAccoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Error("error occured when reading body of the request")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to read request body")
		return
	}
	var accountReq models.HAPIReq

	err = json.Unmarshal(body, &accountReq)
	if err != nil {
		log.WithError(err).Error("error occured when unmarshalling the Creation request")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "malformed json request body")
		return
	}

	err = c.PublishAndListen(accountReq, utils.CreateAction, w)
	if err != nil {
		log.WithError(err).Errorf("error occured when creating account for user %s", accountReq.Username)
		return
	}

	log.Info("successfully created the account!")
	return
}

func (c *ControllerConf) AuthAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Error("error occured when reading body of the request")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to read request body")
		return
	}
	var accountReq models.HAPIReq

	err = json.Unmarshal(body, &accountReq)
	if err != nil {
		log.WithError(err).Error("error occured when unmarshalling the Auth request")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "malformed json request body")
		return
	}

	err = c.PublishAndListen(accountReq, utils.AuthAction, w)
	if err != nil {
		log.WithError(err).Errorf("error occured when creating account for user %s", accountReq.Username)
		return
	}

	log.Info("successfully authenticated the account!")
	return
}
