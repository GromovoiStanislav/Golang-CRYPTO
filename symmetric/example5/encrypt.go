package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Генерируем симметричный ключ
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}

	// Записываем ключ в файл
	err = ioutil.WriteFile("symmetric_key.txt", []byte(hex.EncodeToString(key)), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Сообщение для шифрования
	plaintext := "Привет, мир!"

	// Генерируем случайный IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем новый блочный шифр AES
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем объект шифрования с режимом CBC
	mode := cipher.NewCBCEncrypter(block, iv)

	// Шифруем сообщение
	plaintextBytes := []byte(plaintext)
	paddedPlaintext := PKCS7Pad(plaintextBytes, aes.BlockSize)
	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	// Записываем IV и зашифрованное сообщение в файл
	encryptedData := fmt.Sprintf("%x\n%x", iv, ciphertext)
	err = ioutil.WriteFile("encrypted_message.txt", []byte(encryptedData), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Шифрование завершено.")
}

func PKCS7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := strings.Repeat(string(byte(padding)), padding)
	return append(data, []byte(padText)...)
}
