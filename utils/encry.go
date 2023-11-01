package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
)

// PublicKeyStr RSA public key
const PublicKeyStr = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDAwLU8jpzUCEbzaLVpRgKAF0BG
DuQnqRVHGDvP75bWq2T03JXv2EtJtz6L4n7ZvMYK0gA6+/kfrfMKhap5Kg8Afc5W
Bp1CWE0NsUPAL2OekjhYO7qZue5M1UCGu6XmJ3I6qKJzozWAsEBuQuwV0MrLTQbr
28c+QOTmLnD6a4ZSOQIDAQAB
-----END PUBLIC KEY-----`

func EncryPassWord(password string) (string, string, error) {
	// Parsing the public key
	block, _ := pem.Decode([]byte(PublicKeyStr))
	pub, _ := x509.ParsePKIXPublicKey(block.Bytes)
	rsaPub, _ := pub.(*rsa.PublicKey)

	// RSA encryption
	encryptedPassword, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(password))
	if err != nil {
		log.Fatal("Failed to encrypt: ", err)
		return "", "", err
	}
	// Base64 encode the encrypted password
	strRSAPassword := base64.StdEncoding.EncodeToString(encryptedPassword)
	// MD5 hash the password
	md5Hash := md5.Sum([]byte(password))
	strPassword := fmt.Sprintf("%x", md5Hash)
	return strRSAPassword, strPassword, nil
}
