package request

import (
	"github.com/BJYW/gin-vue-admin/server/model/common/request"
	"github.com/BJYW/gin-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
