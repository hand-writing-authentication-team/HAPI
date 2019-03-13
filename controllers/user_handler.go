package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hand-writing-authentication-team/HAPI/models"
	"github.com/hand-writing-authentication-team/HAPI/utils"
	log "github.com/sirupsen/logrus"
)

func (c *ControllerConf) CreateAccoundHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Error("error occured when reading body of the request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var accountReq models.AccountCreationReq
	err = json.Unmarshal(body, &accountReq)
	if err != nil {
		log.WithError(err).Error("error occured when unmarshalling the account request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var accountResp models.AccountCreationResp
	accountResp.Username = accountReq.Username
	jobID := utils.JobIDGenerator()
	queueActionReq := models.AuthenticationRequest{
		JobID:     jobID,
		Username:  accountReq.Username,
		Password:  accountReq.Password,
		Handwring: accountReq.Handwriting,
		Action:    utils.CreateAction,
	}
	err = c.QC.Publish("", c.QC.QueueName, queueActionReq)
	if err != nil {
		accountResp.Status = utils.StatusError
		log.WithError(err).Error("error occured when unmarshalling the account request")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(accountResp)
		return
	}
	log.Info("successfully created the account!")
	w.WriteHeader(http.StatusOK)
	return
}
