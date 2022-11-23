package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string `json:"username" gorm:"type:varchar(64);comment:用户名"` //不能修改
	Password string `json:"-" gorm:"type:varchar(128);comment:密码"`
	NickName string `json:"nick_name" gorm:"type:varchar(128);comment:昵称"`
	Phone    string `json:"phone" gorm:"type:varchar(11);comment:手机号"`
	Salt     string `json:"-" gorm:"type:varchar(255);comment:加盐"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255);comment:头像"`
	Sex      string `json:"sex" gorm:"type:varchar(255);comment:性别"`
	Email    string `json:"email" gorm:"type:varchar(128);comment:邮箱"`
	Remark   string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	Status   int    `json:"status" gorm:"size:2;comment:状态 -1:删除 1:正常 2:停用 "`
	ControlBy
	ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}

// Encrypt 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}
