package v0

import (
	"github.com/eden-framework/courier"
	"github.com/robotic-framework/srv-configuration/internal/routers/v0/configs"
)

var Router = courier.NewRouter(V0Router{})

func init() {
	Router.Register(configs.Router)
}

type V0Router struct {
	courier.EmptyOperator
}

func (V0Router) Path() string {
	return "/v0"
}
