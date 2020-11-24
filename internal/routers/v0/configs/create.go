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
	Router.Register(courier.NewRouter(CreateConfig{}))
}

// 创建配置
type CreateConfig struct {
	httpx.MethodPost

	Body modules.CreateConfigurationBody `in:"body"`
}

func (req CreateConfig) Path() string {
	return ""
}

func (req CreateConfig) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	err = modules.CreateConfiguration(req.Body, db, global.ClientConfig.ID)
	if err != nil {
		logrus.Errorf("[CreateConfig] modules.CreateConfiguration err: %v, req: %+v", err, req)
	}
	return
}
