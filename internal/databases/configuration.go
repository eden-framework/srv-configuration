package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model Configuration --database Config.DB --with-comments
//go:generate eden generate tag Configuration --defaults=true
// @def primary ID
// @def unique_index U_configuration_id ConfigurationID
// @def unique_index U_stack_key StackID Key
// @def index I_stack StackID
type Configuration struct {
	datatypes.PrimaryID
	// 业务ID
	ConfigurationID uint64 `db:"f_configuration_id" json:"configurationID,string"`
	// StackID
	StackID uint64 `db:"f_stack_id" json:"stackID,string"`
	// Key
	Key string `db:"f_key" json:"key"`
	// Value
	Value string `db:"f_value" json:"value"`
	// Remark
	Remark string `db:"f_remark,default=''" json:"remark"`
	datatypes.OperateTime
}
