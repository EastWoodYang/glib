package glib

import (
	"bufio"
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"strings"
)

/* ================================================================================
 * 安全
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Md5哈希
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Md5(data string) string {
	m := md5.New()
	io.WriteString(m, data)
	return hex.EncodeToString(m.Sum(nil))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Sha1哈希
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Sha256哈希
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Sha256(data string) string {
	t := sha256.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Hmac Sha1哈希
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HmacSha1(data string, key string, args ...bool) string {
	resultString := ""
	isHex := true

	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))

	if len(args) > 0 {
		isHex = args[0]
	}

	if isHex {
		resultString = hex.EncodeToString(mac.Sum(nil))
	} else {
		resultString = string(mac.Sum(nil))
	}
	return resultString
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Hmac Sha256哈希
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func HmacSha256(data string, key string, args ...bool) string {
	resultString := ""
	isHex := true

	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))

	if len(args) > 0 {
		isHex = args[0]
	}

	if isHex {
		resultString = hex.EncodeToString(mac.Sum(nil))
	} else {
		resultString = string(mac.Sum(nil))
	}
	return resultString
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Sha256WithRsa签名
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Sha256WithRsa(data string, privateKey string) (string, error) {
	//key to pem
	privateKey = RsaPrivateToMultipleLine(privateKey)

	//RSA密匙
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("sign private key decode error")
	}

	prk8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	//SHA256哈希
	h := sha256.New()
	h.Write([]byte(data))
	digest := h.Sum(nil)

	//RSA签名
	rsaPrivateKey := prk8.(*rsa.PrivateKey)
	s, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, []byte(digest))
	if err != nil {
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(s)
	return string(sign), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Sha256WithRsa验签
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func Sha256WithRsaVerify(data, sign, publicKey string) (bool, error) {
	// publickey to pem
	publicKey = RsaPublicToMultipleLine(publicKey)

	// 加载RSA的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return false, errors.New("sign public key decode error")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	rsaPublicKey, _ := pub.(*rsa.PublicKey)

	h := sha256.New()
	h.Write([]byte(data))
	digest := h.Sum(nil)

	// base64 decode,支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
	signString, _ := base64.StdEncoding.DecodeString(sign)
	hex.EncodeToString(signString)

	// 调用rsa包的VerifyPKCS1v15验证签名有效性
	if err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, digest, signString); err != nil {
		return false, err
	}

	return true, nil
}

/*
func testAes() {
	key := []byte("axewfd3_r44&98Klaxewfd3_r44&98Kl")
	result, err := AESEncrypt([]byte("mliu"), key)
	if err != nil {
		panic(err)
	}
	s := base64.StdEncoding.EncodeToString(result)
	origData, err := AESDecrypt(result, key)
	ss := string(origData)
}*/

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Des加密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key[0:8])
	if err != nil {
		return nil, err
	}
	origData = Pkcs5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Des解密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key[0:8])
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = Pkcs5UnPadding(origData)
	return origData, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Aes加密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func AesEncrypt(origData, key []byte, args ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key[0:16])
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	iv := key[:blockSize]
	if len(args) > 0 {
		iv = args[0][:blockSize]
	}

	origData = Pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return crypted, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Aes解密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func AesDecrypt(crypted, key []byte, args ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key[0:16])
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	iv := key[:blockSize]
	if len(args) > 0 {
		iv = args[0][:blockSize]
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))

	blockMode.CryptBlocks(origData, crypted)
	origData = Pkcs5UnPadding(origData)

	return origData, nil
}

func Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Rsa加密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RSAEncrypt(origData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(crand.Reader, pub, origData)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Rsa解密
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RSADecrypt(cipherData, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(crand.Reader, priv, cipherData)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 生成RsaKey
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GenRsaKey(bits int) (string, string, error) {
	//私钥
	privateKey, err := rsa.GenerateKey(crand.Reader, bits)
	if err != nil {
		return "", "", err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateBuffer := bytes.NewBuffer(make([]byte, 0))
	privateWriter := bufio.NewWriter(privateBuffer)
	err = pem.Encode(privateWriter, block)
	if err != nil {
		return "", "", err
	}

	// 生成公钥
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	publicBuffer := bytes.NewBuffer(make([]byte, 0))
	publicWriter := bufio.NewWriter(publicBuffer)
	err = pem.Encode(publicWriter, block)
	if err != nil {
		return "", "", err
	}

	return privateBuffer.String(), publicBuffer.String(), nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 将单行的Ras Public字符串格式化为多行格式
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RsaPublicToMultipleLine(privateKey string) string {
	privateKeys := make([]string, 0)
	privateKeys = append(privateKeys, "-----BEGIN PUBLIC KEY-----")

	for i := 0; i < 4; i++ {
		start := i * 64
		end := (i + 1) * 64
		lineKey := ""
		if i == 3 {
			lineKey = privateKey[start:]
		} else {
			lineKey = privateKey[start:end]
		}
		privateKeys = append(privateKeys, lineKey)
	}

	privateKeys = append(privateKeys, "-----END PUBLIC KEY-----")
	privateKey = strings.Join(privateKeys, "\r\n")

	return privateKey
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 将单行的Ras Private字符串格式化为多行格式
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RsaPrivateToMultipleLine(privateKey string) string {
	privateKeys := make([]string, 0)
	privateKeys = append(privateKeys, "-----BEGIN PRIVATE KEY-----")

	for i := 0; i < 26; i++ {
		start := i * 64
		end := (i + 1) * 64
		lineKey := ""
		if i == 25 {
			lineKey = privateKey[start:]
		} else {
			lineKey = privateKey[start:end]
		}
		privateKeys = append(privateKeys, lineKey)
	}

	privateKeys = append(privateKeys, "-----END PRIVATE KEY-----")
	privateKey = strings.Join(privateKeys, "\r\n")

	return privateKey
}
