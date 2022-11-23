package models

import (
	"back-admin/internal/app/common/driver"
	"back-admin/pkg/queue"
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

type SysLoginLog struct {
	Model
	Username      string    `json:"username" gorm:"type:varchar(128);comment:用户名"`
	Status        int       `json:"status" gorm:"type:varchar(4);comment:状态 -1:删除 1:正常 2:停用 "`
	Ipaddr        string    `json:"ipaddr" gorm:"type:varchar(255);comment:ip地址"`
	LoginLocation string    `json:"loginLocation" gorm:"type:varchar(255);comment:归属地"`
	Browser       string    `json:"browser" gorm:"type:varchar(255);comment:浏览器"`
	Os            string    `json:"os" gorm:"type:varchar(255);comment:系统"`
	Platform      string    `json:"platform" gorm:"type:varchar(255);comment:固件"`
	LoginTime     time.Time `json:"loginTime" gorm:"type:timestamp;comment:登录时间"`
	Remark        string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	Msg           string    `json:"msg" gorm:"type:varchar(255);comment:信息"`
	CreatedAt     time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	ControlBy
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}

// SaveLoginLog 从队列中获取登录日志
func SaveLoginLog(message []queue.Messager) error {
	var err error
	api := make([]SysLoginLog, 0)
	for _, v := range message {
		temp := new(SysLoginLog)
		if err = json.Unmarshal(v.GetValues(), temp); err != nil {
			driver.Instance.Log.Error("解析登录日志错误", zap.Error(err))
			continue
		}
		api = append(api, *temp)
	}
	if len(api) > 0 {
		err = driver.Instance.Orm.Create(&api).Error
		if err != nil {
			driver.Instance.Log.Error("保存登录日志错误", zap.Error(err))
			return err
		}
	}
	return nil
}
