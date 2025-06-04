package mysql

import (
	"api/configs"
	"api/internal/models"
	"fmt"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
)

func Setup() {
	config := configs.App.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Database, config.Charset)
	color.Cyan("** dns:" + dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Panic("** initialize database failed, err: "+err.Error(), zap.String("error", "初始化数据库失败!"+err.Error()))
		return
	}
	Database = db
	models.InitializeTables(Database)
}

func Close() {
}
