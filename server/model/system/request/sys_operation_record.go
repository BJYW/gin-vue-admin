package request

import (
	"github.com/BJYW/gin-vue-admin/server/model/common/request"
	"github.com/BJYW/gin-vue-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
