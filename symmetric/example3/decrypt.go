package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

func main() {
	// Считывание симметричного ключа из файла
	key, err := readKeyFromFile("symmetric_key.bin")
	if err != nil {
		fmt.Println("Ошибка при считывании ключа из файла:", err)
		return
	}

	// Считывание зашифрованного сообщения из файла
	encryptedMessage, err := readEncryptedMessageFromFile("encrypted_message.bin")
	if err != nil {
		fmt.Println("Ошибка при считывании зашифрованного сообщения из файла:", err)
		return
	}

	// Дешифрование сообщения
	decryptedMessage, err := decryptMessage(key, encryptedMessage)
	if err != nil {
		fmt.Println("Ошибка при дешифровании сообщения:", err)
		return
	}

	fmt.Println("Дешифрованное сообщение:", string(decryptedMessage))
}

// Считывание ключа из файла
func readKeyFromFile(filename string) ([]byte, error) {
	key, err := os.ReadFile(filename)
	return key, err
}

// Считывание зашифрованного сообщения из файла
func readEncryptedMessageFromFile(filename string) ([]byte, error) {
	encryptedMessage, err := os.ReadFile(filename)
	return encryptedMessage, err
}

// Дешифрование сообщения с использованием симметричного ключа
func decryptMessage(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("неверный размер зашифрованного сообщения")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	decryptedMessage, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return decryptedMessage, nil
}
