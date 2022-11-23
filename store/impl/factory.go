package impl

import (
	"back-admin/store"
	"gorm.io/gorm"
)

type Factory struct {
	gorm *gorm.DB
}

func NewFactory(gorm *gorm.DB) store.Db {
	return &Factory{gorm: gorm}
}

func (f *Factory) Sys() store.Sys {
	return &sys{gorm: f.gorm}
}

func (f *Factory) Schema() store.Schema {
	return &schema{gorm: f.gorm}
}
