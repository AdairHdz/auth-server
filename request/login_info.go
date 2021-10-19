package request

type LoginInfo struct {
	EmailAddress string `json:"emailAddress"`
	Password string `json:"password"`
}