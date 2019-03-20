package models

type HAPIReq struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Handwriting string `json:"handwriting"`
	Race        string `json:"race,omitempty"`
}

type HAPIResp struct {
	Username string `json:"username"`
	ErrorMsg string `json:"error,omitempty"`
	Status   string `json:"status"`
}
