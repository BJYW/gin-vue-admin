package request

import (
	"github.com/BJYW/gin-vue-admin/server/model/{{.Package}}"
	"github.com/BJYW/gin-vue-admin/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
