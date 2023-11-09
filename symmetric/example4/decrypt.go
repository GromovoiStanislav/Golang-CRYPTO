package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

func main() {
	// Считываем симметричный ключ из файла
	key, err := ioutil.ReadFile("symmetric_key.bin")
	if err != nil {
		fmt.Println("Ошибка при считывании ключа из файла:", err)
		return
	}

	// Считываем зашифрованное сообщение из файла
	encryptedMessage, err := ioutil.ReadFile("encrypted_message.bin")
	if err != nil {
		fmt.Println("Ошибка при считывании зашифрованного сообщения из файла:", err)
		return
	}

	// Считываем IV (Initialization Vector) из первых 12 байт зашифрованного сообщения
	iv := encryptedMessage[:12]
	// Оставшиеся байты после первых 12 - это сами зашифрованные данные
	encryptedData := encryptedMessage[12:]

	// Создаем блочный шифр с использованием ключа
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Ошибка при создании блочного шифра:", err)
		return
	}

	// Создаем объект GCM для дешифрования
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Ошибка при создании объекта GCM:", err)
		return
	}

	// Дешифруем данные
	decryptedMessage, err := aesGCM.Open(nil, iv, encryptedData, nil)
	if err != nil {
		fmt.Println("Ошибка при дешифровании данных:", err)
		return
	}

	fmt.Println("Дешифрованное сообщение:", string(decryptedMessage))
}
