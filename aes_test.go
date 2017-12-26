package utils

import (
	"fmt"
	"testing"
)

func Test_AesEncrypt(t *testing.T) {
	data := `{"account": "boochat","password": "boochat", "mcc": 86}`
	encoded, err := AesEncrypt([]byte(data))
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf("%s \n", encoded)
	// out: Mkm9c869ZRD0nhdNZ0zjvJa7UrZQZdUPGhtIQp6ujKG4RBLth3wNrp5k8gj65ecWRNnLySaBIP2EwOLZ/rr4Pw==
}

func Test_AesDecrypt(t *testing.T) {
	data := "Mkm9c869ZRD0nhdNZ0zjvJa7UrZQZdUPGhtIQp6ujKG4RBLth3wNrp5k8gj65ecWRNnLySaBIP2EwOLZ/rr4Pw=="
	decoded, err := AesDecrypt(data)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf("%s \n", decoded)
	// out: {"account": "boochat","password": "boochat", "mcc": 86}
}
