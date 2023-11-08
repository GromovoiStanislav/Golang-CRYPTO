package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	// Создание симметричного ключа (256 бит)
	key, err := generateSymmetricKey()
	if err != nil {
		fmt.Println("Ошибка при создании симметричного ключа:", err)
		return
	}

	// Сообщение, которое вы хотите зашифровать
	message := []byte("Это сообщение, которое мы хотим зашифровать.")

	// Шифрование сообщения
	encryptedMessage, err := encryptMessage(key, message)
	if err != nil {
		fmt.Println("Ошибка при шифровании сообщения:", err)
		return
	}

	fmt.Println("Сообщение успешно зашифровано.")

	// Дешифрование сообщения
	decryptedMessage, err := decryptMessage(key, encryptedMessage)
	if err != nil {
		fmt.Println("Ошибка при дешифровании сообщения:", err)
		return
	}

	fmt.Println("Дешифрованное сообщение:")
	fmt.Println(string(decryptedMessage))
}

// Создание симметричного ключа
func generateSymmetricKey() ([]byte, error) {
	key := make([]byte, 32) // 32 байта (256 бит) для AES-256
	_, err := rand.Read(key)
	return key, err
}

// Шифрование сообщения с использованием симметричного ключа
func encryptMessage(key, plaintext []byte) ([]byte, error) {
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

	encryptedMessage := gcm.Seal(nonce, nonce, plaintext, nil)

	return encryptedMessage, nil
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
