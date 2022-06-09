package response

import "github.com/BJYW/gin-vue-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
