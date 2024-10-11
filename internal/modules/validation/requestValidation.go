package validation

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	EchoRequest interface {
		Bind(obj any) error
	}

	customEchoRequest struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

var (
	once          sync.Once
	validatorInst *validator.Validate
)

func NewCustomEchoRequest(echoRequest echo.Context) EchoRequest {
	once.Do(func() {
		validatorInst = validator.New()
	})

	return &customEchoRequest{
		ctx:       echoRequest,
		validator: validatorInst,
	}
}

func (r *customEchoRequest) Bind(obj any) error {
	if err := r.ctx.Bind(obj); err != nil {
		return err
	}

	if err := r.validator.Struct(obj); err != nil {
		return err
	}

	return nil
}
