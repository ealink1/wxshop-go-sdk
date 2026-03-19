package wxshop_go_sdk

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// MessagePushCrypto 消息推送加解密处理
type MessagePushCrypto struct {
	EncodingAESKey string // EncodingAESKey，来自微信配置
	Token          string // Token 令牌，来自微信配置
	AppID          string // 微信小店 AppID
	AESKey         []byte // 解码后的 AES 密钥
}

// NewMessagePushCrypto 创建一个新的消息推送加解密处理器
// encodingAESKey: 微信配置的 EncodingAESKey
// token: 微信配置的 Token 令牌
// appID: 微信小店 AppID
func NewMessagePushCrypto(encodingAESKey, token, appID string) (*MessagePushCrypto, error) {
	// Base64 解码 AESKey，需要在末尾添加 "="
	aesKey, err := base64.StdEncoding.DecodeString(encodingAESKey + "=")
	if err != nil {
		return nil, fmt.Errorf("解码 AESKey 失败：%w", err)
	}

	if len(aesKey) != 32 {
		return nil, fmt.Errorf("AESKey 长度错误，应为 32 字节，实际：%d", len(aesKey))
	}

	return &MessagePushCrypto{
		EncodingAESKey: encodingAESKey,
		Token:          token,
		AppID:          appID,
		AESKey:         aesKey,
	}, nil
}

// VerifySignature 验证推送消息的签名（用于服务器验证）
// signature: 微信发送的签名参数
// timestamp: 时间戳参数
// nonce: 随机数参数
// echostr: 随机字符串参数
func (c *MessagePushCrypto) VerifySignature(signature, timestamp, nonce, echostr string) error {
	// 将 token、timestamp、nonce 三个参数进行字典序排序
	strs := []string{c.Token, timestamp, nonce}
	sort.Strings(strs)

	// 拼接成一个字符串
	combined := strings.Join(strs, "")

	// 计算 SHA1 签名
	sha := sha1.New()
	sha.Write([]byte(combined))
	computedSignature := fmt.Sprintf("%x", sha.Sum(nil))

	// 校验签名
	if computedSignature != signature {
		return fmt.Errorf("签名验证失败，期望：%s, 计算得到：%s", signature, computedSignature)
	}

	return nil
}

// VerifyMessageSignature 验证消息体的签名（安全模式）
// msgSignature: 微信发送的消息签名
// timestamp: 时间戳参数
// nonce: 随机数参数
// encrypt: 加密的消息内容（Encrypt 字段）
func (c *MessagePushCrypto) VerifyMessageSignature(msgSignature, timestamp, nonce, encrypt string) error {
	// 将 token、timestamp、nonce、encrypt 四个参数进行字典序排序
	strs := []string{c.Token, timestamp, nonce, encrypt}
	sort.Strings(strs)

	// 拼接成一个字符串
	combined := strings.Join(strs, "")

	// 计算 SHA1 签名
	sha := sha1.New()
	sha.Write([]byte(combined))
	computedSignature := fmt.Sprintf("%x", sha.Sum(nil))

	// 校验签名
	if computedSignature != msgSignature {
		return fmt.Errorf("消息签名验证失败，期望：%s, 计算得到：%s", msgSignature, computedSignature)
	}

	return nil
}

// DecryptMessage 解密消息体
// encrypt: Base64 编码的加密消息内容
// 返回解密后的明文 JSON 字符串
func (c *MessagePushCrypto) DecryptMessage(encrypt string) (string, error) {
	// Base64 解码
	decoded, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", fmt.Errorf("Base64 解码失败：%w", err)
	}

	// AES 解密（CBC 模式）
	decrypted, err := c.aesDecrypt(decoded, c.AESKey)
	if err != nil {
		return "", fmt.Errorf("AES 解密失败：%w", err)
	}

	// 去除 PKCS#7 填充
	unpadded := c.pkcs7Unpad(decrypted)
	if unpadded == nil {
		return "", fmt.Errorf("PKCS#7 去填充失败")
	}

	// 解析数据结构：random(16B) + msg_len(4B) + msg + appid
	if len(unpadded) < 20 {
		return "", fmt.Errorf("解密数据长度过短：%d", len(unpadded))
	}

	// 跳过前 16 字节的随机字符串
	// 读取 4 字节的消息长度（网络字节序）
	msgLen := binary.BigEndian.Uint32(unpadded[16:20])

	// 提取消息内容
	if int(msgLen)+20 > len(unpadded) {
		return "", fmt.Errorf("消息长度超出范围，msg_len=%d, 总长度=%d", msgLen, len(unpadded))
	}

	msg := string(unpadded[20 : 20+msgLen])

	// 提取并验证 AppID
	appid := string(unpadded[20+msgLen:])
	if appid != c.AppID {
		return "", fmt.Errorf("AppID 不匹配，期望：%s, 实际：%s", c.AppID, appid)
	}

	return msg, nil
}

// EncryptMessage 加密消息体
// msg: 要加密的明文消息
// 返回 Base64 编码的加密消息
func (c *MessagePushCrypto) EncryptMessage(msg string) (string, error) {
	// 构造数据结构：random(16B) + msg_len(4B) + msg + appid
	randomBytes := c.generateRandomBytes(16)
	msgBytes := []byte(msg)
	appidBytes := []byte(c.AppID)

	// 计算总长度
	totalLen := 16 + 4 + len(msgBytes) + len(appidBytes)

	// 创建缓冲区
	buffer := make([]byte, totalLen)

	// 填充随机字符串
	copy(buffer[0:16], randomBytes)

	// 填充消息长度（网络字节序）
	binary.BigEndian.PutUint32(buffer[16:20], uint32(len(msgBytes)))

	// 填充消息内容
	copy(buffer[20:20+len(msgBytes)], msgBytes)

	// 填充 AppID
	copy(buffer[20+len(msgBytes):], appidBytes)

	// PKCS#7 填充
	padded := c.pkcs7Pad(buffer, 32)

	// AES 加密（CBC 模式）
	encrypted, err := c.aesEncrypt(padded, c.AESKey)
	if err != nil {
		return "", fmt.Errorf("AES 加密失败：%w", err)
	}

	// Base64 编码
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// GenerateCallbackSignature 生成回调消息的签名
// timestamp: 时间戳
// nonce: 随机数
// encrypt: 加密后的消息内容
func (c *MessagePushCrypto) GenerateCallbackSignature(timestamp, nonce, encrypt string) string {
	// 将 token、timestamp、nonce、encrypt 四个参数进行字典序排序
	strs := []string{c.Token, timestamp, nonce, encrypt}
	sort.Strings(strs)

	// 拼接成一个字符串
	combined := strings.Join(strs, "")

	// 计算 SHA1 签名
	sha := sha1.New()
	sha.Write([]byte(combined))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

// BuildEncryptedResponse 构建加密响应包
// msg: 要回复的明文消息
// 返回加密后的响应包（JSON 格式）
func (c *MessagePushCrypto) BuildEncryptedResponse(msg string) (map[string]interface{}, error) {
	// 加密消息
	encrypted, err := c.EncryptMessage(msg)
	if err != nil {
		return nil, fmt.Errorf("加密消息失败：%w", err)
	}

	// 生成时间戳和随机数
	timestamp := time.Now().Unix()
	nonce := fmt.Sprintf("%d", rand.Intn(1000000000))

	// 生成签名
	signature := c.GenerateCallbackSignature(fmt.Sprintf("%d", timestamp), nonce, encrypted)

	// 构建响应包
	response := map[string]interface{}{
		"Encrypt":      encrypted,
		"MsgSignature": signature,
		"TimeStamp":    timestamp,
		"Nonce":        nonce,
	}

	return response, nil
}

// ==================== AES 加密解密辅助方法 ====================

// aesDecrypt AES 解密（CBC 模式）
func (c *MessagePushCrypto) aesDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(ciphertext) < blockSize {
		return nil, fmt.Errorf("密文长度不足")
	}

	// CBC 模式需要提取 IV（前 16 字节）
	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext, nil
}

// aesEncrypt AES 加密（CBC 模式）
func (c *MessagePushCrypto) aesEncrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	// 生成随机 IV
	iv := c.generateRandomBytes(blockSize)

	// PKCS#7 填充
	plaintext = c.pkcs7Pad(plaintext, blockSize)

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	// 返回：IV + 密文
	result := append(iv, ciphertext...)
	return result, nil
}

// pkcs7Pad PKCS#7 填充
func (c *MessagePushCrypto) pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7Unpad PKCS#7 去填充
func (c *MessagePushCrypto) pkcs7Unpad(data []byte) []byte {
	if len(data) == 0 {
		return nil
	}

	padding := int(data[len(data)-1])
	if padding > len(data) || padding > 32 {
		return nil
	}

	for i := 0; i < padding; i++ {
		if data[len(data)-1-i] != byte(padding) {
			return nil
		}
	}

	return data[:len(data)-padding]
}

// generateRandomBytes 生成随机字节
func (c *MessagePushCrypto) generateRandomBytes(length int) []byte {
	bytes := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range bytes {
		bytes[i] = byte(rand.Intn(256))
	}
	return bytes
}

// ==================== 便捷函数 ====================

// VerifyAndDecryptMessage 验证签名并解密消息（一站式处理）
// msgSignature: 消息签名
// timestamp: 时间戳
// nonce: 随机数
// encrypt: 加密的消息内容
// crypto: 加解密处理器实例
func VerifyAndDecryptMessage(crypto *MessagePushCrypto, msgSignature, timestamp, nonce, encrypt string) (string, error) {
	// 验证签名
	if err := crypto.VerifyMessageSignature(msgSignature, timestamp, nonce, encrypt); err != nil {
		return "", err
	}

	// 解密消息
	return crypto.DecryptMessage(encrypt)
}
