package custom

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	EchoRequest interface {
		Bind(i interface{}) error
	}

	customRequest struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

var (
	once          sync.Once
	validatorInst *validator.Validate
)

func NewCustomRequest(ctx echo.Context) EchoRequest {
	once.Do(func() {
		validatorInst = validator.New()
	})

	return &customRequest{
		ctx:       ctx,
		validator: validatorInst,
	}
}

func (r *customRequest) Bind(i interface{}) error {
	if err := r.ctx.Bind(i); err != nil {
		return err
	}

	if err := r.validator.Struct(i); err != nil {
		return err
	}

	return nil
}
