package models

//SysCasbinRule 权限规则表
type SysCasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index;comment:规则类型"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index;comment:角色ID"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index;comment:api路径"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index;comment:api访问方法""`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
}

func (SysCasbinRule) TableName() string {
	return "sys_casbin_rule"
}

/*
pType有2中p和g
	例如:
	p admin    /user  GET   p是策略 admin是角色 /user是资源路径  GET是请求方式

	g zhangsan admin        g是具体某个用户策略 zhangsan是张三  admin是所属角色

*/
