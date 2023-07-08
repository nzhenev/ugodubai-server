// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure for table sys_post.
type SysPost struct {
	PostId    uint64      `json:"postId"    description:"岗位ID"`
	PostCode  string      `json:"postCode"  description:"岗位编码"`
	PostName  string      `json:"postName"  description:"岗位名称"`
	PostSort  int         `json:"postSort"  description:"显示顺序"`
	Status    uint        `json:"status"    description:"状态（0正常 1停用）"`
	Remark    string      `json:"remark"    description:"备注"`
	CreatedBy uint64      `json:"createdBy" description:"创建人"`
	UpdatedBy uint64      `json:"updatedBy" description:"修改人"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
