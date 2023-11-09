package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func main() {
	// Считывание симметричного ключа из файла
	key, err := readKeyFromFile("symmetric_key.bin")
	if err != nil {
		fmt.Println("Ошибка при считывании ключа из файла:", err)
		return
	}

	// Сообщение, которое нужно зашифровать
	message := "Это секретное сообщение, которое мы хотим зашифровать."

	// Шифрование сообщения
	encryptedMessage, err := encryptMessage(key, message)
	if err != nil {
		fmt.Println("Ошибка при шифровании сообщения:", err)
		return
	}

	fmt.Println("Зашифрованное сообщение:")
	fmt.Printf("%x\n", encryptedMessage)

	// Запись зашифрованного сообщения в файл
	err = writeToFile("encrypted_message.bin", encryptedMessage)
	if err != nil {
		fmt.Println("Ошибка при записи зашифрованного сообщения в файл:", err)
		return
	}

	fmt.Println("Зашифрованное сообщение успешно записано в файл.")
}

// Считывание ключа из файла
func readKeyFromFile(filename string) ([]byte, error) {
	key, err := os.ReadFile(filename)
	return key, err
}

// Шифрование сообщения с использованием симметричного ключа
func encryptMessage(key []byte, message string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Шифрование сообщения
	encryptedMessage := gcm.Seal(nonce, nonce, []byte(message), nil)

	return encryptedMessage, nil
}

// Запись ключа в файл
func writeToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}