package service

import (
	"github.com/BJYW/gin-vue-admin/server/service/example"
	"github.com/BJYW/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
