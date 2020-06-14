package Boot

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


//mysql相关
var mysql_db *gorm.DB
func init(){
	InitMysql()
}

func InitMysql() error {
	var err error
	mysql_db, err = gorm.Open("mysql", "root:root@tcp(192.168.2.239:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		mysql_db=nil
		 return NewFatalError(err.Error()) //这里返回致命异常
	}
	mysql_db.SingularTable(true)
	mysql_db.DB().SetMaxIdleConns( 10)
	mysql_db.DB().SetMaxOpenConns( 100)
	mysql_db.LogMode(true)


	return nil
}

func GetDB() *gorm.DB {
	return mysql_db
}