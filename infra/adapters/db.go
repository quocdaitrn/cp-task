package adapters

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/quocdaitrn/cp-task/infra/config"
)

func ProvideMySQL(cfg config.Config) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(cfg.DBDsn), &gorm.Config{})
	return
}
