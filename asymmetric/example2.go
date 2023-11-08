package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

func main() {
    // Генерируем пару ключей RSA
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        fmt.Println("Ошибка при генерации ключей")
        panic(err)
    }

    // Создаем сообщение
    message := "Hello, world!"

    // Шифруем сообщение публичным ключом
    encrypted, err := rsa.EncryptPKCS1v15(rand.Reader,  &privateKey.PublicKey, []byte(message))
	if err != nil {
		fmt.Println("Ошибка при шифровании сообщения")
		panic(err)
	}
    
	// Выводим шифрованное сообщение
    //fmt.Println(encrypted) //байты

    // Преобразование зашифрованного сообщения в Base64
    encryptedBase64 := base64.StdEncoding.EncodeToString(encrypted)
    fmt.Println(encryptedBase64)


    // Декодирование сообщения из Base64 в байты
	encryptedMessage, err := base64.StdEncoding.DecodeString(string(encryptedBase64))

    // Дешифруем сообщение закрытым ключом
    decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedMessage)
    if err != nil {
        fmt.Println("Ошибка при дешифровании сообщения")
        panic(err)
    }

    // Выводим дешифрованное сообщение
    fmt.Println(string(decrypted))
}
