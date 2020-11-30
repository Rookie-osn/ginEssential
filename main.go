package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"rookie.code/ginessential/common"
)



func main(){
	DB:=common.GetDB()
	defer DB.Close()

	r:=gin.Default()
	r=CollectRouter(r)
	panic(r.Run())
}