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
	// Путь к файлу с закрытым ключом
	privateKeyPath := "private_key.pem"

	// Загрузка закрытого ключа из файла
	privateKeyBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Println("Ошибка при чтении закрытого ключа:", err)
		return
	}

	// Преобразование закрытого ключа в структуру rsa.PrivateKey
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil {
		fmt.Println("Не удалось декодировать закрытый ключ")
		return
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Ошибка при парсинге закрытого ключа:", err)
		return
	}

	// Приведение закрытого ключа к типу *rsa.PrivateKey
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Println("Ошибка при приведении закрытого ключа к типу *rsa.PrivateKey")
		return
	}

	// Сообщение, которое вы хотите зашифровать
	message := []byte("Это сообщение, которое мы хотим зашифровать.")

	// Шифрование сообщения
	encryptedMessage, err := rsa.EncryptPKCS1v15(rand.Reader, &rsaPrivateKey.PublicKey, message)
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
