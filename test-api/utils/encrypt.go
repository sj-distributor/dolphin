package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/duke-git/lancet/v2/validator"
	"github.com/maiguangyang/aescrypto"
)

// aes密钥
var secretKey string = "dAcGmFaAgBaAnHbOcLdIbCiUbIuAfFbA"

// 目前前端密码，使用EncryptPassword相同加密方法传过来，再次加密
// 密码加密
func EncryptPassword(text string) string {
	password := EncryptMd5(text)
	password = base64.StdEncoding.EncodeToString([]byte(password))
	password = EncryptMd5(password)
	return password
}

// MD5验证
func DecryptPassword(text, md5Text string) bool {
	password := EncryptMd5(text)
	password = base64.StdEncoding.EncodeToString([]byte(password))
	password = EncryptMd5(password)
	return password == md5Text
}

// MD5加密
func EncryptMd5(text string) string {
	md5Data := fmt.Sprintf("%x", md5.Sum([]byte(text)))
	return md5Data
}

// aes解密
func AesLeftDecrypt(text string, key string) (string, error) {
	if validator.IsEmptyString(key) {
		key = secretKey
	}
	// 先base64解密
	prvPem, err := base64.StdEncoding.DecodeString(text)

	// 再使用hex.DecodeString转换
	dexDecode, err := hex.DecodeString(string(prvPem))

	if err != nil {
		return "", err
	}

	// 最后解密对称加密的内容
	prvPem, err = AesDecrypt(dexDecode, []byte(key))
	if err != nil {
		return "", err
	}

	return string(prvPem), nil
}

// AES CBC加密
func AesCbcEncrypt(text string, key string) string {
	if validator.IsEmptyString(key) {
		key = secretKey
	}
	cryptText, err := aescrypto.AesCbcPkcs7Encrypt([]byte(text), []byte(key), nil)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(cryptText)
}

// AES CBC解密
func AesCbcDecrypt(text string, key string) string {
	if validator.IsEmptyString(key) {
		key = secretKey
	}

	cryptText, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}
	fmt.Println("cryptText", cryptText)

	decData, err := aescrypto.AesCbcPkcs7Decrypt(cryptText, []byte(key), nil)

	if err != nil {
		return ""
	}

	return string(decData)
}
