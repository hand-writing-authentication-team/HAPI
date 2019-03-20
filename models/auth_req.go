package models

type AuthenticationRequest struct {
	JobID     string `json:"jobid"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Race      string `json:"race"`
	Handwring string `json:"handwriting"`
	Action    string `json:"action"`
}
