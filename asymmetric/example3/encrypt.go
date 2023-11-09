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
	// Пути к файлам с открытым и закрытым ключами
	//publicKeyPath := "public_key.pem"
	privateKeyPath := "private_key.pem"

	// Загрузка закрытого ключа из файла
	privateKeyBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Println("Ошибка при чтении закрытого ключа:", err)
		return
	}

	// Преобразование закрытого ключа в структуру rsa.PrivateKey
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		fmt.Println("Неверный формат закрытого ключа")
		return
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Ошибка при парсинге закрытого ключа:", err)
		return
	}

	// Сообщение, которое вы хотите зашифровать
	message := []byte("Это сообщение, которое мы хотим зашифровать.")

	// Шифрование сообщения
	encryptedMessage, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, message)
	if err != nil {
		fmt.Println("Ошибка при шифровании сообщения:", err)
		return
	}

	// Сохранение зашифрованного сообщения в файл
	err = ioutil.WriteFile("encrypted_message.bin", encryptedMessage, 0644)
	if err != nil {
		fmt.Println("Ошибка при сохранении зашифрованного сообщения в файл:", err)
		return
	}

	fmt.Println("Зашифрованное сообщение успешно сохранено в файле.")
}
