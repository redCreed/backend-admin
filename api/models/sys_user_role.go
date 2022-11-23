package models

type SysUserRole struct {
	Id     int `json:"id" gorm:"primaryKey;autoIncrement;comment:id"` //id
	UserId int `json:"user_id" gorm:"size:11;comment:用户id"`           //用户id
	RoleId int `json:"role_id" gorm:"size:11;comment:角色id"`           //角色id
	//Status int `json:"status"  gorm:"size:11;comment:状态 -1:删除 1:正常 2:停用"` //状态 -1:删除 1:正常 2:停用
	//ControlBy
	ModelTime
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
