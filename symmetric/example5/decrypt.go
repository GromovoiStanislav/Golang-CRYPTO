package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Считываем симметричный ключ и IV из файла
	keyHex, err := ioutil.ReadFile("symmetric_key.txt")
	if err != nil {
		log.Fatal(err)
	}

	key, err := hex.DecodeString(strings.TrimSpace(string(keyHex)))
	if err != nil {
		log.Fatal(err)
	}

	encryptedData, err := ioutil.ReadFile("encrypted_message.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Разделяем строку на IV и зашифрованное сообщение
	parts := strings.Split(strings.TrimSpace(string(encryptedData)), "\n")
	iv, err := hex.DecodeString(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	encryptedMessage, err := hex.DecodeString(parts[1])
	if err != nil {
		log.Fatal(err)
	}

	// Создаем новый блочный шифр AES
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем объект расшифровщика с режимом CBC
	mode := cipher.NewCBCDecrypter(block, iv)

	// Дешифруем сообщение
	decrypted := make([]byte, len(encryptedMessage))
	mode.CryptBlocks(decrypted, encryptedMessage)

	// Убираем паддинг, если используется PKCS7
	decrypted = PKCS7Unpad(decrypted)

	// Выводим дешифрованное сообщение
	log.Printf("Дешифрованное сообщение: %s", decrypted)
}

func PKCS7Unpad(data []byte) []byte {
    if len(data) == 0 {
        return nil
    }
    pad := int(data[len(data)-1])
    return data[:len(data)-pad]
}
