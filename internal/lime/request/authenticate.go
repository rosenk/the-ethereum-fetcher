package request

type Authenticate struct {
	Username string `json:"username" validate:"required,notblank"`
	Password string `json:"password" validate:"required,notblank"`
}
