package utils

import (
	"crypto/md5"
	"fmt"
)

func GetMd5(data string, isShort bool) string {
	md5 := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	if isShort {
		return md5[8:24]
	}
	return md5
}

// 字符串解析成10进制
func Bitwise(data string) int32 {
	var hash int32
	for _, v := range []rune(data) {
		hash = ((hash << 5) - hash) + int32(v)
		hash = hash & hash
	}
	return hash
}

// 10进制转换为任意进制
func BinaryTransfer(interger, bitbase int32) (result string) {
	isMinus := false
	basechar := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if bitbase > 62 || bitbase < 0 {
		return ""
	}
	if interger < 0 {
		interger = interger * (-1)
		isMinus = true
	}

	for interger != 0 {
		char := basechar[interger%bitbase]
		interger = interger / bitbase
		result = string(char) + result
	}

	// 讲负号转换为 ’Z‘
	if isMinus {
		return "Z" + result
	}
	return result
}

func ShortHash(username string) string {
	if username == "" {
		return ""
	}
	imid := fmt.Sprintf("%s", username) // 原生写法
	return BinaryTransfer(Bitwise(imid), 61)
}

