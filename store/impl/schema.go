package impl

import (
	"back-admin/api/models"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type schema struct {
	gorm *gorm.DB
}

func (d *schema) CreateRoot() error {
	return nil
}

func (d *schema) CreateSchema() error {
	tx := d.gorm.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", models.Schema))
	if tx.Error != nil {
		return errors.WithStack(tx.Error)
	}

	return nil
}

func (d *schema) HasTable(tb interface{}) bool {
	return d.gorm.Migrator().HasTable(tb)
}

func (d *schema) Close() error {
	db, err := d.gorm.DB()
	if err != nil {
		return errors.WithStack(err)
	}

	if err = db.Close(); err != nil {
		return errors.Wrap(err, "close db err")
	}

	return nil
}

func (d *schema) Migrate(tables []interface{}) error {
	if err := d.gorm.AutoMigrate(tables...); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d *schema) Exec(s string) error {
	if err := d.gorm.Exec(s).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
