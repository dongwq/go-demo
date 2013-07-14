package demo

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/dongwq/godemo/util"
)

func AesEncode(key, src []byte) []byte {

	// 初始向量，长度必须为16个字节(128bit)
	var iv = []byte("abcdef1234567890")[:aes.BlockSize]
	// 得到块，用于加密和解密
	block, err := aes.NewCipher(key)

	util.CheckErr("msg err", err)

	// 加密，使用CFB模式(密文反馈模式)，其他模式参见crypto/cipher
	encrypter := cipher.NewCFBEncrypter(block, iv)

	encrypted := make([]byte, len(src))
	encrypter.XORKeyStream(encrypted, src)

	return encrypted

}

func AesDecode(key, src []byte) []byte {

	// 初始向量，长度必须为16个字节(128bit)
	var iv = []byte("abcdef1234567890")[:aes.BlockSize]
	// 得到块，用于加密和解密
	block, err := aes.NewCipher(key)
	util.CheckErr("msg err", err)

	// 解密
	decrypter := cipher.NewCFBDecrypter(block, iv)

	decrypted := make([]byte, len(src))
	decrypter.XORKeyStream(decrypted, src)

	return decrypted
}

func AesDecodeStr(key, src string) string {
	srcBytes,err := base64.StdEncoding.DecodeString(src)
	util.CheckErr("msg err", err)
	
	result := AesDecode([]byte(key), srcBytes)
	return string(result)
}

func AesEncodeStr(key, src string) string {
	result := AesEncode([]byte(key), []byte(src))

	return base64.StdEncoding.EncodeToString(result)
}

func DemoAES() {
	// 消息明文
	src := []byte("hello, world, what")
	// 密钥，长度可以为16、24、32字节
	key := []byte("1234567890abcdef")

	encrypted := AesEncode(key, src)

	fmt.Printf("Encrypting %s : %v -> %v\n", src, []byte(src), encrypted)

	// 解密
	//decrypter := cipher.NewCFBDecrypter(block, iv)

	//decrypted := make([]byte, len(src))
	//decrypter.XORKeyStream(decrypted, encrypted)
	//fmt.Printf("Decrypting %v -> %v : %s\n", encrypted, decrypted, decrypted)
}

func DemoAESStr() {
	// 消息明文
	src := "hello, world,what"
	// 密钥，长度可以为16、24、32字节
	key := "1234567890abcdef"

	encrypted := AesEncodeStr(key, src)

	fmt.Printf("Encrypting %s : %v -> %v\n", src, []byte(src), encrypted)

	decrypted := AesDecodeStr(key, encrypted)

	fmt.Printf("Decrypting %v -> %v : %s\n", encrypted, decrypted, decrypted)
}
