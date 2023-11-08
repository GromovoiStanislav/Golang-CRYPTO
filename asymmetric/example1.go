package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
    // Генерируем пару ключей RSA
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        panic(err)
    }



    // Создаем сообщение
    message := "Hello, world!"

    // Шифруем сообщение публичным ключом
    encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privateKey.PublicKey, []byte(message),[]byte("my_solt"))
    if err != nil {
        panic(err)
    }
	// Выводим дешифрованное сообщение
    //fmt.Println(encrypted)

    // Дешифруем сообщение закрытым ключом
    decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encrypted,[]byte("my_solt"))
    if err != nil {
        panic(err)
    }

    // Выводим дешифрованное сообщение
    fmt.Println(string(decrypted))
}
