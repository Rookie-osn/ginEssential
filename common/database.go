package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"rookie.code/ginessential/model"
)
var DB *gorm.DB

// 连接数据库
func init(){
	driverName:="mysql"
	host:="localhost"
	port:="3306"
	database:="ginessential"
	username:="root"
	password:="believe9407"
	charset:="utf8"
	args:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,password,host,port,database,charset)
	//	连接数据库
	db,err:=gorm.Open(driverName,args)
	if err!=nil{
		panic("failed to connect database,err: "+err.Error())
	}
	// 根据传入的结构体类型自动生成对应参数的表
	db.AutoMigrate(&model.User{})
	DB=db
}

// 提供给其他包调用返回DB实例
func GetDB() *gorm.DB{
	return DB
}