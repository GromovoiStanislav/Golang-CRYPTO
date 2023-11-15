package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

func main() {
	// Генерация простых чисел p и g
	p, _ := rand.Prime(rand.Reader, 128) // 128 бит для примера, можно выбрать другую длину
	g := big.NewInt(2)                     // генератор

	// Закрытые ключи для Alice и Bob
	privateKeyAlice, _ := rand.Int(rand.Reader, p)
	privateKeyBob, _ := rand.Int(rand.Reader, p)

	// Вычисление открытых ключей
	publicKeyAlice := new(big.Int).Exp(g, privateKeyAlice, p)
	publicKeyBob := new(big.Int).Exp(g, privateKeyBob, p)

	// Обмен открытыми ключами

	// Алиса получает открытый ключ Боба и вычисляет общий секретный ключ
	sharedKeyAlice := new(big.Int).Exp(publicKeyBob, privateKeyAlice, p)

	// Боб получает открытый ключ Алисы и вычисляет общий секретный ключ
	sharedKeyBob := new(big.Int).Exp(publicKeyAlice, privateKeyBob, p)

	// Печать общих секретных ключей
	fmt.Println("Общий секретный ключ Алисы:", sharedKeyAlice)
	fmt.Println("Общий секретный ключ Боба:", sharedKeyBob)


	// Пример сообщения для шифрования
	message := []byte("Привет, Diffie-Hellman!")

	// Используем AES для шифрования
	ciphertext, err := encrypt(message, sharedKeyAlice)
	if err != nil {
		fmt.Println("Ошибка при шифровании:", err)
		return
	}

	fmt.Println("Зашифрованное сообщение:", ciphertext)
	fmt.Println("Зашифрованное сообщение:", string(ciphertext) )

	// Дешифрование
	decryptedMessage, err := decrypt(ciphertext, sharedKeyBob)
	if err != nil {
		fmt.Println("Ошибка при дешифровании:", err)
		return
	}

	fmt.Println("Расшифрованное сообщение:", string(decryptedMessage))
}


func encrypt(plaintext []byte, key *big.Int) ([]byte, error) {
	// Преобразование общего секретного ключа в нужный размер
	keyBytes := key.Bytes()
	keyBytes = append(make([]byte, 0, 32-len(keyBytes)), keyBytes...)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	// Дополнение до размера блока
	plaintext = pkcs7Pad(plaintext, aes.BlockSize)

	ciphertext := make([]byte, len(plaintext))
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	// Добавление IV к шифрованному тексту
	ciphertext = append(iv, ciphertext...)

	return ciphertext, nil
}

func decrypt(ciphertext []byte, key *big.Int) ([]byte, error) {
	// Преобразование общего секретного ключа в нужный размер
	keyBytes := key.Bytes()
	keyBytes = append(make([]byte, 0, 32-len(keyBytes)), keyBytes...)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	// Разделение IV и шифрованного текста
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Дешифрование
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Удаление дополнения
	ciphertext = pkcs7Unpad(ciphertext)

	return ciphertext, nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}