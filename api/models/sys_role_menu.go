package models

type SysRoleMenu struct {
	Id     int `json:"id" gorm:"primaryKey;autoIncrement;comment:id"` //id
	RoleId int `json:"role_id" gorm:"size:11;comment:角色id"`           //角色id
	MenuId int `json:"menu_id" gorm:"size:11;comment:菜单id"`           //菜单id
	//Status int `json:"status"  gorm:"size:11;comment:状态 -1:删除 1:正常 2:停用"` //状态 -1:删除 1:正常 2:停用
	//ControlBy
	ModelTime
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
