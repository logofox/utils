package utils

import (
	"fmt"
	"testing"
)

func Test_Md5(t *testing.T) {
	data := "logo_fox"
	fmt.Println("Test_Md5:", GetMd5(data, false))
	fmt.Println("Test_Md5 short:", GetMd5(data, true))
}

func Test_Bitwise(t *testing.T) {
	fmt.Println("Test_Bitwise [foobar@example.com]:", Bitwise("foobar@example.com"))
	fmt.Println("Test_Bitwise:", Bitwise("12"))

	fmt.Println("Test_Bitwise:", Bitwise("z"))
	fmt.Println("Test_Bitwise:", Bitwise("zf"))

	fmt.Println("Test_Bitwise:", Bitwise("logo"))
	fmt.Println("Test_Bitwise:", Bitwise("logof"))
	fmt.Println("Test_Bitwise:", Bitwise("logofox"))

	fmt.Println("Test_Bitwise:", Bitwise("some string"))
}

func Test_BinaryTransfer(t *testing.T) {
	fmt.Println("Test_GetBooId [-1007577910]:", BinaryTransfer(-1007577910, 61))
	fmt.Println("Test_GetBooId:", BinaryTransfer(1097329740, 61))
}

func Test_ShortHash(t *testing.T) {
	fmt.Println("Test_ShortHash [logofox]:", ShortHash("logofox"))
	fmt.Println("Test_ShortHash [foobar@example.com]:", ShortHash("foobar@example.com"))
	fmt.Println("Test_ShortHash [my name is really big big and big...]:", ShortHash("my name is really big big and big..."))
	fmt.Println("Test_ShortHash [万里长城永不倒。。。]:", ShortHash("万里长城永不倒。。。"))
	fmt.Println("Test_ShortHash [和平]:", ShortHash("和平"))
}
