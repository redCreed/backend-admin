package models

type SysApi struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Handle string `json:"handle" gorm:"size:128;comment:handle"`
	Title  string `json:"title" gorm:"size:128;comment:标题"`
	Path   string `json:"path" gorm:"size:128;comment:地址"`
	Method string `json:"method" gorm:"size:16;comment:请求方式"`
	ModelTime
	ControlBy
}

func (SysApi) TableName() string {
	return "sys_api"
}
