package main

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"testing"
)

func TestName(t *testing.T) {
	ctx := context.Background()
	var privateKey = ""
	client, err := alipay.New("", privateKey, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	//使用支付宝公钥
	var aliPublicKey = ""
	err = client.LoadAliPayPublicKey(aliPublicKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	var p = alipay.BillAccountLogQuery{
		StartTime: "2025-01-01 00:00:00",
		EndTime:   "2025-01-06 00:00:00",
	}
	result, err := client.BillAccountLogQuery(ctx, p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}

func getPrivateKeyFromPKCS8(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("private key error: no PEM block found")
	}

	privateKeyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	privateKeyRSA, ok := privateKeyInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key error: not *rsa.PrivateKey type")
	}

	return privateKeyRSA, nil
}

// doSign 使用SHA256withRSA算法对内容进行签名
func doSign(content, privateKey, charset string) (string, error) {
	privateKeyRSA, err := getPrivateKeyFromPKCS8(privateKey)
	if err != nil {
		return "", err
	}

	hash := sha256.New()
	hash.Write([]byte(content))
	hashed := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKeyRSA, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}

	// Base64编码签名结果
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	return encodedSignature, nil
}
