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
	Router.Register(courier.NewRouter(GetConfigurations{}))
}

// 批量获取配置
type GetConfigurations struct {
	httpx.MethodGet
	// StackID
	StackID uint64 `name:"stackID,string" in:"query"`
}

func (req GetConfigurations) Path() string {
	return ""
}

func (req GetConfigurations) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.SlaveDB.Get()
	resp, err := modules.FetchConfiguration(req.StackID, db)
	if err != nil {
		logrus.Errorf("[GetConfigurations] modules.FetchConfiguration err: %v, req: %+v", err, req)
	}
	return resp, err
}
