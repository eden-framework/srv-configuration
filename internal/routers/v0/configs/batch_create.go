package configs

import (
	"context"
	"github.com/eden-framework/client/client_srv_id"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-framework/sqlx"
	"github.com/robotic-framework/srv-configuration/internal/global"
	"github.com/robotic-framework/srv-configuration/internal/modules"
	"github.com/sirupsen/logrus"
)

func init() {
	Router.Register(courier.NewRouter(BatchCreateConfigs{}))
}

// 批量创建配置
type BatchCreateConfigs struct {
	httpx.MethodPost

	Body []modules.CreateConfigurationBody `in:"body"`
}

func (req BatchCreateConfigs) Path() string {
	return "/batch"
}

func (req BatchCreateConfigs) Output(ctx context.Context) (result interface{}, err error) {
	db := global.Config.MasterDB.Get()
	tx := sqlx.NewTasks(db)

	for _, param := range req.Body {
		tx = tx.With(GetTransaction(param, global.ClientConfig.ID))
	}

	if err = tx.Do(); err != nil {
		logrus.Errorf("[BatchCreateConfig] transaction modules.CreateConfiguration err: %v, req: %+v", err, req.Body)
	}
	return
}

func GetTransaction(req modules.CreateConfigurationBody, idClient client_srv_id.ClientSrvIDInterface) func(db sqlx.DBExecutor) error {
	return func(db sqlx.DBExecutor) error {
		err := modules.CreateConfiguration(req, db, idClient)
		if err != nil {
			return err
		}
		return nil
	}
}
