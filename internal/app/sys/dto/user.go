package dto

import (
	"time"
)

type UserListParam struct {
	PageNo   int `form:"page_no" json:"page_no"  validate:"required,min=1" `          //页面
	PageSize int `form:"page_no" json:"page_size" validate:"required,min=1,max=200" ` //页大小
}

type UserListResp struct {
	Data  []SysUserList `json:"data"`  //数据集合
	Count int64         `json:"count"` //数据总数
}

type SysUserList struct {
	UserId    int        `json:"user_id"`    //用户id
	Username  string     `json:"username" `  //用户名(唯一且不能修改)
	NickName  string     `json:"nick_name" ` //昵称
	Phone     string     `json:"phone" `     //手机号
	Avatar    string     `json:"avatar" `    //头像
	Sex       string     `json:"sex" `       //性别
	Email     string     `json:"email"`      //邮箱
	Remark    string     `json:"remark" `    //备注
	Status    int        `json:"status" `    //状态 -1:删除 1:正常 2:停用
	CreateBy  int        `json:"createBy" `  //创建者
	UpdateBy  int        `json:"updateBy" `  //更新者
	CreatedAt time.Time  `json:"createdAt" ` //创建时间
	UpdatedAt time.Time  `json:"updatedAt" ` //最后更新时间
	Roles     []RoleInfo `json:"roles"`      //角色集合
}

type RoleInfo struct {
	UserId   int    `json:"user_id"`   //用户id
	RoleId   int    `json:"role_id"`   //角色id
	RoleName string `json:"role_name"` //角色名称
}

type UserLoginReq struct {
	Username string `form:"username" json:"username"  validate:"required"` //用户名称
	Password string `form:"password" json:"password"  validate:"required"` //用户密码
}

type UserLoginResp struct {
	//UserId   int    `json:"user_id"`   //用户id
	//Nickname string `json:"nickname"`  //用户昵称
	//RoleId   int    `json:"role_id"`   //角色id
	//RoleName string `json:"role_name"` //角色名称
	Token string `json:"token"` //登录token
}

type UserToken struct {
	Token string `json:"token"` //登录token
}

type UpdateUserRoleReq struct {
	RoleId []int `form:"role_id" json:"role_id"  validate:"gte=0,lte=200,unique"` //角色id集合
}

type AddUserReq struct {
	Username string `form:"username" json:"username"   validate:"required"`                   //用户名(唯一且不能修改)
	Password string `form:"password" json:"password" validate:"required"`                     //密码
	NickName string `json:"nick_name" `                                                       //昵称
	Phone    string `json:"phone" `                                                           //手机号
	Avatar   string `json:"avatar" `                                                          //头像
	Sex      string `json:"sex" `                                                             //性别
	Email    string `json:"email"`                                                            //邮箱
	Remark   string `json:"remark" `                                                          //备注
	RoleId   []int  `form:"role_id" json:"role_id"  validate:"required,gte=0,lte=200,unique"` //角色id集合
}

type UpdateUserReq struct {
	Password string `json:"password" `                                                        //密码
	NickName string `json:"nick_name" `                                                       //昵称
	Phone    string `json:"phone" `                                                           //手机号
	Avatar   string `json:"avatar" `                                                          //头像
	Sex      string `json:"sex" `                                                             //性别
	Email    string `json:"email"`                                                            //邮箱
	Remark   string `json:"remark" `                                                          //备注
	RoleId   []int  `form:"role_id" json:"role_id"  validate:"required,gte=0,lte=200,unique"` //角色id集合
}
