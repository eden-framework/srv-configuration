package modules

import (
	"github.com/eden-framework/client"
	"github.com/eden-framework/client/client_srv_id"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/robotic-framework/srv-configuration/internal/constants/errors"
	"github.com/robotic-framework/srv-configuration/internal/databases"
)

type CreateConfigurationBody struct {
	// StackID
	StackID uint64 `json:"stackID,string"`
	// Key
	Key string `json:"key"`
	// Value
	Value string `json:"value"`
	// Remark
	Remark string `json:"remark" default:""`
}

func CreateConfiguration(req CreateConfigurationBody, db sqlx.DBExecutor, clientID client_srv_id.ClientSrvIDInterface) error {
	id, err := client.GetUniqueID(clientID)
	if err != nil {
		return err
	}
	model := &databases.Configuration{
		ConfigurationID: id,
		StackID:         req.StackID,
		Key:             req.Key,
		Value:           req.Value,
		Remark:          req.Remark,
	}

	return model.Create(db)
}

func FetchConfiguration(stackID uint64, db sqlx.DBExecutor) (result []databases.Configuration, err error) {
	model := &databases.Configuration{}
	table := db.T(model)
	if stackID == 0 {
		err = errors.BadRequest
		return
	}

	condition := builder.And(table.F("StackID").Eq(stackID))
	result, err = model.List(db, condition)
	return
}

type UpdateConfigurationBody struct {
	// StackID
	StackID uint64 `json:"stackID,string" default:""`
	// Key
	Key string `json:"key" default:""`
	// Value
	Value string `json:"value" default:""`
	// Remark
	Remark string `json:"remark" default:""`
}

func (req UpdateConfigurationBody) Validation(configID uint64) error {
	if configID == 0 && req.StackID == 0 && req.Key == "" {
		return errors.BadRequest.StatusError().WithMsg("configurationID stackID key 不能全为空").WithErrTalk()
	}
	if configID == 0 && ((req.StackID != 0 && req.Key == "") || (req.StackID == 0 && req.Key != "")) {
		return errors.BadRequest.StatusError().WithMsg("stackID key 不能单一为空").WithErrTalk()
	}
	return nil
}

func UpdateConfiguration(configID uint64, req UpdateConfigurationBody, db sqlx.DBExecutor) error {
	err := req.Validation(configID)
	if err != nil {
		return err
	}
	model := &databases.Configuration{
		Value:  req.Value,
		Remark: req.Remark,
	}
	if configID != 0 {
		model.ConfigurationID = configID
		err = model.UpdateByConfigurationIDWithStruct(db)
	} else {
		model.StackID = req.StackID
		model.Key = req.Key
		err = model.UpdateByStackIDAndKeyWithStruct(db)
	}
	return err
}
