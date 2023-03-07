package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/isso-719/gaya-on-server/pkg/config"
	"github.com/jinzhu/gorm"
)

type SQLConnector struct {
	Conn *gorm.DB
}

func NewSQLConnector() *SQLConnector {
	conf := config.LoadConfig()

	dsn := sqlConnInfo(*conf.SQLInfo)
	conn, err := gorm.Open(conf.SQLInfo.DBType, dsn)
	if err != nil {
		panic(err)
	}

	// Connector 作成時に自動でマイグレーションを実行する
	conn.AutoMigrate()

	return &SQLConnector{
		Conn: conn,
	}
}

func sqlConnInfo(sqlInfo config.SQLInfo) string {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		sqlInfo.SQLUser,
		sqlInfo.SQLPassword,
		sqlInfo.SQLAddress,
		sqlInfo.SQLDBName,
	)

	return dataSourceName
}

func (sc *SQLConnector) Find(out interface{}, where ...interface{}) *gorm.DB {
	return sc.Conn.Find(out, where...)
}

func (sc *SQLConnector) Exec(sql string, values ...interface{}) *gorm.DB {
	return sc.Conn.Exec(sql, values...)
}

func (sc *SQLConnector) Raw(sql string, values ...interface{}) *gorm.DB {
	return sc.Conn.Raw(sql, values...)
}

func (sc *SQLConnector) Create(value interface{}) *gorm.DB {
	return sc.Conn.Create(value)
}

func (sc *SQLConnector) Save(value interface{}) *gorm.DB {
	return sc.Conn.Save(value)
}

func (sc *SQLConnector) Delete(value interface{}, where ...interface{}) *gorm.DB {
	return sc.Conn.Delete(value, where...)
}

func (sc *SQLConnector) Where(query interface{}, args ...interface{}) *gorm.DB {
	return sc.Conn.Where(query, args...)
}
