package svc

import (
	"automatix/app/gate/cmd/api/internal/config"

	"github.com/bwmarrin/snowflake"
)

type ServiceContext struct {
	Config    config.Config
	Snowflake *snowflake.Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflake, _ := snowflake.NewNode(c.Id)
	return &ServiceContext{
		Config:    c,
		Snowflake: snowflake,
	}
}
