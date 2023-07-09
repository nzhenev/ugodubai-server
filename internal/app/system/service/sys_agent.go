package service

import (
	"context"

	"ugodubai-server/api/v1/system"
)

type ISysAgent interface {
	List(ctx context.Context, req *system.AgentListReq) (res *system.AgentListRes, err error)
	Get(ctx context.Context, req *system.AgentGetReq) (res *system.AgentGetRes, err error)
}

var localSysAgent ISysAgent

func SysAgent() ISysAgent {
	if localSysAgent == nil {
		panic("implement not found for interface ISysAgent, forgot register?")
	}
	return localSysAgent
}

func RegisterSysAgent(i ISysAgent) {
	localSysAgent = i
}
