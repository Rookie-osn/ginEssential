package util

import (
	"math/rand"
	"time"
)

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
