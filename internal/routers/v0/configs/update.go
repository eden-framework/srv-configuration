package configs

import (
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/robotic-framework/srv-configuration/internal/global"
	"github.com/robotic-framework/srv-configuration/internal/modules"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(UpdateConfig{}))
}

// 更新配置
type UpdateConfig struct {
	httpx.MethodPatch
	ConfigurationID uint64                          `name:"configID,string" in:"path" default:""`
	Body            modules.UpdateConfigurationBody `name:"body" in:"body"`
}

func (req UpdateConfig) Path() string {
	return "/:configID"
}

func (req UpdateConfig) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	err = modules.UpdateConfiguration(req.ConfigurationID, req.Body, db)
	if err != nil {
		logrus.Errorf("[UpdateConfiguration] modules.UpdateConfiguration err: %v, request: %+v", err, req)
	}
	return
}
