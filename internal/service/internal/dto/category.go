// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Category is the golang structure of table gf_category for DAO operations like Where/Data.
type Category struct {
	g.Meta      `orm:"table:gf_category, dto:true"`
	Id          interface{} // 分类ID，自增主键
	ContentType interface{} // 内容类型：topic, ask, article, reply
	Key         interface{} // 栏目唯一键名，用于程序部分场景硬编码，一般不会用得到
	ParentId    interface{} // 父级分类ID，用于层级管理
	UserId      interface{} // 创建的用户ID
	Name        interface{} // 分类名称
	Sort        interface{} // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Thumb       interface{} // 封面图
	Brief       interface{} // 简述
	Content     interface{} // 详细介绍
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
