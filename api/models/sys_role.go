package models

type SysRole struct {
	RoleId   int    `json:"roleId" gorm:"primaryKey;autoIncrement"` //角色编码
	RoleName string `json:"roleName" gorm:"size:128;"`              //角色名称
	Status   int    `json:"status" gorm:"type:varchar(4);comment:状态 -1:删除 1:正常 2:停用 "`
	RoleKey  string `json:"roleKey" gorm:"size:128;"` //角色唯一key 不能修改
	RoleSort int    `json:"roleSort" gorm:""`         //角色排序
	Flag     string `json:"flag" gorm:"size:128;"`    //标志位
	Remark   string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	Admin    bool   `json:"admin" gorm:"size:4;"` //是否是管理员角色
	ControlBy
	ModelTime
}

func (SysRole) TableName() string {
	return "sys_role"
}
