package models

type SysMenuApi struct {
	Id     int `json:"id" gorm:"primaryKey;autoIncrement;comment:id"` //id
	MenuId int `json:"menu_id" gorm:"size:11;comment:菜单id"`           //菜单id
	ApiId  int `json:"api_id" gorm:"size:11;comment:action id"`       //action id
	//Status int `json:"status"  gorm:"size:11;comment:状态 -1:删除 1:正常 2:停用"` //状态 -1:删除 1:正常 2:停用
	//ControlBy
	ModelTime
}

func (SysMenuApi) TableName() string {
	return "sys_menu_api"
}
