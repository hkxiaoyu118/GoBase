package main

import (
	"fmt"
	"gobase/AesCrypter"
)

func main() {
	//result :=BaseString.GetRandomString(16)
	//fmt.Println(result)
	//
	//var str string="123123123sjdflksjldf"
	//result2:=BaseString.GetRandStringEx(16,str)
	//fmt.Println(result2)

	str := "this is a test"
	key := "1234567890123456"
	result := AesCrypter.AESCbcEncrypt([]byte(str), []byte(key))
	fmt.Println(result)
	result2:=AesCrypter.AESCbcDecrypt(result, []byte(key))
	fmt.Println(string(result2))
}
