package main

import (
	"../AesCrypter"
	"fmt"
)

func main() {
	//result :=BaseString.GetRandomString(16)
	//fmt.Println(result)
	//
	//var str string="123123123sjdflksjldf"
	//result2:=BaseString.GetRandStringEx(16,str)
	//fmt.Println(result2)

	str := "this is a test";
	key := "1234567890123456"
	result := AesCrypter.AESCtrEncryptV2([]byte(str), []byte(key))
	fmt.Println(result)
	result2:=AesCrypter.AESCtrDecryptV2(result, []byte(key))
	fmt.Println(string(result2))
}
