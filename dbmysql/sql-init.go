package dbmysql

import (
	"context"
	"fmt"
	"log"
	"time"
	"yana-golang/db_interface"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct { // conform
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n===============\n", sql)
}

func Init() *db_interface.MYSQL {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.schema"),
	)
	dial := mysql.Open(dsn)
	gormClient, err := gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		log.Panic("init gorm")
	}
	return &db_interface.MYSQL{
		User: InitUserDB(gormClient),
	}
}
