package response

type Authenticate struct {
	Token string `json:"token" validate:"required,notblank"`
}
