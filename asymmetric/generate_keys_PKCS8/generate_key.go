package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func main() {
	// Генерируем пару ключей RSA
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		panic(err)
	}

	// Кодируем закрытый ключ в формате PEM
	pemEncodedPrivateKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// Сохраняем закрытый ключ в файл
	if err := ioutil.WriteFile("private.pem", pemEncodedPrivateKey, 0600); err != nil {
		panic(err)
	}

	// Извлекаем открытый ключ из закрытого
	publicKey := &privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}

	// Кодируем открытый ключ в формате PEM
	pemEncodedPublicKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	// Сохраняем открытый ключ в файл
	if err := ioutil.WriteFile("public.pem", pemEncodedPublicKey, 0644); err != nil {
		panic(err)
	}
}
