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

	// Кодируем закрытый ключ в формате PEM
	pemEncodedPrivateKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	// Сохраняем закрытый ключ в файл
	if err := ioutil.WriteFile("private.pem", pemEncodedPrivateKey, 0600); err != nil {
		panic(err)
	}

	// Извлекаем открытый ключ из закрытого
	publicKey := &privateKey.PublicKey

	// Кодируем открытый ключ в формате PEM
	pemEncodedPublicKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})


    // Сохраняем открытый ключ в файл
	if err := ioutil.WriteFile("public.pem", pemEncodedPublicKey, 0644); err != nil {
		panic(err)
	}
}
