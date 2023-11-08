package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func main() {
    // Генерируем ключ для шифрования сообщений
    key := make([]byte, 32)
    _, err := rand.Read(key)
    if err != nil {
      fmt.Println("Ошибка при создании ключа:", err)
      return
    }

     // сообщение
     message := "Hello, world!"

    // Шифруем сообщение
    block, err := aes.NewCipher(key)
    if err != nil {
      fmt.Println("Ошибка при шифровании сообщения")
      panic(err)
    }
    // Создаем шифратор
    gcm, err := cipher.NewGCM(block)
    if err != nil {
      fmt.Println("Ошибка при шифровании сообщения")
      panic(err)
    }
    // Шифруем 
    nonce := make([]byte, gcm.NonceSize())
    if _, err := rand.Read(nonce); err != nil {
      fmt.Println("Ошибка при шифровании сообщения")
      panic(err)
  }
    encrypted := gcm.Seal(nonce, nonce, []byte(message), nil)

 


    // Дешифруем сообщение
    block2, err := aes.NewCipher(key)
    if err != nil {
      fmt.Println("Ошибка при дешифровании сообщения")
      panic(err)
    }

    decipher, err := cipher.NewGCM(block2)
    if err != nil {
      fmt.Println("Ошибка при дешифровании сообщения")
      panic(err)
    }

    nonceSize := gcm.NonceSize()
    if len(encrypted) < nonceSize {
      fmt.Println("неверный размер зашифрованного сообщения")
      panic(err)
    }
  
    nonce, encrypted = encrypted[:nonceSize], encrypted[nonceSize:]

    decrypted, err := decipher.Open(nil, nonce, encrypted, nil)
    if err != nil {
        panic(err)
    }

    // Выводим дешифрованное сообщение
    fmt.Println(string(decrypted))
}
