package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hand-writing-authentication-team/HAPI/models"
	"github.com/hand-writing-authentication-team/HAPI/utils"
	log "github.com/sirupsen/logrus"
)

func (c *ControllerConf) PublishAndListen(req models.HAPIReq, action string, w http.ResponseWriter) error {
	var resp models.HAPIResp
	resp.Username = req.Username
	jobID := utils.JobIDGenerator()
	queueActionReq := models.AuthenticationRequest{
		JobID:     jobID,
		Username:  req.Username,
		Password:  req.Password,
		Handwring: req.Handwriting,
		Action:    action,
	}
	err := c.QC.Publish("", c.QC.QueueName, queueActionReq)
	if err != nil {
		resp.Status = utils.StatusError
		log.WithError(err).Error("error occured when unmarshalling the account request")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return err
	}
	result, err := c.RQ.Listen(jobID)
	if err != nil {
		resp.Status = utils.StatusError
		switch err.Error() {
		case utils.ErrorMsgTimeout:
			w.WriteHeader(http.StatusGatewayTimeout)
			break
		case utils.ErrorMsgInternalServerError:
			w.WriteHeader(http.StatusInternalServerError)
			break
		default:
			w.WriteHeader(http.StatusInternalServerError)
			break
		}
		json.NewEncoder(w).Encode(resp)
		return err
	}

	status := result.Status
	resp.Status = status
	log.Infof("job %s status %s", jobID, status)
	switch status {
	case utils.StatusError:
		resp.ErrorMsg = result.ErrorMsg
		log.Errorf("account creation failed as backend failed for job %s", jobID)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return errors.New(status)
	case utils.StatusConflict:
		resp.ErrorMsg = result.ErrorMsg
		log.Errorf("account creation failed as account already exists for job %s", jobID)
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(resp)
		return errors.New(status)
	case utils.StatusCreated:
		break
	case utils.StatusAuthenticated:
		break
	case utils.StatusSuccess:
		break
	default:
		resp.Status = utils.StatusError
		log.Errorf("met unrecognizable status %s for jobid %s", status, jobID)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return errors.New(status)
	}
	json.NewEncoder(w).Encode(resp)
	return nil
}
