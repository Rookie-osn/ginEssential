package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math/rand"
	"net/http"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11),not null;unique"`
	Password string `gorm:"size:255;not null"'`
}

func main(){
	db:=InitDB()
	defer db.Close()

	r:=gin.Default()
	r.POST("/api/auth/register",func(ctx *gin.Context){
		// 获取参数
		name:=ctx.PostForm("name")
		telephone:=ctx.PostForm("telephone")
		password:=ctx.PostForm("password")
		// 数据验证
		if len(telephone)!=11{
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"手机号必须为11位"})
			return
		}
		if len(password)<6{
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
			return
		}
		// 如果名称没有传，给一个10位的随机字符串
		if len(name)==0{
			name=RandomString(10)
		}
		// 判断手机号是否存在
		if isTelephoneExist(db,telephone){
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户已经存在"})
			return
		}
		// 创建用户
		newUser:=User{
			Name:name,
			Telephone: telephone,
			Password: password,
		}
		db.Create(&newUser)
		// 返回结果
		ctx.JSON(200,gin.H{"msg":"注册成功"})
	})

	r.Run(":8080")
}

// 查询手机号是否已存在
func isTelephoneExist(db *gorm.DB,telephone string)bool{
	var user User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID!=0{
		return true
	}
	return false
}

// 生成长度为n的随机字符串
func RandomString(n int)string{
	var letters=[]byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result:=make([]byte,n)

	rand.Seed(time.Now().Unix())
	for i:=range result{
		result[i]=letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 连接数据库
func InitDB()*gorm.DB{
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
	db.AutoMigrate(&User{})

	return db
}