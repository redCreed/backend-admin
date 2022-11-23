package models

type SysMenu struct {
	MenuId   int    `json:"menuId" gorm:"primaryKey;autoIncrement"`            //id
	MenuName string `json:"menuName" gorm:"size:128;"`                         //名称
	Title    string `json:"title" gorm:"size:128;"`                            //标题
	Icon     string `json:"icon" gorm:"size:128;"`                             //图标
	Path     string `json:"path" gorm:"size:128;"`                             //路径
	MenuType int    `json:"menuType" gorm:"size:1;"`                           //类型 1:目录 2:菜单 3:按钮
	ParentId int    `json:"parentId" gorm:"size:11;"`                          //父id
	IdPath   string `json:"id_path" gorm:"type:varchar(255);comment:父级id全路径"`  //父级id全路径  /0/1/
	Sort     int    `json:"sort" gorm:"size:4;"`                               //排序
	IsShow   int    `json:"is_show" gorm:"size:4;"`                            //是否显示1:正常 0:停用
	Status   int    `json:"status"  gorm:"size:11;comment:状态 -1:删除 1:正常 2:停用"` //状态 -1:删除 1:正常 2:停用
	Remark   string `json:"remark" gorm:"type:varchar(255);comment:备注"`        //备注
	ControlBy
	ModelTime
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
