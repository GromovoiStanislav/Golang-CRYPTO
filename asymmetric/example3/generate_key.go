package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {
	// Генерация пары асимметричных ключей
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Ошибка при генерации ключей:", err)
		return
	}

	// Сохранение закрытого ключа в файл
	savePrivateKeyToFile("private_key.pem", privateKey)

	// Сохранение открытого ключа в файл (или отправьте его другой стороне)
	savePublicKeyToFile("public_key.pem", &privateKey.PublicKey)


	fmt.Println("Закрытый и открытый ключи успешно сохранены в файлах.")
}

// Сохранение закрытого ключа в файл
func savePrivateKeyToFile(filename string, privateKey *rsa.PrivateKey) {
	privASN1 := x509.MarshalPKCS1PrivateKey(privateKey)
	privPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privASN1})

	err := ioutil.WriteFile(filename, privPEM, 0644)
	if err != nil {
		fmt.Println("Ошибка при сохранении закрытого ключа в файл:", err)
		return
	}
}

// Сохранение открытого ключа в файл
func savePublicKeyToFile(filename string, publicKey *rsa.PublicKey) {
	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Ошибка при сохранении открытого ключа:", err)
		return
	}

	pubPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubASN1})

	err = ioutil.WriteFile(filename, pubPEM, 0644)
	if err != nil {
		fmt.Println("Ошибка при сохранении открытого ключа в файл:", err)
		return
	}
}
