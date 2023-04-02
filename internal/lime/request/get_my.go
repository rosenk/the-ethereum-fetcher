package request

type GetMy struct {
	AuthToken string `json:"-" validate:"required,notblank"`
	UserID    *int64 `json:"-" validate:""`
}
