package wxshop_go_sdk

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"sort"
	"strings"
)

// VerifyGet 绑定消息通知接口验证时候用
func VerifyGet(timeStamp, nonce, token string) string {
	//mergeStr := fmt.Sprintf("%s%s%s", timeStamp, nonce, token)
	//// 1. 创建 SHA1 哈希对象
	//h := sha1.New()
	//// 2. 写入数据
	//// Write 方法永远不会返回错误，所以可以忽略返回值
	//h.Write([]byte(mergeStr))
	//// 3. 计算最终的和 (字节切片)
	//sum := h.Sum(nil)
	//
	//// 4. 将字节切片转换为十六进制字符串
	//return hex.EncodeToString(sum)

	parts := []string{timeStamp, nonce, token}
	sort.Strings(parts)
	merged := strings.Join(parts, "")
	sum := sha1.Sum([]byte(merged))
	return hex.EncodeToString(sum[:])
}

// VerifyPost 微信调用消息通知接口验证时候用
func VerifyPost(timeStamp, nonce, token, encrypt string) string {
	//mergeStr := fmt.Sprintf("%s%s%s%s", encrypt, timeStamp, nonce, token)
	//mergeStr := fmt.Sprintf("%s%s%s%s", encrypt, timeStamp, nonce, token)
	//// 1. 创建 SHA1 哈希对象
	//h := sha1.New()
	//// 2. 写入数据
	//// Write 方法永远不会返回错误，所以可以忽略返回值
	//h.Write([]byte(mergeStr))
	//// 3. 计算最终的和 (字节切片)
	//sum := h.Sum(nil)
	//
	//// 4. 将字节切片转换为十六进制字符串
	//return hex.EncodeToString(sum)

	parts := []string{encrypt, timeStamp, nonce, token}
	sort.Strings(parts)
	merged := strings.Join(parts, "")
	sum := sha1.Sum([]byte(merged))
	return hex.EncodeToString(sum[:])
}

func DecryptMessage(encrypt, encodingAESKey string) (string, string, error) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return "", "", err
	}
	if len(aesKey) < 32 {
		return "", "", errors.New("encodingaeskey长度不足")
	}
	cipherData, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", "", err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", "", err
	}
	if len(cipherData)%aes.BlockSize != 0 {
		return "", "", errors.New("密文长度非法")
	}
	iv := aesKey[:aes.BlockSize]
	mode := cipher.NewCBCDecrypter(block, iv)
	plain := make([]byte, len(cipherData))
	mode.CryptBlocks(plain, cipherData)
	plain, err = weixinPKCS7Unpad(plain, 32)
	if err != nil {
		return "", "", err
	}
	if len(plain) < 20 {
		return "", "", errors.New("解密数据长度不足")
	}
	msgLen := binary.BigEndian.Uint32(plain[16:20])
	msgStart := 20
	msgEnd := msgStart + int(msgLen)
	if msgEnd > len(plain) {
		return "", "", errors.New("解密消息长度非法")
	}
	msg := string(plain[msgStart:msgEnd])
	appId := ""
	if msgEnd < len(plain) {
		appId = string(plain[msgEnd:])
	}
	return msg, appId, nil
}

func weixinPKCS7Unpad(plain []byte, blockSize int) ([]byte, error) {
	if len(plain) == 0 {
		return nil, errors.New("解密数据为空")
	}
	padLen := int(plain[len(plain)-1])
	if padLen == 0 || padLen > blockSize || padLen > len(plain) {
		return nil, errors.New("padding非法")
	}
	for _, v := range plain[len(plain)-padLen:] {
		if int(v) != padLen {
			return nil, errors.New("padding不匹配")
		}
	}
	return plain[:len(plain)-padLen], nil
}
