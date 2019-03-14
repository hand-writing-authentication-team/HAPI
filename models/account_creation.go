package models

type AccountCreationReq struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Handwriting string `json:"handwriting"`
}

type AccountCreationResp struct {
	Username string `json:"username"`
	ErrorMsg string `json:"error,omitempty"`
	Status   string `json:"status"`
}
