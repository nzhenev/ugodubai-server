package system

import (
	commonApi "ugodubai-server/api/v1/common"
	"ugodubai-server/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// LoginLogListReq 查询列表请求参数
type LoginLogListReq struct {
	g.Meta        `path:"/loginLog/list" tags:"登录日志管理" method:"get" summary:"日志列表"`
	LoginName     string `p:"userName"`      //登陆名
	Status        string `p:"status"`        //状态
	Ipaddr        string `p:"ipaddr"`        //登录地址
	SortName      string `p:"orderByColumn"` //排序字段
	SortOrder     string `p:"isAsc"`         //排序方式
	LoginLocation string `p:"loginLocation"` //登录地点
	commonApi.PageReq
}

type LoginLogListRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*entity.SysLoginLog `json:"list"`
}

type LoginLogDelReq struct {
	g.Meta `path:"/loginLog/delete" tags:"登录日志管理" method:"delete" summary:"删除日志"`
	Ids    []int `p:"ids" v:"required#ids必须"`
}

type LoginLogDelRes struct {
}

type LoginLogClearReq struct {
	g.Meta `path:"/loginLog/clear" tags:"登录日志管理" method:"delete" summary:"清除日志"`
}

type LoginLogClearRes struct {
}
