// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Setting is the golang structure of table gf_setting for DAO operations like Where/Data.
type Setting struct {
	g.Meta    `orm:"table:gf_setting, dto:true"`
	K         interface{} // 键名
	V         interface{} // 键值
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
