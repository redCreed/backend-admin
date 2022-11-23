package dto

import "back-admin/api/models"

type RoleListParam struct {
	PageNo   int `form:"page_no" json:"page_no"  validate:"required,min=1" `            //页面
	PageSize int `form:"page_size" json:"page_size" validate:"required,min=1,max=200" ` //页大小
}

type RoleListResp struct {
	Data  []models.SysRole `json:"data"`  //数据集合
	Count int64            `json:"count"` //数据总数
}

type RoleList struct {
	RoleId   int    `form:"role_id" json:"role_id" gorm:"primaryKey;autoIncrement"` // 角色编码
	RoleName string `form:"role_name" json:"role_name" gorm:"size:128;"`            // 角色名称
	Status   string `form:"status" json:"status" gorm:"size:4;"`                    //状态 1:正常 2:禁用 -1:删除
	RoleKey  string `form:"role_key" json:"role_key" gorm:"size:128;"`              //角色唯一key 不能修改
	RoleSort int    `form:"role_sort" json:"role_sort" gorm:""`                     //角色排序
	Flag     string `form:"flag" json:"flag" gorm:"size:128;"`                      //标志位
	Remark   string `form:"remark" json:"remark" gorm:"size:255;"`                  //备注
	Admin    bool   `form:"admin" json:"admin" gorm:"size:4;"`                      //是否root
}

type AddRoleParam struct {
	RoleName string `form:"role_name" json:"role_name" validate:"required" ` //角色名称
	RoleKey  string `form:"role_key" json:"role_key"  validate:"required" `  //角色唯一key
	RoleSort int    `form:"role_sort" json:"role_sort" validate:"gte=0"`     //角色排序
	Flag     string `form:"flag" json:"flag" `                               //标志位
	Remark   string `form:"remark" json:"remark" `                           //备注
	MenuIds  []int  `form:"menu_ids" json:"menu_ids"  validate:"required"`   //菜单id集合
}

type UpdateRoleParam struct {
	RoleId   int    `form:"role_id" json:"role_id" `                         // 角色编码
	RoleName string `form:"role_name" json:"role_name" validate:"required" ` //角色名称
	//RoleKey  *string `form:"role_key" json:"role_key"  validate:"required" `  //角色key不能修改
	RoleSort *int    `form:"role_sort" json:"role_sort" `                   //角色排序
	Flag     *string `form:"flag" json:"flag" `                             //标志位
	Remark   *string `form:"remark" json:"remark" `                         //备注
	MenuIds  []int   `form:"menu_ids" json:"menu_ids"  validate:"required"` //菜单id集合
}
