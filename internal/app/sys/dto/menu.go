package dto

import "back-admin/api/models"

type AddMenuParam struct {
	MenuName string `form:"menu_name" json:"menu_name" validate:"required"`             //菜单名称
	Title    string `form:"title" json:"title"`                                         //菜单标题
	Icon     string `form:"icon" json:"icon" `                                          //菜单图标
	Path     string `form:"path" json:"path" `                                          //菜单路径
	MenuType int    `form:"menu_type" json:"menu_type" validate:"required,oneof=1 2 3"` //菜单类型 1:目录 2:菜单 3:按钮
	ParentId int    `form:"parent_id" json:"parent_id" `                                //父级菜单
	Sort     int    `form:"sort" json:"sort" validate:"gte=0"`                          //排序
	IsShow   int    `form:"is_show" json:"is_show" `                                    //是否显示1:正常 0:停用
	Remark   string `form:"remark" json:"remark" `                                      //备注
	Status   int    `form:"status" json:"status"  validate:"required,oneof=1 2"`        //状态 -1:删除 1:正常 2:停用
	Apis     []int  `form:"apis" json:"apis" `                                          //api接口id集合
}

type UpdateMenuParam struct {
	MenuId   int     `form:"menu_id" json:"menu_id"`                                     //id
	MenuName string  `form:"menu_name" json:"menu_name" validate:"required"`             //菜单名称
	Title    *string `form:"title" json:"title"`                                         //菜单标题
	Icon     *string `form:"icon" json:"icon" `                                          //菜单图标
	Path     *string `form:"path" json:"path" `                                          //菜单路径
	MenuType int     `form:"menu_type" json:"menu_type" validate:"required,oneof=1 2 3"` //菜单类型 1:目录 2:菜单 3:按钮
	ParentId int     `form:"parent_id" json:"parent_id" validate:"gte=0"`                //父级菜单
	Sort     *int    `form:"sort" json:"sort" validate:"omitempty,gte=0"`                //排序
	IsShow   *int    `form:"is_show" json:"is_show" `                                    //是否显示1:正常 0:停用
	Remark   *string `form:"remark" json:"remark" `                                      //备注
	Status   int     `form:"status" json:"status"  validate:"required,oneof=1 2"`        //状态 -1:删除 1:正常 2:停用
	Apis     []int   `form:"apis" json:"apis"  `                                         //api接口id集合
}

type MenuTree struct {
	MenuId   int             `json:"menuId" gorm:"primaryKey;autoIncrement"`            //id
	MenuName string          `json:"menuName" gorm:"size:128;"`                         //名称
	Title    string          `json:"title" gorm:"size:128;"`                            //标题
	Icon     string          `json:"icon" gorm:"size:128;"`                             //图标
	Path     string          `json:"path" gorm:"size:128;"`                             //路径
	MenuType int             `json:"menuType" gorm:"size:1;"`                           //类型 1:目录 2:菜单 3:按钮
	ParentId int             `json:"parentId" gorm:"size:11;"`                          //父id
	IdPath   string          `json:"id_path" gorm:"type:varchar(255);comment:父级id全路径"`  //父级id全路径  /0/1/
	Sort     int             `json:"sort" gorm:"size:4;"`                               //排序
	IsShow   int             `json:"is_show" gorm:"size:4;"`                            //是否显示1:正常 0:停用
	Status   int             `json:"status"  gorm:"size:11;comment:状态 -1:删除 1:正常 2:停用"` //状态 -1:删除 1:正常 2:停用
	Remark   string          `json:"remark" gorm:"type:varchar(255);comment:备注"`        //备注
	Children []MenuTree      `json:"children"`
	Api      []models.SysApi `json:"api"`
}
