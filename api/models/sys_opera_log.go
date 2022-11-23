package models

import (
	"back-admin/internal/app/common/driver"
	"back-admin/pkg/queue"
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

type SysOperaLog struct {
	Model
	Title         string    `json:"title" gorm:"type:varchar(255);comment:操作模块"`
	BusinessType  string    `json:"businessType" gorm:"type:varchar(128);comment:操作类型"`
	BusinessTypes string    `json:"businessTypes" gorm:"type:varchar(128);comment:BusinessTypes"`
	Method        string    `json:"method" gorm:"type:varchar(128);comment:函数"`
	RequestMethod string    `json:"requestMethod" gorm:"type:varchar(128);comment:请求方式: GET POST PUT DELETE"`
	OperatorType  string    `json:"operatorType" gorm:"type:varchar(128);comment:操作类型"`
	OperName      string    `json:"operName" gorm:"type:varchar(128);comment:操作者"`
	DeptName      string    `json:"deptName" gorm:"type:varchar(128);comment:部门名称"`
	OperUrl       string    `json:"operUrl" gorm:"type:varchar(255);comment:访问地址"`
	OperIp        string    `json:"operIp" gorm:"type:varchar(128);comment:客户端ip"`
	OperLocation  string    `json:"operLocation" gorm:"type:varchar(128);comment:访问位置"`
	OperParam     string    `json:"operParam" gorm:"type:text;comment:请求参数"`
	Status        int       `json:"status" gorm:"type:varchar(4);comment:状态 -1:删除 1:正常 2:停用 "`
	OperTime      time.Time `json:"operTime" gorm:"type:timestamp;comment:操作时间"`
	JsonResult    string    `json:"jsonResult" gorm:"type:varchar(255);comment:返回数据"`
	Remark        string    `json:"remark" gorm:"type:varchar(255);comment:备注"`
	LatencyTime   string    `json:"latencyTime" gorm:"type:varchar(128);comment:耗时"`
	UserAgent     string    `json:"userAgent" gorm:"type:varchar(255);comment:ua"`
	CreatedAt     time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	ControlBy
}

func (SysOperaLog) TableName() string {
	return "sys_opera_log"
}

// SaveOperaLog 从队列中获取操作日志日志
func SaveOperaLog(message []queue.Messager) error {
	var err error
	api := make([]SysOperaLog, 0)
	for _, v := range message {
		temp := new(SysOperaLog)
		if err = json.Unmarshal(v.GetValues(), temp); err != nil {
			driver.Instance.Log.Error("解析操作日志错误", zap.Error(err))
			continue
		}
		api = append(api, *temp)
	}
	if len(api) > 0 {
		err = driver.Instance.Orm.Create(&api).Error
		if err != nil {
			driver.Instance.Log.Error("保存操作日志错误", zap.Error(err))
			return err
		}
	}
	return nil
}
