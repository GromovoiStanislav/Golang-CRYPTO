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
	// Считывание закрытого ключа из файла
	privateKeyBytes, err := ioutil.ReadFile("private_key.pem")
	if err != nil {
		fmt.Println("Ошибка при чтении закрытого ключа:", err)
		return
	}

	// Декодирование закрытого ключа
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		fmt.Println("Некорректный формат закрытого ключа")
		return
	}

	// Преобразование закрытого ключа
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Ошибка при преобразовании закрытого ключа:", err)
		return
	}

	// Считывание зашифрованного сообщения из файла
	encryptedMessage, err := ioutil.ReadFile("encrypted_message.bin")
	if err != nil {
		fmt.Println("Ошибка при чтении зашифрованного сообщения:", err)
		return
	}

	// Дешифрование сообщения
	decryptedMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedMessage)
	if err != nil {
		fmt.Println("Ошибка при дешифровании сообщения:", err)
		return
	}

	fmt.Println("Дешифрованное сообщение:", string(decryptedMessage))
}
