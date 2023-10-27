package requests

type UserRegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginRequest struct {
	Number   int64  `jsonL:"number"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type TopupReq struct {
	Amount int64 `json:"amount"`
}
