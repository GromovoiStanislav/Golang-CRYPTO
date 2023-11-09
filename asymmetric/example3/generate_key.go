package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
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
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		fmt.Println("Ошибка при создании файла для закрытого ключа:", err)
		return
	}
	defer privateKeyFile.Close()

	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		fmt.Println("Ошибка при маршалинге закрытого ключа:", err)
		return
	}

	privateKeyBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	err = pem.Encode(privateKeyFile, privateKeyBlock)
	if err != nil {
		fmt.Println("Ошибка при записи закрытого ключа в файл:", err)
		return
	}

	fmt.Println("Закрытый ключ успешно сохранен в файле private_key.pem.")
}

// Сохранение открытого ключа в файл
func savePublicKeyToFile(filename string, publicKey *rsa.PublicKey) {
	publicKeyFile, err := os.Create("public_key.pem")
	if err != nil {
		fmt.Println("Ошибка при создании файла для открытого ключа:", err)
		return
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Ошибка при маршалинге открытого ключа:", err)
		return
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	err = pem.Encode(publicKeyFile, publicKeyBlock)
	if err != nil {
		fmt.Println("Ошибка при записи открытого ключа в файл:", err)
		return
	}

	fmt.Println("Открытый ключ успешно сохранен в файле public_key.pem.")
}