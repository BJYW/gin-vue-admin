package global

{{- if .HasGlobal }}

import "github.com/BJYW/gin-vue-admin/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}