package responses

type LoginResponse struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}
